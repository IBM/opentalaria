// protocol has been generated from message format json - DO NOT EDIT
package protocol

// TopicRequest_DescribeTopicPartitionsRequest contains the topics to fetch details for.
type TopicRequest_DescribeTopicPartitionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
}

func (t *TopicRequest_DescribeTopicPartitionsRequest) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicRequest_DescribeTopicPartitionsRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Cursor_DescribeTopicPartitionsRequest contains the first topic and partition index to fetch details for.
type Cursor_DescribeTopicPartitionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the name for the first topic to process.
	TopicName string
	// PartitionIndex contains the partition index to start with.
	PartitionIndex int32
}

func (c *Cursor_DescribeTopicPartitionsRequest) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.TopicName); err != nil {
		return err
	}

	pe.putInt32(c.PartitionIndex)

	pe.putUVarint(0)
	return nil
}

func (c *Cursor_DescribeTopicPartitionsRequest) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.TopicName, err = pd.getString(); err != nil {
		return err
	}

	if c.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeTopicPartitionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains the topics to fetch details for.
	Topics []TopicRequest_DescribeTopicPartitionsRequest
	// ResponsePartitionLimit contains the maximum number of partitions included in the response.
	ResponsePartitionLimit int32
	// Cursor contains the first topic and partition index to fetch details for.
	Cursor Cursor_DescribeTopicPartitionsRequest
}

func (r *DescribeTopicPartitionsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putInt32(r.ResponsePartitionLimit)

	if err := r.Cursor.encode(pe, r.Version); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeTopicPartitionsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicRequest_DescribeTopicPartitionsRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicRequest_DescribeTopicPartitionsRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.ResponsePartitionLimit, err = pd.getInt32(); err != nil {
		return err
	}

	tmpCursor := Cursor_DescribeTopicPartitionsRequest{}
	if err := tmpCursor.decode(pd, r.Version); err != nil {
		return err
	}
	r.Cursor = tmpCursor

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeTopicPartitionsRequest) GetKey() int16 {
	return 75
}

func (r *DescribeTopicPartitionsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeTopicPartitionsRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *DescribeTopicPartitionsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeTopicPartitionsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeTopicPartitionsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
