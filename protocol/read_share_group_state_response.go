// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// StateBatch_ReadShareGroupStateResponse contains the state batches for this share-partition.
type StateBatch_ReadShareGroupStateResponse struct {
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

func (s *StateBatch_ReadShareGroupStateResponse) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	pe.putInt64(s.FirstOffset)

	pe.putInt64(s.LastOffset)

	pe.putInt8(s.DeliveryState)

	pe.putInt16(s.DeliveryCount)

	pe.putUVarint(0)
	return nil
}

func (s *StateBatch_ReadShareGroupStateResponse) decode(pd packetDecoder, version int16) (err error) {
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

// PartitionResult_ReadShareGroupStateResponse contains the results for the partitions.
type PartitionResult_ReadShareGroupStateResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// StateEpoch contains the state epoch for this share-partition.
	StateEpoch int32
	// StartOffset contains the share-partition start offset, which can be -1 if it is not yet initialized.
	StartOffset int64
	// StateBatches contains the state batches for this share-partition.
	StateBatches []StateBatch_ReadShareGroupStateResponse
}

func (p *PartitionResult_ReadShareGroupStateResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	pe.putInt32(p.StateEpoch)

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

func (p *PartitionResult_ReadShareGroupStateResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if p.StateEpoch, err = pd.getInt32(); err != nil {
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
		p.StateBatches = make([]StateBatch_ReadShareGroupStateResponse, numStateBatches)
		for i := 0; i < numStateBatches; i++ {
			var block StateBatch_ReadShareGroupStateResponse
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

// ReadStateResult contains the read results.
type ReadStateResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic identifier.
	TopicID uuid.UUID
	// Partitions contains the results for the partitions.
	Partitions []PartitionResult_ReadShareGroupStateResponse
}

func (r *ReadStateResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
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

func (r *ReadStateResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		r.Partitions = make([]PartitionResult_ReadShareGroupStateResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionResult_ReadShareGroupStateResponse
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

type ReadShareGroupStateResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Results contains the read results.
	Results []ReadStateResult
}

func (r *ReadShareGroupStateResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ReadShareGroupStateResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResults > 0 {
		r.Results = make([]ReadStateResult, numResults)
		for i := 0; i < numResults; i++ {
			var block ReadStateResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Results[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ReadShareGroupStateResponse) GetKey() int16 {
	return 84
}

func (r *ReadShareGroupStateResponse) GetVersion() int16 {
	return r.Version
}

func (r *ReadShareGroupStateResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *ReadShareGroupStateResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ReadShareGroupStateResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ReadShareGroupStateResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
