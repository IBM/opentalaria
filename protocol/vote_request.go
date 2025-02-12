// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// PartitionData_VoteRequest contains the partition data.
type PartitionData_VoteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ReplicaEpoch contains the epoch of the voter sending the request
	ReplicaEpoch int32
	// ReplicaID contains the replica id of the voter sending the request
	ReplicaID int32
	// ReplicaDirectoryID contains the directory id of the voter sending the request
	ReplicaDirectoryID uuid.UUID
	// VoterDirectoryID contains the directory id of the voter receiving the request
	VoterDirectoryID uuid.UUID
	// LastOffsetEpoch contains the epoch of the last record written to the metadata log.
	LastOffsetEpoch int32
	// LastOffset contains the log end offset of the metadata log of the voter sending the request.
	LastOffset int64
	// PreVote contains a Whether the request is a PreVote request (not persisted) or not.
	PreVote bool
}

func (p *PartitionData_VoteRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.ReplicaEpoch)

	pe.putInt32(p.ReplicaID)

	if p.Version >= 1 {
		if err := pe.putUUID(p.ReplicaDirectoryID); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if err := pe.putUUID(p.VoterDirectoryID); err != nil {
			return err
		}
	}

	pe.putInt32(p.LastOffsetEpoch)

	pe.putInt64(p.LastOffset)

	if p.Version >= 2 {
		pe.putBool(p.PreVote)
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_VoteRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ReplicaEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ReplicaID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 1 {
		if p.ReplicaDirectoryID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if p.VoterDirectoryID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if p.LastOffsetEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LastOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if p.PreVote, err = pd.getBool(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_VoteRequest contains the topic data.
type TopicData_VoteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partition data.
	Partitions []PartitionData_VoteRequest
}

func (t *TopicData_VoteRequest) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.TopicName); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(t.Partitions)); err != nil {
		return err
	}
	for _, block := range t.Partitions {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicData_VoteRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_VoteRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_VoteRequest
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type VoteRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the cluster id.
	ClusterID *string
	// VoterID contains the replica id of the voter receiving the request.
	VoterID int32
	// Topics contains the topic data.
	Topics []TopicData_VoteRequest
}

func (r *VoteRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.ClusterID); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putInt32(r.VoterID)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *VoteRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.VoterID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_VoteRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_VoteRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *VoteRequest) GetKey() int16 {
	return 52
}

func (r *VoteRequest) GetVersion() int16 {
	return r.Version
}

func (r *VoteRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *VoteRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *VoteRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
