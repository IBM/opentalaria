// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// Listener_UpdateRaftVoterRequest contains the endpoint that can be used to communicate with the leader.
type Listener_UpdateRaftVoterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the hostname.
	Host string
	// Port contains the port.
	Port uint16
}

func (l *Listener_UpdateRaftVoterRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (l *Listener_UpdateRaftVoterRequest) decode(pd packetDecoder, version int16) (err error) {
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

// KRaftVersionFeature_UpdateRaftVoterRequest contains the range of versions of the protocol that the replica supports.
type KRaftVersionFeature_UpdateRaftVoterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MinSupportedVersion contains the minimum supported KRaft protocol version.
	MinSupportedVersion int16
	// MaxSupportedVersion contains the maximum supported KRaft protocol version.
	MaxSupportedVersion int16
}

func (k *KRaftVersionFeature_UpdateRaftVoterRequest) encode(pe packetEncoder, version int16) (err error) {
	k.Version = version
	pe.putInt16(k.MinSupportedVersion)

	pe.putInt16(k.MaxSupportedVersion)

	pe.putUVarint(0)
	return nil
}

func (k *KRaftVersionFeature_UpdateRaftVoterRequest) decode(pd packetDecoder, version int16) (err error) {
	k.Version = version
	if k.MinSupportedVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if k.MaxSupportedVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type UpdateRaftVoterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the cluster id.
	ClusterID *string
	// CurrentLeaderEpoch contains the current leader epoch of the partition, -1 for unknown leader epoch.
	CurrentLeaderEpoch int32
	// VoterID contains the replica id of the voter getting updated in the topic partition.
	VoterID int32
	// VoterDirectoryID contains the directory id of the voter getting updated in the topic partition.
	VoterDirectoryID uuid.UUID
	// Listeners contains the endpoint that can be used to communicate with the leader.
	Listeners []Listener_UpdateRaftVoterRequest
	// KRaftVersionFeature contains the range of versions of the protocol that the replica supports.
	KRaftVersionFeature KRaftVersionFeature_UpdateRaftVoterRequest
}

func (r *UpdateRaftVoterRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.ClusterID); err != nil {
		return err
	}

	pe.putInt32(r.CurrentLeaderEpoch)

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

	if err := r.KRaftVersionFeature.encode(pe, r.Version); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *UpdateRaftVoterRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.CurrentLeaderEpoch, err = pd.getInt32(); err != nil {
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
		r.Listeners = make([]Listener_UpdateRaftVoterRequest, numListeners)
		for i := 0; i < numListeners; i++ {
			var block Listener_UpdateRaftVoterRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Listeners[i] = block
		}
	}

	tmpKRaftVersionFeature := KRaftVersionFeature_UpdateRaftVoterRequest{}
	if err := tmpKRaftVersionFeature.decode(pd, r.Version); err != nil {
		return err
	}
	r.KRaftVersionFeature = tmpKRaftVersionFeature

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *UpdateRaftVoterRequest) GetKey() int16 {
	return 82
}

func (r *UpdateRaftVoterRequest) GetVersion() int16 {
	return r.Version
}

func (r *UpdateRaftVoterRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *UpdateRaftVoterRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *UpdateRaftVoterRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
