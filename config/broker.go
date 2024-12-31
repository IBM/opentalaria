package config

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/netip"
	"net/url"
	"opentalaria/utils"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Broker struct {
	BrokerID int32
	Rack     *string
	// https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#listeners
	Listeners []Listener
	// https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#advertised-listeners
	AdvertisedListeners []Listener
}

type Listener struct {
	Host string
	Port int32
	// If the listener name is a security protocol, like PLAINTEXT,SSL,SASL_PLAINTEXT,SASL_SSL,
	// the name will be set as SecurityProtocol. Otherwise the name should be mapped in listener.security.protocol.map.
	// see https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#listener-security-protocol-map.
	SecurityProtocol SecurityProtocol
	ListenerName     string
}

// var (
// 	// by default in KRaft mode, generated broker IDs start from reserved.broker.max.id + 1,
// 	// where reserved.broker.max.id=1000 if the property is not set.
// 	// KRaft mode is the default Kafka mode, since Kafka v3.3.1, so OpenTalaria will implement default settings in KRaft mode.
// 	RESERVED_BROKER_MAX_ID = 1000
// )

// NewBroker returns a new instance of Broker.
// For now OpenTalaria does not support rack awareness, but this will change in the future.
func NewBroker(env *viper.Viper) (*Broker, error) {
	broker := Broker{}

	listenerStr := env.GetString("listeners")
	if listenerStr == "" {
		return &Broker{}, errors.New("no listeners set")
	}
	listeners := strings.Split(strings.ReplaceAll(listenerStr, " ", ""), ",")

	var advertisedListeners []string
	advListenerStr := env.GetString("advertised.listeners")
	if advListenerStr == "" {
		advertisedListeners = listeners
	} else {
		advertisedListeners = strings.Split(strings.ReplaceAll(advListenerStr, " ", ""), ",")
	}

	listenersArray, err := parseListeners(listeners, false)
	if err != nil {
		return &Broker{}, err
	}
	broker.Listeners = append(broker.Listeners, listenersArray...)

	err = validateListeners(&broker)
	if err != nil {
		return &Broker{}, err
	}

	advertisedListenersArr, err := parseListeners(advertisedListeners, true)
	if err != nil {
		return &Broker{}, err
	}
	broker.AdvertisedListeners = append(broker.AdvertisedListeners, advertisedListenersArr...)

	err = validateAdvertisedListeners(&broker)
	if err != nil {
		return &Broker{}, err
	}

	brokerId := env.GetInt("broker.id")
	reservedBrokerMaxId := env.GetInt("reserved.max.broker.id")

	// validate Broker ID
	if brokerId > reservedBrokerMaxId {
		return &Broker{}, fmt.Errorf("the configured node ID is greater than `reserved.broker.max.id`. Please adjust the `reserved.broker.max.id` setting. [%d > %d]",
			brokerId,
			reservedBrokerMaxId)
	}

	if brokerId == -1 {
		brokerId = reservedBrokerMaxId + 1
	}

	broker.BrokerID = int32(brokerId)

	if len(broker.Listeners) > 1 {
		return &Broker{}, errors.New("OpenTalaria does not support more than one listener for now. See https://github.com/IBM/opentalaria/issues/18")
	}

	return &broker, nil
}

func parseListeners(listeners []string, advertised bool) ([]Listener, error) {
	result := []Listener{}

	for _, l := range listeners {
		if l == "" {
			continue
		}

		listener, err := parseListener(l, advertised)
		if err != nil {
			return []Listener{}, err
		}

		result = append(result, listener)
	}

	return result, nil
}

func parseListener(l string, advertised bool) (Listener, error) {
	listener, err := url.Parse(l)
	if err != nil {
		return Listener{}, err
	}

	// parse the security protocol from the url scheme.
	// If the protocol is unknown treat the scheme as broker name and check the listener.security.protocol.map
	listenerName, securityProtocol, err := getBrokerNameComponents(listener.Scheme)
	if err != nil {
		return Listener{}, err
	}

	host, port, err := net.SplitHostPort(listener.Host)
	if err != nil {
		return Listener{}, err
	}

	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		return Listener{}, err
	}

	// The empty host was most likely inherited from the listeners variable.
	// Since it's not allowed to advertise an empty host, we will get the IPv4 address of the first network interface.
	if advertised && host == "" {
		ifaces, err := net.Interfaces()
		if err != nil {
			return Listener{}, err
		}

		for _, iface := range ifaces {
			if (iface.Flags&net.FlagUp) != 0 && (iface.Flags&net.FlagLoopback) == 0 {
				addrs, err := iface.Addrs()
				if err != nil {
					return Listener{}, err
				}

				for _, addr := range addrs {
					ipnet, ok := addr.(*net.IPNet)
					if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
						host = ipnet.IP.To4().String()
						slog.Info("Advertised listeners not set, and listener host is empty. Setting first network iface IP as a listener", "IP", host)

						break
					}
				}

				// we found a host, break the interfaces loop
				if host != "" {
					break
				}
			}
		}
	}

	return Listener{
		Host:             host,
		Port:             int32(parsedPort),
		SecurityProtocol: securityProtocol,
		ListenerName:     listenerName,
	}, nil
}

