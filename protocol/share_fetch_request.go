// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// AcknowledgementBatch_ShareFetchRequest contains a Record batches to acknowledge.
type AcknowledgementBatch_ShareFetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// FirstOffset contains a First offset of batch of records to acknowledge.
	FirstOffset int64
	// LastOffset contains a Last offset (inclusive) of batch of records to acknowledge.
	LastOffset int64
	// AcknowledgeTypes contains a Array of acknowledge types - 0:Gap,1:Accept,2:Release,3:Reject.
	AcknowledgeTypes []int8
}

func (a *AcknowledgementBatch_ShareFetchRequest) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt64(a.FirstOffset)

	pe.putInt64(a.LastOffset)

	if err := pe.putInt8Array(a.AcknowledgeTypes); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (a *AcknowledgementBatch_ShareFetchRequest) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.FirstOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if a.LastOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if a.AcknowledgeTypes, err = pd.getInt8Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// FetchPartition_ShareFetchRequest contains the partitions to fetch.
type FetchPartition_ShareFetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// PartitionMaxBytes contains the maximum bytes to fetch from this partition. 0 when only acknowledgement with no fetching is required. See KIP-74 for cases where this limit may not be honored.
	PartitionMaxBytes int32
	// AcknowledgementBatches contains a Record batches to acknowledge.
	AcknowledgementBatches []AcknowledgementBatch_ShareFetchRequest
}

func (p *FetchPartition_ShareFetchRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.PartitionMaxBytes)

	if err := pe.putArrayLength(len(p.AcknowledgementBatches)); err != nil {
		return err
	}
	for _, block := range p.AcknowledgementBatches {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (p *FetchPartition_ShareFetchRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.PartitionMaxBytes, err = pd.getInt32(); err != nil {
		return err
	}

	var numAcknowledgementBatches int
	if numAcknowledgementBatches, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numAcknowledgementBatches > 0 {
		p.AcknowledgementBatches = make([]AcknowledgementBatch_ShareFetchRequest, numAcknowledgementBatches)
		for i := 0; i < numAcknowledgementBatches; i++ {
			var block AcknowledgementBatch_ShareFetchRequest
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.AcknowledgementBatches[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// FetchTopic_ShareFetchRequest contains the topics to fetch.
type FetchTopic_ShareFetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// Partitions contains the partitions to fetch.
	Partitions []FetchPartition_ShareFetchRequest
}

func (t *FetchTopic_ShareFetchRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *FetchTopic_ShareFetchRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]FetchPartition_ShareFetchRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block FetchPartition_ShareFetchRequest
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

// ForgottenTopic_ShareFetchRequest contains the partitions to remove from this share session.
type ForgottenTopic_ShareFetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// Partitions contains the partitions indexes to forget.
	Partitions []int32
}

func (f *ForgottenTopic_ShareFetchRequest) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	if err := pe.putUUID(f.TopicID); err != nil {
		return err
	}

	if err := pe.putInt32Array(f.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (f *ForgottenTopic_ShareFetchRequest) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	if f.Partitions, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ShareFetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID *string
	// MemberID contains the member ID.
	MemberID *string
	// ShareSessionEpoch contains the current share session epoch: 0 to open a share session; -1 to close it; otherwise increments for consecutive requests.
	ShareSessionEpoch int32
	// MaxWaitMs contains the maximum time in milliseconds to wait for the response.
	MaxWaitMs int32
	// MinBytes contains the minimum bytes to accumulate in the response.
	MinBytes int32
	// MaxBytes contains the maximum bytes to fetch.  See KIP-74 for cases where this limit may not be honored.
	MaxBytes int32
	// BatchSize contains the optimal number of records for batches of acquired records and acknowledgements.
	BatchSize int32
	// Topics contains the topics to fetch.
	Topics []FetchTopic_ShareFetchRequest
	// ForgottenTopicsData contains the partitions to remove from this share session.
	ForgottenTopicsData []ForgottenTopic_ShareFetchRequest
}

func (r *ShareFetchRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.GroupID); err != nil {
		return err
	}

	if err := pe.putNullableString(r.MemberID); err != nil {
		return err
	}

	pe.putInt32(r.ShareSessionEpoch)

	pe.putInt32(r.MaxWaitMs)

	pe.putInt32(r.MinBytes)

	pe.putInt32(r.MaxBytes)

	pe.putInt32(r.BatchSize)

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.ForgottenTopicsData)); err != nil {
		return err
	}
	for _, block := range r.ForgottenTopicsData {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ShareFetchRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.GroupID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.MemberID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.ShareSessionEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.MaxWaitMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.MinBytes, err = pd.getInt32(); err != nil {
		return err
	}

	if r.MaxBytes, err = pd.getInt32(); err != nil {
		return err
	}

	if r.BatchSize, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]FetchTopic_ShareFetchRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block FetchTopic_ShareFetchRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	var numForgottenTopicsData int
	if numForgottenTopicsData, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numForgottenTopicsData > 0 {
		r.ForgottenTopicsData = make([]ForgottenTopic_ShareFetchRequest, numForgottenTopicsData)
		for i := 0; i < numForgottenTopicsData; i++ {
			var block ForgottenTopic_ShareFetchRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.ForgottenTopicsData[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ShareFetchRequest) GetKey() int16 {
	return 78
}

func (r *ShareFetchRequest) GetVersion() int16 {
	return r.Version
}

func (r *ShareFetchRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *ShareFetchRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ShareFetchRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ShareFetchRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
