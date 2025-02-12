// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

type RemoveRaftVoterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the cluster id of the request.
	ClusterID *string
	// VoterID contains the replica id of the voter getting removed from the topic partition.
	VoterID int32
	// VoterDirectoryID contains the directory id of the voter getting removed from the topic partition.
	VoterDirectoryID uuid.UUID
}

func (r *RemoveRaftVoterRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.ClusterID); err != nil {
		return err
	}

	pe.putInt32(r.VoterID)

	if err := pe.putUUID(r.VoterDirectoryID); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *RemoveRaftVoterRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.VoterID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.VoterDirectoryID, err = pd.getUUID(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *RemoveRaftVoterRequest) GetKey() int16 {
	return 81
}

func (r *RemoveRaftVoterRequest) GetVersion() int16 {
	return r.Version
}

func (r *RemoveRaftVoterRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *RemoveRaftVoterRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *RemoveRaftVoterRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
