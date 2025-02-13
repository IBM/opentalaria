// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// PartitionData_InitializeShareGroupStateRequest contains the data for the partitions.
type PartitionData_InitializeShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// StateEpoch contains the state epoch for this share-partition.
	StateEpoch int32
	// StartOffset contains the share-partition start offset, or -1 if the start offset is not being initialized.
	StartOffset int64
}

func (p *PartitionData_InitializeShareGroupStateRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	pe.putInt32(p.StateEpoch)

	pe.putInt64(p.StartOffset)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_InitializeShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.StateEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.StartOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// InitializeStateData contains the data for the topics.
type InitializeStateData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic identifier.
	TopicID uuid.UUID
	// Partitions contains the data for the partitions.
	Partitions []PartitionData_InitializeShareGroupStateRequest
}

func (t *InitializeStateData) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putUUID(t.TopicID); err != nil {
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

func (t *InitializeStateData) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_InitializeShareGroupStateRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_InitializeShareGroupStateRequest
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

type InitializeShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID string
	// Topics contains the data for the topics.
	Topics []InitializeStateData
}

func (r *InitializeShareGroupStateRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putString(r.GroupID); err != nil {
		return err
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

func (r *InitializeShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]InitializeStateData, numTopics)
		for i := 0; i < numTopics; i++ {
			var block InitializeStateData
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

func (r *InitializeShareGroupStateRequest) GetKey() int16 {
	return 83
}

func (r *InitializeShareGroupStateRequest) GetVersion() int16 {
	return r.Version
}

func (r *InitializeShareGroupStateRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *InitializeShareGroupStateRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *InitializeShareGroupStateRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
