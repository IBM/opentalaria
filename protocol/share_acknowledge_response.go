// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// LeaderIdAndEpoch_ShareAcknowledgeResponse contains the current leader of the partition.
type LeaderIdAndEpoch_ShareAcknowledgeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch.
	LeaderEpoch int32
}

func (c *LeaderIdAndEpoch_ShareAcknowledgeResponse) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	pe.putInt32(c.LeaderID)

	pe.putInt32(c.LeaderEpoch)

	pe.putUVarint(0)
	return nil
}

func (c *LeaderIdAndEpoch_ShareAcknowledgeResponse) decode(pd packetDecoder, version int16) (err error) {
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

// PartitionData_ShareAcknowledgeResponse contains the topic partitions.
type PartitionData_ShareAcknowledgeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// CurrentLeader contains the current leader of the partition.
	CurrentLeader LeaderIdAndEpoch_ShareAcknowledgeResponse
}

func (p *PartitionData_ShareAcknowledgeResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	if err := p.CurrentLeader.encode(pe, p.Version); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_ShareAcknowledgeResponse) decode(pd packetDecoder, version int16) (err error) {
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

	tmpCurrentLeader := LeaderIdAndEpoch_ShareAcknowledgeResponse{}
	if err := tmpCurrentLeader.decode(pd, p.Version); err != nil {
		return err
	}
	p.CurrentLeader = tmpCurrentLeader

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// ShareAcknowledgeTopicResponse contains the response topics.
type ShareAcknowledgeTopicResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the unique topic ID.
	TopicID uuid.UUID
	// Partitions contains the topic partitions.
	Partitions []PartitionData_ShareAcknowledgeResponse
}

func (r *ShareAcknowledgeTopicResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (r *ShareAcknowledgeTopicResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		r.Partitions = make([]PartitionData_ShareAcknowledgeResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_ShareAcknowledgeResponse
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

// NodeEndpoint_ShareAcknowledgeResponse contains a Endpoints for all current leaders enumerated in PartitionData with error NOT_LEADER_OR_FOLLOWER.
type NodeEndpoint_ShareAcknowledgeResponse struct {
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

func (n *NodeEndpoint_ShareAcknowledgeResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (n *NodeEndpoint_ShareAcknowledgeResponse) decode(pd packetDecoder, version int16) (err error) {
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

type ShareAcknowledgeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code.
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// Responses contains the response topics.
	Responses []ShareAcknowledgeTopicResponse
	// NodeEndpoints contains a Endpoints for all current leaders enumerated in PartitionData with error NOT_LEADER_OR_FOLLOWER.
	NodeEndpoints []NodeEndpoint_ShareAcknowledgeResponse
}

func (r *ShareAcknowledgeResponse) encode(pe packetEncoder) (err error) {
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

func (r *ShareAcknowledgeResponse) decode(pd packetDecoder, version int16) (err error) {
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
		r.Responses = make([]ShareAcknowledgeTopicResponse, numResponses)
		for i := 0; i < numResponses; i++ {
			var block ShareAcknowledgeTopicResponse
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
		r.NodeEndpoints = make([]NodeEndpoint_ShareAcknowledgeResponse, numNodeEndpoints)
		for i := 0; i < numNodeEndpoints; i++ {
			var block NodeEndpoint_ShareAcknowledgeResponse
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

func (r *ShareAcknowledgeResponse) GetKey() int16 {
	return 79
}

func (r *ShareAcknowledgeResponse) GetVersion() int16 {
	return r.Version
}

func (r *ShareAcknowledgeResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *ShareAcknowledgeResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ShareAcknowledgeResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ShareAcknowledgeResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ShareAcknowledgeResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
