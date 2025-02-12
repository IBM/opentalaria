// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// LeaderIdAndEpoch_ShareFetchResponse contains the current leader of the partition.
type LeaderIdAndEpoch_ShareFetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch.
	LeaderEpoch int32
}

func (c *LeaderIdAndEpoch_ShareFetchResponse) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	pe.putInt32(c.LeaderID)

	pe.putInt32(c.LeaderEpoch)

	pe.putUVarint(0)
	return nil
}

func (c *LeaderIdAndEpoch_ShareFetchResponse) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if c.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// AcquiredRecords contains the acquired records.
type AcquiredRecords struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// FirstOffset contains the earliest offset in this batch of acquired records.
	FirstOffset int64
	// LastOffset contains the last offset of this batch of acquired records.
	LastOffset int64
	// DeliveryCount contains the delivery count of this batch of acquired records.
	DeliveryCount int16
}

func (a *AcquiredRecords) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt64(a.FirstOffset)

	pe.putInt64(a.LastOffset)

	pe.putInt16(a.DeliveryCount)

	pe.putUVarint(0)
	return nil
}

func (a *AcquiredRecords) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.FirstOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if a.LastOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if a.DeliveryCount, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// PartitionData_ShareFetchResponse contains the topic partitions.
type PartitionData_ShareFetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the fetch error code, or 0 if there was no fetch error.
	ErrorCode int16
	// ErrorMessage contains the fetch error message, or null if there was no fetch error.
	ErrorMessage *string
	// AcknowledgeErrorCode contains the acknowledge error code, or 0 if there was no acknowledge error.
	AcknowledgeErrorCode int16
	// AcknowledgeErrorMessage contains the acknowledge error message, or null if there was no acknowledge error.
	AcknowledgeErrorMessage *string
	// CurrentLeader contains the current leader of the partition.
	CurrentLeader LeaderIdAndEpoch_ShareFetchResponse
	// Records contains the record data.
	Records RecordBatch
	// AcquiredRecords contains the acquired records.
	AcquiredRecords []AcquiredRecords
}

func (p *PartitionData_ShareFetchResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	pe.putInt16(p.AcknowledgeErrorCode)

	if err := pe.putNullableString(p.AcknowledgeErrorMessage); err != nil {
		return err
	}

	if err := p.CurrentLeader.encode(pe, p.Version); err != nil {
		return err
	}

	if err := p.Records.encode(pe, p.Version); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(p.AcquiredRecords)); err != nil {
		return err
	}
	for _, block := range p.AcquiredRecords {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_ShareFetchResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if p.AcknowledgeErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.AcknowledgeErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	tmpCurrentLeader := LeaderIdAndEpoch_ShareFetchResponse{}
	if err := tmpCurrentLeader.decode(pd, p.Version); err != nil {
		return err
	}
	p.CurrentLeader = tmpCurrentLeader

	tmpRecords := RecordBatch{}
	if err := tmpRecords.decode(pd, p.Version); err != nil {
		return err
	}
	p.Records = tmpRecords

	var numAcquiredRecords int
	if numAcquiredRecords, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numAcquiredRecords > 0 {
		p.AcquiredRecords = make([]AcquiredRecords, numAcquiredRecords)
		for i := 0; i < numAcquiredRecords; i++ {
			var block AcquiredRecords
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.AcquiredRecords[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// ShareFetchableTopicResponse contains the response topics.
type ShareFetchableTopicResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// Partitions contains the topic partitions.
	Partitions []PartitionData_ShareFetchResponse
}

func (r *ShareFetchableTopicResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (r *ShareFetchableTopicResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		r.Partitions = make([]PartitionData_ShareFetchResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_ShareFetchResponse
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

// NodeEndpoint_ShareFetchResponse contains a Endpoints for all current leaders enumerated in PartitionData with error NOT_LEADER_OR_FOLLOWER.
type NodeEndpoint_ShareFetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// NodeID contains the ID of the associated node.
	NodeID int32
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port int32
	// Rack contains the rack of the node, or null if it has not been assigned to a rack.
	Rack *string
}

func (n *NodeEndpoint_ShareFetchResponse) encode(pe packetEncoder, version int16) (err error) {
	n.Version = version
	pe.putInt32(n.NodeID)

	if err := pe.putString(n.Host); err != nil {
		return err
	}

	pe.putInt32(n.Port)

	if err := pe.putNullableString(n.Rack); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (n *NodeEndpoint_ShareFetchResponse) decode(pd packetDecoder, version int16) (err error) {
	n.Version = version
	if n.NodeID, err = pd.getInt32(); err != nil {
		return err
	}

	if n.Host, err = pd.getString(); err != nil {
		return err
	}

	if n.Port, err = pd.getInt32(); err != nil {
		return err
	}

	if n.Rack, err = pd.getNullableString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ShareFetchResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level response error code.
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// Responses contains the response topics.
	Responses []ShareFetchableTopicResponse
	// NodeEndpoints contains a Endpoints for all current leaders enumerated in PartitionData with error NOT_LEADER_OR_FOLLOWER.
	NodeEndpoints []NodeEndpoint_ShareFetchResponse
}

func (r *ShareFetchResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Responses)); err != nil {
		return err
	}
	for _, block := range r.Responses {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.NodeEndpoints)); err != nil {
		return err
	}
	for _, block := range r.NodeEndpoints {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ShareFetchResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	var numResponses int
	if numResponses, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResponses > 0 {
		r.Responses = make([]ShareFetchableTopicResponse, numResponses)
		for i := 0; i < numResponses; i++ {
			var block ShareFetchableTopicResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Responses[i] = block
		}
	}

	var numNodeEndpoints int
	if numNodeEndpoints, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numNodeEndpoints > 0 {
		r.NodeEndpoints = make([]NodeEndpoint_ShareFetchResponse, numNodeEndpoints)
		for i := 0; i < numNodeEndpoints; i++ {
			var block NodeEndpoint_ShareFetchResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.NodeEndpoints[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ShareFetchResponse) GetKey() int16 {
	return 78
}

func (r *ShareFetchResponse) GetVersion() int16 {
	return r.Version
}

func (r *ShareFetchResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ShareFetchResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ShareFetchResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ShareFetchResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
