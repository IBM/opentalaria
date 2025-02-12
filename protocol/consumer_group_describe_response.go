// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// TopicPartitions_ConsumerGroupDescribeResponse contains a
type TopicPartitions_ConsumerGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic ID.
	TopicID uuid.UUID
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partitions.
	Partitions []int32
}

func (t *TopicPartitions_ConsumerGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putUUID(t.TopicID); err != nil {
		return err
	}

	if err := pe.putString(t.TopicName); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicPartitions_ConsumerGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

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

// Assignment_ConsumerGroupDescribeResponse contains a
type Assignment_ConsumerGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicPartitions contains the assigned topic-partitions to the member.
	TopicPartitions []TopicPartitions_ConsumerGroupDescribeResponse
}

func (a *Assignment_ConsumerGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putArrayLength(len(a.TopicPartitions)); err != nil {
		return err
	}
	for _, block := range a.TopicPartitions {
		if err := block.encode(pe, a.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (a *Assignment_ConsumerGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	var numTopicPartitions int
	if numTopicPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopicPartitions > 0 {
		a.TopicPartitions = make([]TopicPartitions_ConsumerGroupDescribeResponse, numTopicPartitions)
		for i := 0; i < numTopicPartitions; i++ {
			var block TopicPartitions_ConsumerGroupDescribeResponse
			if err := block.decode(pd, a.Version); err != nil {
				return err
			}
			a.TopicPartitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Member_ConsumerGroupDescribeResponse contains the members.
type Member_ConsumerGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the member ID.
	MemberID string
	// InstanceID contains the member instance ID.
	InstanceID *string
	// RackID contains the member rack ID.
	RackID *string
	// MemberEpoch contains the current member epoch.
	MemberEpoch int32
	// ClientID contains the client ID.
	ClientID string
	// ClientHost contains the client host.
	ClientHost string
	// SubscribedTopicNames contains the subscribed topic names.
	SubscribedTopicNames []string
	// SubscribedTopicRegex contains a the subscribed topic regex otherwise or null of not provided.
	SubscribedTopicRegex *string
	// Assignment contains the current assignment.
	Assignment Assignment_ConsumerGroupDescribeResponse
	// TargetAssignment contains the target assignment.
	TargetAssignment Assignment_ConsumerGroupDescribeResponse
	// MemberType contains a -1 for unknown. 0 for classic member. +1 for consumer member.
	MemberType int8
}

func (m *Member_ConsumerGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	if err := pe.putString(m.MemberID); err != nil {
		return err
	}

	if err := pe.putNullableString(m.InstanceID); err != nil {
		return err
	}

	if err := pe.putNullableString(m.RackID); err != nil {
		return err
	}

	pe.putInt32(m.MemberEpoch)

	if err := pe.putString(m.ClientID); err != nil {
		return err
	}

	if err := pe.putString(m.ClientHost); err != nil {
		return err
	}

	if err := pe.putStringArray(m.SubscribedTopicNames); err != nil {
		return err
	}

	if err := pe.putNullableString(m.SubscribedTopicRegex); err != nil {
		return err
	}

	if err := m.Assignment.encode(pe, m.Version); err != nil {
		return err
	}

	if err := m.TargetAssignment.encode(pe, m.Version); err != nil {
		return err
	}

	if m.Version >= 1 {
		pe.putInt8(m.MemberType)
	}

	pe.putUVarint(0)
	return nil
}

func (m *Member_ConsumerGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if m.InstanceID, err = pd.getNullableString(); err != nil {
		return err
	}

	if m.RackID, err = pd.getNullableString(); err != nil {
		return err
	}

	if m.MemberEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if m.ClientID, err = pd.getString(); err != nil {
		return err
	}

	if m.ClientHost, err = pd.getString(); err != nil {
		return err
	}

	if m.SubscribedTopicNames, err = pd.getStringArray(); err != nil {
		return err
	}

	if m.SubscribedTopicRegex, err = pd.getNullableString(); err != nil {
		return err
	}

	tmpAssignment := Assignment_ConsumerGroupDescribeResponse{}
	if err := tmpAssignment.decode(pd, m.Version); err != nil {
		return err
	}
	m.Assignment = tmpAssignment

	tmpTargetAssignment := Assignment_ConsumerGroupDescribeResponse{}
	if err := tmpTargetAssignment.decode(pd, m.Version); err != nil {
		return err
	}
	m.TargetAssignment = tmpTargetAssignment

	if m.Version >= 1 {
		if m.MemberType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// DescribedGroup_ConsumerGroupDescribeResponse contains each described group.
type DescribedGroup_ConsumerGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the describe error, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// GroupID contains the group ID string.
	GroupID string
	// GroupState contains the group state string, or the empty string.
	GroupState string
	// GroupEpoch contains the group epoch.
	GroupEpoch int32
	// AssignmentEpoch contains the assignment epoch.
	AssignmentEpoch int32
	// AssignorName contains the selected assignor.
	AssignorName string
	// Members contains the members.
	Members []Member_ConsumerGroupDescribeResponse
	// AuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this group.
	AuthorizedOperations int32
}

func (g *DescribedGroup_ConsumerGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	g.Version = version
	pe.putInt16(g.ErrorCode)

	if err := pe.putNullableString(g.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putString(g.GroupID); err != nil {
		return err
	}

	if err := pe.putString(g.GroupState); err != nil {
		return err
	}

	pe.putInt32(g.GroupEpoch)

	pe.putInt32(g.AssignmentEpoch)

	if err := pe.putString(g.AssignorName); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(g.Members)); err != nil {
		return err
	}
	for _, block := range g.Members {
		if err := block.encode(pe, g.Version); err != nil {
			return err
		}
	}

	pe.putInt32(g.AuthorizedOperations)

	pe.putUVarint(0)
	return nil
}

func (g *DescribedGroup_ConsumerGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	g.Version = version
	if g.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if g.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if g.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if g.GroupState, err = pd.getString(); err != nil {
		return err
	}

	if g.GroupEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if g.AssignmentEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if g.AssignorName, err = pd.getString(); err != nil {
		return err
	}

	var numMembers int
	if numMembers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numMembers > 0 {
		g.Members = make([]Member_ConsumerGroupDescribeResponse, numMembers)
		for i := 0; i < numMembers; i++ {
			var block Member_ConsumerGroupDescribeResponse
			if err := block.decode(pd, g.Version); err != nil {
				return err
			}
			g.Members[i] = block
		}
	}

	if g.AuthorizedOperations, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ConsumerGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Groups contains each described group.
	Groups []DescribedGroup_ConsumerGroupDescribeResponse
}

func (r *ConsumerGroupDescribeResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Groups)); err != nil {
		return err
	}
	for _, block := range r.Groups {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ConsumerGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numGroups int
	if numGroups, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numGroups > 0 {
		r.Groups = make([]DescribedGroup_ConsumerGroupDescribeResponse, numGroups)
		for i := 0; i < numGroups; i++ {
			var block DescribedGroup_ConsumerGroupDescribeResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Groups[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ConsumerGroupDescribeResponse) GetKey() int16 {
	return 69
}

func (r *ConsumerGroupDescribeResponse) GetVersion() int16 {
	return r.Version
}

func (r *ConsumerGroupDescribeResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ConsumerGroupDescribeResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *ConsumerGroupDescribeResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ConsumerGroupDescribeResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
