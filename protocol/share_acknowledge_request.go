// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// AcknowledgementBatch_ShareAcknowledgeRequest contains a Record batches to acknowledge.
type AcknowledgementBatch_ShareAcknowledgeRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// FirstOffset contains a First offset of batch of records to acknowledge.
	FirstOffset int64
	// LastOffset contains a Last offset (inclusive) of batch of records to acknowledge.
	LastOffset int64
	// AcknowledgeTypes contains a Array of acknowledge types - 0:Gap,1:Accept,2:Release,3:Reject.
	AcknowledgeTypes []int8
}

func (a *AcknowledgementBatch_ShareAcknowledgeRequest) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt64(a.FirstOffset)

	pe.putInt64(a.LastOffset)

	if err := pe.putInt8Array(a.AcknowledgeTypes); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (a *AcknowledgementBatch_ShareAcknowledgeRequest) decode(pd packetDecoder, version int16) (err error) {
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

// AcknowledgePartition contains the partitions containing records to acknowledge.
type AcknowledgePartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// AcknowledgementBatches contains a Record batches to acknowledge.
	AcknowledgementBatches []AcknowledgementBatch_ShareAcknowledgeRequest
}

func (p *AcknowledgePartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

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

func (p *AcknowledgePartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	var numAcknowledgementBatches int
	if numAcknowledgementBatches, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numAcknowledgementBatches > 0 {
		p.AcknowledgementBatches = make([]AcknowledgementBatch_ShareAcknowledgeRequest, numAcknowledgementBatches)
		for i := 0; i < numAcknowledgementBatches; i++ {
			var block AcknowledgementBatch_ShareAcknowledgeRequest
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

// AcknowledgeTopic contains the topics containing records to acknowledge.
type AcknowledgeTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// Partitions contains the partitions containing records to acknowledge.
	Partitions []AcknowledgePartition
}

func (t *AcknowledgeTopic) encode(pe packetEncoder, version int16) (err error) {
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

func (t *AcknowledgeTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]AcknowledgePartition, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block AcknowledgePartition
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

type ShareAcknowledgeRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID *string
	// MemberID contains the member ID.
	MemberID *string
	// ShareSessionEpoch contains the current share session epoch: 0 to open a share session; -1 to close it; otherwise increments for consecutive requests.
	ShareSessionEpoch int32
	// Topics contains the topics containing records to acknowledge.
	Topics []AcknowledgeTopic
}

func (r *ShareAcknowledgeRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putNullableString(r.GroupID); err != nil {
		return err
	}

	if err := pe.putNullableString(r.MemberID); err != nil {
		return err
	}

	pe.putInt32(r.ShareSessionEpoch)

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

func (r *ShareAcknowledgeRequest) decode(pd packetDecoder, version int16) (err error) {
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

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]AcknowledgeTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block AcknowledgeTopic
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

func (r *ShareAcknowledgeRequest) GetKey() int16 {
	return 79
}

func (r *ShareAcknowledgeRequest) GetVersion() int16 {
	return r.Version
}

func (r *ShareAcknowledgeRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ShareAcknowledgeRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ShareAcknowledgeRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
