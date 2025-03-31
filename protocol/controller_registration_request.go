// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// Listener_ControllerRegistrationRequest contains the listeners of this controller.
type Listener_ControllerRegistrationRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the hostname.
	Host string
	// Port contains the port.
	Port uint16
	// SecurityProtocol contains the security protocol.
	SecurityProtocol int16
}

func (l *Listener_ControllerRegistrationRequest) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if err := pe.putString(l.Name); err != nil {
		return err
	}

	if err := pe.putString(l.Host); err != nil {
		return err
	}

	pe.putUint16(l.Port)

	pe.putInt16(l.SecurityProtocol)

	pe.putUVarint(0)
	return nil
}

func (l *Listener_ControllerRegistrationRequest) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.Name, err = pd.getString(); err != nil {
		return err
	}

	if l.Host, err = pd.getString(); err != nil {
		return err
	}

	if l.Port, err = pd.getUint16(); err != nil {
		return err
	}

	if l.SecurityProtocol, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Feature_ControllerRegistrationRequest contains the features on this controller.
type Feature_ControllerRegistrationRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the feature name.
	Name string
	// MinSupportedVersion contains the minimum supported feature level.
	MinSupportedVersion int16
	// MaxSupportedVersion contains the maximum supported feature level.
	MaxSupportedVersion int16
}

func (f *Feature_ControllerRegistrationRequest) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	if err := pe.putString(f.Name); err != nil {
		return err
	}

	pe.putInt16(f.MinSupportedVersion)

	pe.putInt16(f.MaxSupportedVersion)

	pe.putUVarint(0)
	return nil
}

func (f *Feature_ControllerRegistrationRequest) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.Name, err = pd.getString(); err != nil {
		return err
	}

	if f.MinSupportedVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if f.MaxSupportedVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ControllerRegistrationRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ControllerID contains the ID of the controller to register.
	ControllerID int32
	// IncarnationID contains the controller incarnation ID, which is unique to each process run.
	IncarnationID uuid.UUID
	// ZkMigrationReady contains a Set if the required configurations for ZK migration are present.
	ZkMigrationReady bool
	// Listeners contains the listeners of this controller.
	Listeners []Listener_ControllerRegistrationRequest
	// Features contains the features on this controller.
	Features []Feature_ControllerRegistrationRequest
}

func (r *ControllerRegistrationRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ControllerID)

	if err := pe.putUUID(r.IncarnationID); err != nil {
		return err
	}

	pe.putBool(r.ZkMigrationReady)

	if err := pe.putArrayLength(len(r.Listeners)); err != nil {
		return err
	}
	for _, block := range r.Listeners {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Features)); err != nil {
		return err
	}
	for _, block := range r.Features {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ControllerRegistrationRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ControllerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.IncarnationID, err = pd.getUUID(); err != nil {
		return err
	}

	if r.ZkMigrationReady, err = pd.getBool(); err != nil {
		return err
	}

	var numListeners int
	if numListeners, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numListeners > 0 {
		r.Listeners = make([]Listener_ControllerRegistrationRequest, numListeners)
		for i := 0; i < numListeners; i++ {
			var block Listener_ControllerRegistrationRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Listeners[i] = block
		}
	}

	var numFeatures int
	if numFeatures, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numFeatures > 0 {
		r.Features = make([]Feature_ControllerRegistrationRequest, numFeatures)
		for i := 0; i < numFeatures; i++ {
			var block Feature_ControllerRegistrationRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Features[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ControllerRegistrationRequest) GetKey() int16 {
	return 70
}

func (r *ControllerRegistrationRequest) GetVersion() int16 {
	return r.Version
}

func (r *ControllerRegistrationRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *ControllerRegistrationRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ControllerRegistrationRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ControllerRegistrationRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
