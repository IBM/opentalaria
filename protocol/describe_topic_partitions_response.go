// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// DescribeTopicPartitionsResponsePartition contains each partition in the topic.
type DescribeTopicPartitionsResponsePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the partition error, or 0 if there was no error.
	ErrorCode int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// LeaderID contains the ID of the leader broker.
	LeaderID int32
	// LeaderEpoch contains the leader epoch of this partition.
	LeaderEpoch int32
	// ReplicaNodes contains the set of all nodes that host this partition.
	ReplicaNodes []int32
	// IsrNodes contains the set of nodes that are in sync with the leader for this partition.
	IsrNodes []int32
	// EligibleLeaderReplicas contains the new eligible leader replicas otherwise.
	EligibleLeaderReplicas []int32
	// LastKnownElr contains the last known ELR.
	LastKnownElr []int32
	// OfflineReplicas contains the set of offline replicas of this partition.
	OfflineReplicas []int32
}

func (p *DescribeTopicPartitionsResponsePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	if err := pe.putInt32Array(p.ReplicaNodes); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.IsrNodes); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.EligibleLeaderReplicas); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.LastKnownElr); err != nil {
		return err
	}

	if err := pe.putInt32Array(p.OfflineReplicas); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *DescribeTopicPartitionsResponsePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ReplicaNodes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.IsrNodes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.EligibleLeaderReplicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.LastKnownElr, err = pd.getInt32Array(); err != nil {
		return err
	}

	if p.OfflineReplicas, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// DescribeTopicPartitionsResponseTopic contains each topic in the response.
type DescribeTopicPartitionsResponseTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the topic error, or 0 if there was no error.
	ErrorCode int16
	// Name contains the topic name.
	Name *string
	// TopicID contains the topic id.
	TopicID uuid.UUID
	// IsInternal contains a True if the topic is internal.
	IsInternal bool
	// Partitions contains each partition in the topic.
	Partitions []DescribeTopicPartitionsResponsePartition
	// TopicAuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this topic.
	TopicAuthorizedOperations int32
}

func (t *DescribeTopicPartitionsResponseTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	pe.putInt16(t.ErrorCode)

	if err := pe.putNullableString(t.Name); err != nil {
		return err
	}

	if err := pe.putUUID(t.TopicID); err != nil {
		return err
	}

	pe.putBool(t.IsInternal)

	if err := pe.putArrayLength(len(t.Partitions)); err != nil {
		return err
	}
	for _, block := range t.Partitions {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	pe.putInt32(t.TopicAuthorizedOperations)

	pe.putUVarint(0)
	return nil
}

func (t *DescribeTopicPartitionsResponseTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if t.Name, err = pd.getNullableString(); err != nil {
		return err
	}

	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	if t.IsInternal, err = pd.getBool(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]DescribeTopicPartitionsResponsePartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block DescribeTopicPartitionsResponsePartition
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.TopicAuthorizedOperations, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Cursor_DescribeTopicPartitionsResponse contains the next topic and partition index to fetch details for.
type Cursor_DescribeTopicPartitionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the name for the first topic to process.
	TopicName string
	// PartitionIndex contains the partition index to start with.
	PartitionIndex int32
}

func (n *Cursor_DescribeTopicPartitionsResponse) encode(pe packetEncoder, version int16) (err error) {
	n.Version = version
	if err := pe.putString(n.TopicName); err != nil {
		return err
	}

	pe.putInt32(n.PartitionIndex)

	pe.putUVarint(0)
	return nil
}

func (n *Cursor_DescribeTopicPartitionsResponse) decode(pd packetDecoder, version int16) (err error) {
	n.Version = version
	if n.TopicName, err = pd.getString(); err != nil {
		return err
	}

	if n.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeTopicPartitionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains each topic in the response.
	Topics []DescribeTopicPartitionsResponseTopic
	// NextCursor contains the next topic and partition index to fetch details for.
	NextCursor Cursor_DescribeTopicPartitionsResponse
}

func (r *DescribeTopicPartitionsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := r.NextCursor.encode(pe, r.Version); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeTopicPartitionsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]DescribeTopicPartitionsResponseTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block DescribeTopicPartitionsResponseTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	tmpNextCursor := Cursor_DescribeTopicPartitionsResponse{}
	if err := tmpNextCursor.decode(pd, r.Version); err != nil {
		return err
	}
	r.NextCursor = tmpNextCursor

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeTopicPartitionsResponse) GetKey() int16 {
	return 75
}

func (r *DescribeTopicPartitionsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeTopicPartitionsResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *DescribeTopicPartitionsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeTopicPartitionsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeTopicPartitionsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeTopicPartitionsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