// getBrokerNameComponents checks if the broker name, inferred from the URL schema is a valid security protocol.
// If not, it checks the listener.security.protocol.map for mapping for custom broker names and returns the broker name/security protocol pair.
// If no mapping is found in the case of custom broker name, the function returns an error.
func getBrokerNameComponents(s string) (string, SecurityProtocol, error) {
	securityProtocol, ok := ParseSecurityProtocol(s)

	if ok {
		return s, securityProtocol, nil
	} else {
		// the listener schema is not a known security protocol, treat is as broker name
		// and extract the security protocol from listener.security.protocol.map
		listenerSpmStr, _ := utils.GetEnvVar("listener.security.protocol.map", "")
		spm := strings.Split(strings.ReplaceAll(listenerSpmStr, " ", ""), ",")

		for _, sp := range spm {
			components := strings.Split(sp, ":")

			if strings.EqualFold(s, components[0]) {
				securityProtocol, ok := ParseSecurityProtocol(components[1])
				if !ok {
					return "", UNDEFINED_SECURITY_PROTOCOL, fmt.Errorf("unknown security protocol for listener %s", components[0])
				}

				return s, securityProtocol, nil
			}
		}
	}

	return "", UNDEFINED_SECURITY_PROTOCOL, fmt.Errorf("broker %s not found in listener.security.protocol.map", s)
}

// validateListeners performs common checks on the listeners as per Kafka specification https://kafka.apache.org/documentation/#brokerconfigs_listeners.
// Broker name and port have to be unique. The exception is if the host for two entries is IPv4 and IPv6 respectively.
func validateListeners(b *Broker) error {
	ports := map[int32]string{}
	listenerNames := map[string]string{}

	for _, listener := range b.Listeners {
		// Check uniqueness for ports
		if val, ok := ports[listener.Port]; ok {
			if areIpProtocolsSame(listener.Host, val) {
				return fmt.Errorf("listener port is not unique for listener %s", listener.ListenerName)
			}
		}

		// Check uniqueness for broker names
		if val, ok := listenerNames[listener.ListenerName]; ok {
			if areIpProtocolsSame(listener.Host, val) {
				return fmt.Errorf("listener name is not unique for listener %s", listener.ListenerName)
			}
		}

		ports[listener.Port] = listener.Host
		listenerNames[listener.ListenerName] = listener.Host
	}

	return nil
}

func areIpProtocolsSame(host1, host2 string) bool {
	// ignore errors from ParseAddr, which will be thrown if a hostname is provided, we care only about IP addresses.
	addr1, _ := netip.ParseAddr(host1)
	existingAddrIPVer := addr1.Is4()

	addr2, _ := netip.ParseAddr(host2)
	newAddrIPVer := addr2.Is4()

	return existingAddrIPVer == newAddrIPVer
}

// validateAdvertisedListeners performs common checks on the advertised listers as per Kafka specification https://kafka.apache.org/documentation/#brokerconfigs_advertised.listeners.
// Unlike with listeners, having duplicated ports is allowed. The only constraint is advertising to 0.0.0.0 is not allowed.
func validateAdvertisedListeners(b *Broker) error {
	for _, listener := range b.AdvertisedListeners {
		if strings.EqualFold(listener.Host, "0.0.0.0") || listener.Host == "" {
			return fmt.Errorf("advertising listener on 0.0.0.0 address is not allowed for listener %s", listener.ListenerName)
		}
	}

	return nil
}

/**
 * Unit test helpers
 */

// MockBroker generates a mock object used for unit testing.
func MockBroker() *Broker {
	broker := Broker{}

	broker.BrokerID = 1
	broker.Rack = nil
	broker.Listeners = append(broker.Listeners, Listener{
		Host:             "localhost",
		Port:             9092,
		SecurityProtocol: PLAINTEXT,
		ListenerName:     "PLAINTEXT",
	})

	broker.AdvertisedListeners = append(broker.AdvertisedListeners, Listener{
		Host:             "127.0.0.1",
		Port:             9092,
		SecurityProtocol: PLAINTEXT,
		ListenerName:     "PLAINTEXT",
	})

	return &broker
}
