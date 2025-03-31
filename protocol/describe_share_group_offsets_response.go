// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// DescribeShareGroupOffsetsResponsePartition contains a
type DescribeShareGroupOffsetsResponsePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// StartOffset contains the share-partition start offset.
	StartOffset int64
	// LeaderEpoch contains the leader epoch of the partition.
	LeaderEpoch int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
}

func (p *DescribeShareGroupOffsetsResponsePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt64(p.StartOffset)

	pe.putInt32(p.LeaderEpoch)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *DescribeShareGroupOffsetsResponsePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.StartOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// DescribeShareGroupOffsetsResponseTopic contains the results for each topic.
type DescribeShareGroupOffsetsResponseTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// Partitions contains a
	Partitions []DescribeShareGroupOffsetsResponsePartition
}

func (r *DescribeShareGroupOffsetsResponseTopic) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.TopicName); err != nil {
		return err
	}

	if err := pe.putUUID(r.TopicID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Partitions)); err != nil {
		return err
	}
	for _, block := range r.Partitions {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeShareGroupOffsetsResponseTopic) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicName, err = pd.getString(); err != nil {
		return err
	}

	if r.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		r.Partitions = make([]DescribeShareGroupOffsetsResponsePartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block DescribeShareGroupOffsetsResponsePartition
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Partitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeShareGroupOffsetsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Responses contains the results for each topic.
	Responses []DescribeShareGroupOffsetsResponseTopic
}

func (r *DescribeShareGroupOffsetsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeShareGroupOffsetsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResponses > 0 {
		r.Responses = make([]DescribeShareGroupOffsetsResponseTopic, numResponses)
		for i := 0; i < numResponses; i++ {
			var block DescribeShareGroupOffsetsResponseTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Responses[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeShareGroupOffsetsResponse) GetKey() int16 {
	return 90
}

func (r *DescribeShareGroupOffsetsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeShareGroupOffsetsResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *DescribeShareGroupOffsetsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeShareGroupOffsetsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeShareGroupOffsetsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeShareGroupOffsetsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
