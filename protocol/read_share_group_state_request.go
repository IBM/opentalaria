// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// PartitionData_ReadShareGroupStateRequest contains the data for the partitions.
type PartitionData_ReadShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// LeaderEpoch contains the leader epoch of the share-partition.
	LeaderEpoch int32
}

func (p *PartitionData_ReadShareGroupStateRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	pe.putInt32(p.LeaderEpoch)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_ReadShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// ReadStateData contains the data for the topics.
type ReadStateData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic identifier.
	TopicID uuid.UUID
	// Partitions contains the data for the partitions.
	Partitions []PartitionData_ReadShareGroupStateRequest
}

func (t *ReadStateData) encode(pe packetEncoder, version int16) (err error) {
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

func (t *ReadStateData) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_ReadShareGroupStateRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_ReadShareGroupStateRequest
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

type ReadShareGroupStateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID string
	// Topics contains the data for the topics.
	Topics []ReadStateData
}

func (r *ReadShareGroupStateRequest) encode(pe packetEncoder) (err error) {
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

func (r *ReadShareGroupStateRequest) decode(pd packetDecoder, version int16) (err error) {
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
		r.Topics = make([]ReadStateData, numTopics)
		for i := 0; i < numTopics; i++ {
			var block ReadStateData
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

func (r *ReadShareGroupStateRequest) GetKey() int16 {
	return 84
}

func (r *ReadShareGroupStateRequest) GetVersion() int16 {
	return r.Version
}

func (r *ReadShareGroupStateRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *ReadShareGroupStateRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ReadShareGroupStateRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ReadShareGroupStateRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
