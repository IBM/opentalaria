// protocol has been generated from message format json - DO NOT EDIT
package protocol

// DescribeShareGroupOffsetsRequestTopic contains the topics to describe offsets for.
type DescribeShareGroupOffsetsRequestTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partitions.
	Partitions []int32
}

func (t *DescribeShareGroupOffsetsRequestTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.TopicName); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *DescribeShareGroupOffsetsRequestTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	if t.Partitions, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeShareGroupOffsetsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID string
	// Topics contains the topics to describe offsets for.
	Topics []DescribeShareGroupOffsetsRequestTopic
}

func (r *DescribeShareGroupOffsetsRequest) encode(pe packetEncoder) (err error) {
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

func (r *DescribeShareGroupOffsetsRequest) decode(pd packetDecoder, version int16) (err error) {
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
		r.Topics = make([]DescribeShareGroupOffsetsRequestTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block DescribeShareGroupOffsetsRequestTopic
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

func (r *DescribeShareGroupOffsetsRequest) GetKey() int16 {
	return 90
}

func (r *DescribeShareGroupOffsetsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeShareGroupOffsetsRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *DescribeShareGroupOffsetsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeShareGroupOffsetsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeShareGroupOffsetsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
