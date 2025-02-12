// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// Endpoint_StreamsGroupDescribeResponse contains a
type Endpoint_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Host contains a host of the endpoint
	Host string
	// Port contains a port of the endpoint
	Port uint16
}

func (e *Endpoint_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if err := pe.putString(e.Host); err != nil {
		return err
	}

	pe.putUint16(e.Port)

	pe.putUVarint(0)
	return nil
}

func (e *Endpoint_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	e.Version = version
	if e.Host, err = pd.getString(); err != nil {
		return err
	}

	if e.Port, err = pd.getUint16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TaskOffset_StreamsGroupDescribeResponse contains a
type TaskOffset_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains the subtopology identifier.
	SubtopologyID string
	// Partition contains the partition.
	Partition int32
	// Offset contains the offset.
	Offset int64
}

func (t *TaskOffset_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.SubtopologyID); err != nil {
		return err
	}

	pe.putInt32(t.Partition)

	pe.putInt64(t.Offset)

	pe.putUVarint(0)
	return nil
}

func (t *TaskOffset_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.SubtopologyID, err = pd.getString(); err != nil {
		return err
	}

	if t.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if t.Offset, err = pd.getInt64(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicPartitions_StreamsGroupDescribeResponse contains a
type TopicPartitions_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic ID.
	TopicID uuid.UUID
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partitions.
	Partitions []int32
}

func (t *TopicPartitions_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicPartitions_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
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

// Assignment_StreamsGroupDescribeResponse contains a
type Assignment_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ActiveTasks contains a Active tasks for this client.
	ActiveTasks []TaskIds_StreamsGroupDescribeResponse
	// StandbyTasks contains a Standby tasks for this client.
	StandbyTasks []TaskIds_StreamsGroupDescribeResponse
	// WarmupTasks contains a Warm-up tasks for this client.
	WarmupTasks []TaskIds_StreamsGroupDescribeResponse
}

func (a *Assignment_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putArrayLength(len(a.ActiveTasks)); err != nil {
		return err
	}
	for _, block := range a.ActiveTasks {
		if err := block.encode(pe, a.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(a.StandbyTasks)); err != nil {
		return err
	}
	for _, block := range a.StandbyTasks {
		if err := block.encode(pe, a.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(a.WarmupTasks)); err != nil {
		return err
	}
	for _, block := range a.WarmupTasks {
		if err := block.encode(pe, a.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (a *Assignment_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	var numActiveTasks int
	if numActiveTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numActiveTasks > 0 {
		a.ActiveTasks = make([]TaskIds_StreamsGroupDescribeResponse, numActiveTasks)
		for i := 0; i < numActiveTasks; i++ {
			var block TaskIds_StreamsGroupDescribeResponse
			if err := block.decode(pd, a.Version); err != nil {
				return err
			}
			a.ActiveTasks[i] = block
		}
	}

	var numStandbyTasks int
	if numStandbyTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numStandbyTasks > 0 {
		a.StandbyTasks = make([]TaskIds_StreamsGroupDescribeResponse, numStandbyTasks)
		for i := 0; i < numStandbyTasks; i++ {
			var block TaskIds_StreamsGroupDescribeResponse
			if err := block.decode(pd, a.Version); err != nil {
				return err
			}
			a.StandbyTasks[i] = block
		}
	}

	var numWarmupTasks int
	if numWarmupTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numWarmupTasks > 0 {
		a.WarmupTasks = make([]TaskIds_StreamsGroupDescribeResponse, numWarmupTasks)
		for i := 0; i < numWarmupTasks; i++ {
			var block TaskIds_StreamsGroupDescribeResponse
			if err := block.decode(pd, a.Version); err != nil {
				return err
			}
			a.WarmupTasks[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TaskIds_StreamsGroupDescribeResponse contains a
type TaskIds_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains the subtopology identifier.
	SubtopologyID string
	// Partitions contains the partitions of the input topics processed by this member.
	Partitions []int32
}

func (t *TaskIds_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.SubtopologyID); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TaskIds_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.SubtopologyID, err = pd.getString(); err != nil {
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

// KeyValue_StreamsGroupDescribeResponse contains a
type KeyValue_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Key contains a key of the config
	Key string
	// Value contains a value of the config
	Value string
}

func (k *KeyValue_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	k.Version = version
	if err := pe.putString(k.Key); err != nil {
		return err
	}

	if err := pe.putString(k.Value); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (k *KeyValue_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	k.Version = version
	if k.Key, err = pd.getString(); err != nil {
		return err
	}

	if k.Value, err = pd.getString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicInfo_StreamsGroupDescribeResponse contains a
type TopicInfo_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the topic.
	Name string
	// Partitions contains the number of partitions in the topic. Can be 0 if no specific number of partitions is enforced. Always 0 for changelog topics.
	Partitions int32
	// ReplicationFactor contains the replication factor of the topic. Can be 0 if the default replication factor should be used.
	ReplicationFactor int16
	// TopicConfigs contains a Topic-level configurations as key-value pairs.
	TopicConfigs []KeyValue_StreamsGroupDescribeResponse
}

func (t *TopicInfo_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	pe.putInt32(t.Partitions)

	pe.putInt16(t.ReplicationFactor)

	if err := pe.putArrayLength(len(t.TopicConfigs)); err != nil {
		return err
	}
	for _, block := range t.TopicConfigs {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicInfo_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.Partitions, err = pd.getInt32(); err != nil {
		return err
	}

	if t.ReplicationFactor, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopicConfigs int
	if numTopicConfigs, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopicConfigs > 0 {
		t.TopicConfigs = make([]KeyValue_StreamsGroupDescribeResponse, numTopicConfigs)
		for i := 0; i < numTopicConfigs; i++ {
			var block KeyValue_StreamsGroupDescribeResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.TopicConfigs[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Subtopology_StreamsGroupDescribeResponse contains the subtopologies of the streams application. This contains the configured subtopologies, where the number of partitions are set and any regular expressions are resolved to actual topics. Null if the group is uninitialized, source topics are missing or incorrectly partitioned.
type Subtopology_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains a String to uniquely identify the subtopology.
	SubtopologyID string
	// SourceTopics contains the topics the subtopology reads from.
	SourceTopics []string
	// RepartitionSinkTopics contains the repartition topics the subtopology writes to.
	RepartitionSinkTopics []string
	// StateChangelogTopics contains the set of state changelog topics associated with this subtopology. Created automatically.
	StateChangelogTopics []TopicInfo_StreamsGroupDescribeResponse
	// RepartitionSourceTopics contains the set of source topics that are internally created repartition topics. Created automatically.
	RepartitionSourceTopics []TopicInfo_StreamsGroupDescribeResponse
}

func (s *Subtopology_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	if err := pe.putString(s.SubtopologyID); err != nil {
		return err
	}

	if err := pe.putStringArray(s.SourceTopics); err != nil {
		return err
	}

	if err := pe.putStringArray(s.RepartitionSinkTopics); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(s.StateChangelogTopics)); err != nil {
		return err
	}
	for _, block := range s.StateChangelogTopics {
		if err := block.encode(pe, s.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(s.RepartitionSourceTopics)); err != nil {
		return err
	}
	for _, block := range s.RepartitionSourceTopics {
		if err := block.encode(pe, s.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (s *Subtopology_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.SubtopologyID, err = pd.getString(); err != nil {
		return err
	}

	if s.SourceTopics, err = pd.getStringArray(); err != nil {
		return err
	}

	if s.RepartitionSinkTopics, err = pd.getStringArray(); err != nil {
		return err
	}

	var numStateChangelogTopics int
	if numStateChangelogTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numStateChangelogTopics > 0 {
		s.StateChangelogTopics = make([]TopicInfo_StreamsGroupDescribeResponse, numStateChangelogTopics)
		for i := 0; i < numStateChangelogTopics; i++ {
			var block TopicInfo_StreamsGroupDescribeResponse
			if err := block.decode(pd, s.Version); err != nil {
				return err
			}
			s.StateChangelogTopics[i] = block
		}
	}

	var numRepartitionSourceTopics int
	if numRepartitionSourceTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numRepartitionSourceTopics > 0 {
		s.RepartitionSourceTopics = make([]TopicInfo_StreamsGroupDescribeResponse, numRepartitionSourceTopics)
		for i := 0; i < numRepartitionSourceTopics; i++ {
			var block TopicInfo_StreamsGroupDescribeResponse
			if err := block.decode(pd, s.Version); err != nil {
				return err
			}
			s.RepartitionSourceTopics[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Topology_StreamsGroupDescribeResponse contains the topology metadata currently initialized for the streams application. Can be null in case of a describe error.
type Topology_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Epoch contains the epoch of the currently initialized topology for this group.
	Epoch int32
	// Subtopologies contains the subtopologies of the streams application. This contains the configured subtopologies, where the number of partitions are set and any regular expressions are resolved to actual topics. Null if the group is uninitialized, source topics are missing or incorrectly partitioned.
	Subtopologies []Subtopology_StreamsGroupDescribeResponse
}

func (t *Topology_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	pe.putInt32(t.Epoch)

	if err := pe.putArrayLength(len(t.Subtopologies)); err != nil {
		return err
	}
	for _, block := range t.Subtopologies {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (t *Topology_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Epoch, err = pd.getInt32(); err != nil {
		return err
	}

	var numSubtopologies int
	if numSubtopologies, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numSubtopologies > 0 {
		t.Subtopologies = make([]Subtopology_StreamsGroupDescribeResponse, numSubtopologies)
		for i := 0; i < numSubtopologies; i++ {
			var block Subtopology_StreamsGroupDescribeResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Subtopologies[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Member_StreamsGroupDescribeResponse contains the members.
type Member_StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the member ID.
	MemberID string
	// MemberEpoch contains the member epoch.
	MemberEpoch int32
	// InstanceID contains the member instance ID for static membership.
	InstanceID *string
	// RackID contains the rack ID.
	RackID *string
	// ClientID contains the client ID.
	ClientID string
	// ClientHost contains the client host.
	ClientHost string
	// TopologyEpoch contains the epoch of the topology on the client.
	TopologyEpoch int32
	// ProcessID contains a Identity of the streams instance that may have multiple clients.
	ProcessID string
	// UserEndpoint contains a User-defined endpoint for Interactive Queries. Null if not defined for this client.
	UserEndpoint Endpoint_StreamsGroupDescribeResponse
	// ClientTags contains a Used for rack-aware assignment algorithm.
	ClientTags []KeyValue_StreamsGroupDescribeResponse
	// TaskOffsets contains a Cumulative changelog offsets for tasks.
	TaskOffsets []TaskOffset_StreamsGroupDescribeResponse
	// TaskEndOffsets contains a Cumulative changelog end offsets for tasks.
	TaskEndOffsets []TaskOffset_StreamsGroupDescribeResponse
	// Assignment contains the current assignment.
	Assignment Assignment_StreamsGroupDescribeResponse
	// TargetAssignment contains the target assignment.
	TargetAssignment Assignment_StreamsGroupDescribeResponse
	// IsClassic contains a True for classic members that have not been upgraded yet.
	IsClassic bool
}

func (m *Member_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	if err := pe.putString(m.MemberID); err != nil {
		return err
	}

	pe.putInt32(m.MemberEpoch)

	if err := pe.putNullableString(m.InstanceID); err != nil {
		return err
	}

	if err := pe.putNullableString(m.RackID); err != nil {
		return err
	}

	if err := pe.putString(m.ClientID); err != nil {
		return err
	}

	if err := pe.putString(m.ClientHost); err != nil {
		return err
	}

	pe.putInt32(m.TopologyEpoch)

	if err := pe.putString(m.ProcessID); err != nil {
		return err
	}

	if err := m.UserEndpoint.encode(pe, m.Version); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(m.ClientTags)); err != nil {
		return err
	}
	for _, block := range m.ClientTags {
		if err := block.encode(pe, m.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(m.TaskOffsets)); err != nil {
		return err
	}
	for _, block := range m.TaskOffsets {
		if err := block.encode(pe, m.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(m.TaskEndOffsets)); err != nil {
		return err
	}
	for _, block := range m.TaskEndOffsets {
		if err := block.encode(pe, m.Version); err != nil {
			return err
		}
	}

	if err := m.Assignment.encode(pe, m.Version); err != nil {
		return err
	}

	if err := m.TargetAssignment.encode(pe, m.Version); err != nil {
		return err
	}

	pe.putBool(m.IsClassic)

	pe.putUVarint(0)
	return nil
}

func (m *Member_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if m.MemberEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if m.InstanceID, err = pd.getNullableString(); err != nil {
		return err
	}

	if m.RackID, err = pd.getNullableString(); err != nil {
		return err
	}

	if m.ClientID, err = pd.getString(); err != nil {
		return err
	}

	if m.ClientHost, err = pd.getString(); err != nil {
		return err
	}

	if m.TopologyEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if m.ProcessID, err = pd.getString(); err != nil {
		return err
	}

	tmpUserEndpoint := Endpoint_StreamsGroupDescribeResponse{}
	if err := tmpUserEndpoint.decode(pd, m.Version); err != nil {
		return err
	}
	m.UserEndpoint = tmpUserEndpoint

	var numClientTags int
	if numClientTags, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numClientTags > 0 {
		m.ClientTags = make([]KeyValue_StreamsGroupDescribeResponse, numClientTags)
		for i := 0; i < numClientTags; i++ {
			var block KeyValue_StreamsGroupDescribeResponse
			if err := block.decode(pd, m.Version); err != nil {
				return err
			}
			m.ClientTags[i] = block
		}
	}

	var numTaskOffsets int
	if numTaskOffsets, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTaskOffsets > 0 {
		m.TaskOffsets = make([]TaskOffset_StreamsGroupDescribeResponse, numTaskOffsets)
		for i := 0; i < numTaskOffsets; i++ {
			var block TaskOffset_StreamsGroupDescribeResponse
			if err := block.decode(pd, m.Version); err != nil {
				return err
			}
			m.TaskOffsets[i] = block
		}
	}

	var numTaskEndOffsets int
	if numTaskEndOffsets, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTaskEndOffsets > 0 {
		m.TaskEndOffsets = make([]TaskOffset_StreamsGroupDescribeResponse, numTaskEndOffsets)
		for i := 0; i < numTaskEndOffsets; i++ {
			var block TaskOffset_StreamsGroupDescribeResponse
			if err := block.decode(pd, m.Version); err != nil {
				return err
			}
			m.TaskEndOffsets[i] = block
		}
	}

	tmpAssignment := Assignment_StreamsGroupDescribeResponse{}
	if err := tmpAssignment.decode(pd, m.Version); err != nil {
		return err
	}
	m.Assignment = tmpAssignment

	tmpTargetAssignment := Assignment_StreamsGroupDescribeResponse{}
	if err := tmpTargetAssignment.decode(pd, m.Version); err != nil {
		return err
	}
	m.TargetAssignment = tmpTargetAssignment

	if m.IsClassic, err = pd.getBool(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// DescribedGroup_StreamsGroupDescribeResponse contains each described group.
type DescribedGroup_StreamsGroupDescribeResponse struct {
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
	// Topology contains the topology metadata currently initialized for the streams application. Can be null in case of a describe error.
	Topology Topology_StreamsGroupDescribeResponse
	// Members contains the members.
	Members []Member_StreamsGroupDescribeResponse
	// AuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this group.
	AuthorizedOperations int32
}

func (g *DescribedGroup_StreamsGroupDescribeResponse) encode(pe packetEncoder, version int16) (err error) {
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

	if err := g.Topology.encode(pe, g.Version); err != nil {
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

func (g *DescribedGroup_StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
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

	tmpTopology := Topology_StreamsGroupDescribeResponse{}
	if err := tmpTopology.decode(pd, g.Version); err != nil {
		return err
	}
	g.Topology = tmpTopology

	var numMembers int
	if numMembers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numMembers > 0 {
		g.Members = make([]Member_StreamsGroupDescribeResponse, numMembers)
		for i := 0; i < numMembers; i++ {
			var block Member_StreamsGroupDescribeResponse
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

type StreamsGroupDescribeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Groups contains each described group.
	Groups []DescribedGroup_StreamsGroupDescribeResponse
}

func (r *StreamsGroupDescribeResponse) encode(pe packetEncoder) (err error) {
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

func (r *StreamsGroupDescribeResponse) decode(pd packetDecoder, version int16) (err error) {
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
		r.Groups = make([]DescribedGroup_StreamsGroupDescribeResponse, numGroups)
		for i := 0; i < numGroups; i++ {
			var block DescribedGroup_StreamsGroupDescribeResponse
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

func (r *StreamsGroupDescribeResponse) GetKey() int16 {
	return 89
}

func (r *StreamsGroupDescribeResponse) GetVersion() int16 {
	return r.Version
}

func (r *StreamsGroupDescribeResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *StreamsGroupDescribeResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *StreamsGroupDescribeResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *StreamsGroupDescribeResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
