package config

import (
	"net"

	"github.com/ibm/opentalaria/protocol"
)

type Request struct {
	Header  protocol.RequestHeader
	Message []byte
	Conn    net.Conn
	Config  *Config
}

type RegisteredAPI struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

type ApiDefinition struct {
	ApiRequest  protocol.API
	HandlerFunc func(req Request, apiVersion int16, opts ...any) ([]byte, int16, error)
}
