// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// Listener_AddRaftVoterRequest contains the endpoints that can be used to communicate with the voter.
type Listener_AddRaftVoterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the hostname.
	Host string
	// Port contains the port.
	Port uint16
}

func (l *Listener_AddRaftVoterRequest) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if err := pe.putString(l.Name); err != nil {
		return err
	}

	if err := pe.putString(l.Host); err != nil {
		return err
	}

	pe.putUint16(l.Port)

	pe.putUVarint(0)
	return nil
}

func (l *Listener_AddRaftVoterRequest) decode(pd packetDecoder, version int16) (err error) {
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

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type AddRaftVoterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the cluster id.
	ClusterID *string
	// TimeoutMs contains the maximum time to wait for the request to complete before returning.
	TimeoutMs int32
	// VoterID contains the replica id of the voter getting added to the topic partition.
	VoterID int32
	// VoterDirectoryID contains the directory id of the voter getting added to the topic partition.
	VoterDirectoryID uuid.UUID
	// Listeners contains the endpoints that can be used to communicate with the voter.
	Listeners []Listener_AddRaftVoterRequest
}

func (r *AddRaftVoterRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.ClusterID); err != nil {
		return err
	}

	pe.putInt32(r.TimeoutMs)

	pe.putInt32(r.VoterID)

	if err := pe.putUUID(r.VoterDirectoryID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Listeners)); err != nil {
		return err
	}
	for _, block := range r.Listeners {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *AddRaftVoterRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.VoterID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.VoterDirectoryID, err = pd.getUUID(); err != nil {
		return err
	}

	var numListeners int
	if numListeners, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numListeners > 0 {
		r.Listeners = make([]Listener_AddRaftVoterRequest, numListeners)
		for i := 0; i < numListeners; i++ {
			var block Listener_AddRaftVoterRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Listeners[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *AddRaftVoterRequest) GetKey() int16 {
	return 80
}

func (r *AddRaftVoterRequest) GetVersion() int16 {
	return r.Version
}

func (r *AddRaftVoterRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *AddRaftVoterRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AddRaftVoterRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
