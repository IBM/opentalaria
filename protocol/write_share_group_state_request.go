// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// StateBatch_WriteShareGroupStateRequest contains the state batches for the share-partition.
type StateBatch_WriteShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// FirstOffset contains the base offset of this state batch.
	FirstOffset int64
	// LastOffset contains the last offset of this state batch.
	LastOffset int64
	// DeliveryState contains the state - 0:Available,2:Acked,4:Archived.
	DeliveryState int8
	// DeliveryCount contains the delivery count.
	DeliveryCount int16
}

func (s *StateBatch_WriteShareGroupStateRequest) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	pe.putInt64(s.FirstOffset)

	pe.putInt64(s.LastOffset)

	pe.putInt8(s.DeliveryState)

	pe.putInt16(s.DeliveryCount)

	pe.putUVarint(0)
	return nil
}

func (s *StateBatch_WriteShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.FirstOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if s.LastOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if s.DeliveryState, err = pd.getInt8(); err != nil {
		return err
	}

	if s.DeliveryCount, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// PartitionData_WriteShareGroupStateRequest contains the data for the partitions.
type PartitionData_WriteShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// StateEpoch contains the state epoch for this share-partition.
	StateEpoch int32
	// LeaderEpoch contains the leader epoch of the share-partition.
	LeaderEpoch int32
	// StartOffset contains the share-partition start offset, or -1 if the start offset is not being written.
	StartOffset int64
	// StateBatches contains the state batches for the share-partition.
	StateBatches []StateBatch_WriteShareGroupStateRequest
}

func (p *PartitionData_WriteShareGroupStateRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	pe.putInt32(p.StateEpoch)

	pe.putInt32(p.LeaderEpoch)

	pe.putInt64(p.StartOffset)

	if err := pe.putArrayLength(len(p.StateBatches)); err != nil {
		return err
	}
	for _, block := range p.StateBatches {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_WriteShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.StateEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.StartOffset, err = pd.getInt64(); err != nil {
		return err
	}

	var numStateBatches int
	if numStateBatches, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numStateBatches > 0 {
		p.StateBatches = make([]StateBatch_WriteShareGroupStateRequest, numStateBatches)
		for i := 0; i < numStateBatches; i++ {
			var block StateBatch_WriteShareGroupStateRequest
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.StateBatches[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// WriteStateData contains the data for the topics.
type WriteStateData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic identifier.
	TopicID uuid.UUID
	// Partitions contains the data for the partitions.
	Partitions []PartitionData_WriteShareGroupStateRequest
}

func (t *WriteStateData) encode(pe packetEncoder, version int16) (err error) {
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

func (t *WriteStateData) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_WriteShareGroupStateRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_WriteShareGroupStateRequest
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

type WriteShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID string
	// Topics contains the data for the topics.
	Topics []WriteStateData
}

func (r *WriteShareGroupStateRequest) encode(pe packetEncoder) (err error) {
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

func (r *WriteShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
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
		r.Topics = make([]WriteStateData, numTopics)
		for i := 0; i < numTopics; i++ {
			var block WriteStateData
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

func (r *WriteShareGroupStateRequest) GetKey() int16 {
	return 85
}

func (r *WriteShareGroupStateRequest) GetVersion() int16 {
	return r.Version
}

func (r *WriteShareGroupStateRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *WriteShareGroupStateRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *WriteShareGroupStateRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *WriteShareGroupStateRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
