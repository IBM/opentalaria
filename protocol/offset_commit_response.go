// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// OffsetCommitResponsePartition contains the responses for each partition in the topic.
type OffsetCommitResponsePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (p *OffsetCommitResponsePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if p.Version >= 8 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *OffsetCommitResponsePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Version >= 8 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetCommitResponseTopic contains the responses for each topic.
type OffsetCommitResponseTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the responses for each partition in the topic.
	Partitions []OffsetCommitResponsePartition
}

func (t *OffsetCommitResponseTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
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

	if t.Version >= 8 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *OffsetCommitResponseTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]OffsetCommitResponsePartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block OffsetCommitResponsePartition
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 8 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type OffsetCommitResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Topics contains the responses for each topic.
	Topics []OffsetCommitResponseTopic
}

func (r *OffsetCommitResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 8 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 8 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *OffsetCommitResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 8 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]OffsetCommitResponseTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block OffsetCommitResponseTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 8 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *OffsetCommitResponse) GetKey() int16 {
	return 8
}

func (r *OffsetCommitResponse) GetVersion() int16 {
	return r.Version
}

func (r *OffsetCommitResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *OffsetCommitResponse) GetHeaderVersion() int16 {
	if r.Version >= 8 {
		return 1
	}
	return 0
}

func (r *OffsetCommitResponse) IsValidVersion() bool {
	return r.Version >= 2 && r.Version <= 9
}

func (r *OffsetCommitResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *OffsetCommitResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
