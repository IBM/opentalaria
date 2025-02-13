// protocol has been generated from message format json - DO NOT EDIT
package protocol

// KeyValue_StreamsGroupHeartbeatRequest contains a
type KeyValue_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Key contains a key of the config
	Key string
	// Value contains a value of the config
	Value string
}

func (k *KeyValue_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (k *KeyValue_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
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

// TopicInfo_StreamsGroupHeartbeatRequest contains a
type TopicInfo_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the topic.
	Name string
	// Partitions contains the number of partitions in the topic. Can be 0 if no specific number of partitions is enforced. Always 0 for changelog topics.
	Partitions int32
	// ReplicationFactor contains the replication factor of the topic. Can be 0 if the default replication factor should be used.
	ReplicationFactor int16
	// TopicConfigs contains a Topic-level configurations as key-value pairs.
	TopicConfigs []KeyValue_StreamsGroupHeartbeatRequest
}

func (t *TopicInfo_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicInfo_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
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
		t.TopicConfigs = make([]KeyValue_StreamsGroupHeartbeatRequest, numTopicConfigs)
		for i := 0; i < numTopicConfigs; i++ {
			var block KeyValue_StreamsGroupHeartbeatRequest
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

// Endpoint_StreamsGroupHeartbeatRequest contains a
type Endpoint_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Host contains a host of the endpoint
	Host string
	// Port contains a port of the endpoint
	Port uint16
}

func (e *Endpoint_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if err := pe.putString(e.Host); err != nil {
		return err
	}

	pe.putUint16(e.Port)

	pe.putUVarint(0)
	return nil
}

func (e *Endpoint_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
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

// TaskOffset_StreamsGroupHeartbeatRequest contains a
type TaskOffset_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains the subtopology identifier.
	SubtopologyID string
	// Partition contains the partition.
	Partition int32
	// Offset contains the offset.
	Offset int64
}

func (t *TaskOffset_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.SubtopologyID); err != nil {
		return err
	}

	pe.putInt32(t.Partition)

	pe.putInt64(t.Offset)

	pe.putUVarint(0)
	return nil
}

func (t *TaskOffset_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
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

// TaskIds_StreamsGroupHeartbeatRequest contains a
type TaskIds_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains the subtopology identifier.
	SubtopologyID string
	// Partitions contains the partitions of the input topics processed by this member.
	Partitions []int32
}

func (t *TaskIds_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TaskIds_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
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

// CopartitionGroup contains a A subset of source topics that must be copartitioned.
type CopartitionGroup struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SourceTopics contains the topics the topology reads from. Index into the array on the subtopology level.
	SourceTopics []int16
	// SourceTopicRegex contains a Regular expressions identifying topics the subtopology reads from. Index into the array on the subtopology level.
	SourceTopicRegex []int16
	// RepartitionSourceTopics contains the set of source topics that are internally created repartition topics. Index into the array on the subtopology level.
	RepartitionSourceTopics []int16
}

func (c *CopartitionGroup) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putInt16Array(c.SourceTopics); err != nil {
		return err
	}

	if err := pe.putInt16Array(c.SourceTopicRegex); err != nil {
		return err
	}

	if err := pe.putInt16Array(c.RepartitionSourceTopics); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (c *CopartitionGroup) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.SourceTopics, err = pd.getInt16Array(); err != nil {
		return err
	}

	if c.SourceTopicRegex, err = pd.getInt16Array(); err != nil {
		return err
	}

	if c.RepartitionSourceTopics, err = pd.getInt16Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Subtopology_StreamsGroupHeartbeatRequest contains the sub-topologies of the streams application.
type Subtopology_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains a String to uniquely identify the subtopology. Deterministically generated from the topology
	SubtopologyID string
	// SourceTopics contains the topics the topology reads from.
	SourceTopics []string
	// SourceTopicRegex contains the regular expressions identifying topics the subtopology reads from.
	SourceTopicRegex []string
	// StateChangelogTopics contains the set of state changelog topics associated with this subtopology. Created automatically.
	StateChangelogTopics []TopicInfo_StreamsGroupHeartbeatRequest
	// RepartitionSinkTopics contains the repartition topics the subtopology writes to.
	RepartitionSinkTopics []string
	// RepartitionSourceTopics contains the set of source topics that are internally created repartition topics. Created automatically.
	RepartitionSourceTopics []TopicInfo_StreamsGroupHeartbeatRequest
	// CopartitionGroups contains a A subset of source topics that must be copartitioned.
	CopartitionGroups []CopartitionGroup
}

func (s *Subtopology_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	if err := pe.putString(s.SubtopologyID); err != nil {
		return err
	}

	if err := pe.putStringArray(s.SourceTopics); err != nil {
		return err
	}

	if err := pe.putStringArray(s.SourceTopicRegex); err != nil {
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

	if err := pe.putStringArray(s.RepartitionSinkTopics); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(s.RepartitionSourceTopics)); err != nil {
		return err
	}
	for _, block := range s.RepartitionSourceTopics {
		if err := block.encode(pe, s.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(s.CopartitionGroups)); err != nil {
		return err
	}
	for _, block := range s.CopartitionGroups {
		if err := block.encode(pe, s.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (s *Subtopology_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.SubtopologyID, err = pd.getString(); err != nil {
		return err
	}

	if s.SourceTopics, err = pd.getStringArray(); err != nil {
		return err
	}

	if s.SourceTopicRegex, err = pd.getStringArray(); err != nil {
		return err
	}

	var numStateChangelogTopics int
	if numStateChangelogTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numStateChangelogTopics > 0 {
		s.StateChangelogTopics = make([]TopicInfo_StreamsGroupHeartbeatRequest, numStateChangelogTopics)
		for i := 0; i < numStateChangelogTopics; i++ {
			var block TopicInfo_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, s.Version); err != nil {
				return err
			}
			s.StateChangelogTopics[i] = block
		}
	}

	if s.RepartitionSinkTopics, err = pd.getStringArray(); err != nil {
		return err
	}

	var numRepartitionSourceTopics int
	if numRepartitionSourceTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numRepartitionSourceTopics > 0 {
		s.RepartitionSourceTopics = make([]TopicInfo_StreamsGroupHeartbeatRequest, numRepartitionSourceTopics)
		for i := 0; i < numRepartitionSourceTopics; i++ {
			var block TopicInfo_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, s.Version); err != nil {
				return err
			}
			s.RepartitionSourceTopics[i] = block
		}
	}

	var numCopartitionGroups int
	if numCopartitionGroups, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numCopartitionGroups > 0 {
		s.CopartitionGroups = make([]CopartitionGroup, numCopartitionGroups)
		for i := 0; i < numCopartitionGroups; i++ {
			var block CopartitionGroup
			if err := block.decode(pd, s.Version); err != nil {
				return err
			}
			s.CopartitionGroups[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Topology_StreamsGroupHeartbeatRequest contains the topology metadata of the streams application. Used to initialize the topology of the group and to check if the topology corresponds to the topology initialized for the group. Only sent when memberEpoch = 0, must be non-empty. Null otherwise.
type Topology_StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Epoch contains the epoch of the topology. Used to check if the topology corresponds to the topology initialized on the brokers.
	Epoch int32
	// Subtopologies contains the sub-topologies of the streams application.
	Subtopologies []Subtopology_StreamsGroupHeartbeatRequest
}

func (t *Topology_StreamsGroupHeartbeatRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *Topology_StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Epoch, err = pd.getInt32(); err != nil {
		return err
	}

	var numSubtopologies int
	if numSubtopologies, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numSubtopologies > 0 {
		t.Subtopologies = make([]Subtopology_StreamsGroupHeartbeatRequest, numSubtopologies)
		for i := 0; i < numSubtopologies; i++ {
			var block Subtopology_StreamsGroupHeartbeatRequest
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

type StreamsGroupHeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID string
	// MemberID contains the member ID generated by the streams consumer. The member ID must be kept during the entire lifetime of the streams consumer process.
	MemberID string
	// MemberEpoch contains the current member epoch; 0 to join the group; -1 to leave the group; -2 to indicate that the static member will rejoin.
	MemberEpoch int32
	// InstanceID contains a null if not provided or if it didn't change since the last heartbeat; the instance ID for static membership otherwise.
	InstanceID *string
	// RackID contains a null if not provided or if it didn't change since the last heartbeat; the rack ID of the member otherwise.
	RackID *string
	// RebalanceTimeoutMs contains a -1 if it didn't change since the last heartbeat; the maximum time in milliseconds that the coordinator will wait on the member to revoke its tasks otherwise.
	RebalanceTimeoutMs int32
	// Topology contains the topology metadata of the streams application. Used to initialize the topology of the group and to check if the topology corresponds to the topology initialized for the group. Only sent when memberEpoch = 0, must be non-empty. Null otherwise.
	Topology Topology_StreamsGroupHeartbeatRequest
	// ActiveTasks contains a Currently owned active tasks for this client. Null if unchanged since last heartbeat.
	ActiveTasks []TaskIds_StreamsGroupHeartbeatRequest
	// StandbyTasks contains a Currently owned standby tasks for this client. Null if unchanged since last heartbeat.
	StandbyTasks []TaskIds_StreamsGroupHeartbeatRequest
	// WarmupTasks contains a Currently owned warm-up tasks for this client. Null if unchanged since last heartbeat.
	WarmupTasks []TaskIds_StreamsGroupHeartbeatRequest
	// ProcessID contains a Identity of the streams instance that may have multiple consumers. Null if unchanged since last heartbeat.
	ProcessID *string
	// UserEndpoint contains a User-defined endpoint for Interactive Queries. Null if unchanged since last heartbeat, or if not defined on the client.
	UserEndpoint Endpoint_StreamsGroupHeartbeatRequest
	// ClientTags contains a Used for rack-aware assignment algorithm. Null if unchanged since last heartbeat.
	ClientTags []KeyValue_StreamsGroupHeartbeatRequest
	// TaskOffsets contains a Cumulative changelog offsets for tasks. Only updated when a warm-up task has caught up, and according to the task offset interval. Null if unchanged since last heartbeat.
	TaskOffsets []TaskOffset_StreamsGroupHeartbeatRequest
	// TaskEndOffsets contains a Cumulative changelog end-offsets for tasks. Only updated when a warm-up task has caught up, and according to the task offset interval. Null if unchanged since last heartbeat.
	TaskEndOffsets []TaskOffset_StreamsGroupHeartbeatRequest
	// ShutdownApplication contains a Whether all Streams clients in the group should shut down.
	ShutdownApplication bool
}

func (r *StreamsGroupHeartbeatRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	if err := pe.putString(r.MemberID); err != nil {
		return err
	}

	pe.putInt32(r.MemberEpoch)

	if err := pe.putNullableString(r.InstanceID); err != nil {
		return err
	}

	if err := pe.putNullableString(r.RackID); err != nil {
		return err
	}

	pe.putInt32(r.RebalanceTimeoutMs)

	if err := r.Topology.encode(pe, r.Version); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.ActiveTasks)); err != nil {
		return err
	}
	for _, block := range r.ActiveTasks {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.StandbyTasks)); err != nil {
		return err
	}
	for _, block := range r.StandbyTasks {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.WarmupTasks)); err != nil {
		return err
	}
	for _, block := range r.WarmupTasks {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putNullableString(r.ProcessID); err != nil {
		return err
	}

	if err := r.UserEndpoint.encode(pe, r.Version); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.ClientTags)); err != nil {
		return err
	}
	for _, block := range r.ClientTags {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.TaskOffsets)); err != nil {
		return err
	}
	for _, block := range r.TaskOffsets {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.TaskEndOffsets)); err != nil {
		return err
	}
	for _, block := range r.TaskEndOffsets {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putBool(r.ShutdownApplication)

	pe.putUVarint(0)
	return nil
}

func (r *StreamsGroupHeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if r.MemberEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.InstanceID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.RackID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.RebalanceTimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	tmpTopology := Topology_StreamsGroupHeartbeatRequest{}
	if err := tmpTopology.decode(pd, r.Version); err != nil {
		return err
	}
	r.Topology = tmpTopology

	var numActiveTasks int
	if numActiveTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numActiveTasks > 0 {
		r.ActiveTasks = make([]TaskIds_StreamsGroupHeartbeatRequest, numActiveTasks)
		for i := 0; i < numActiveTasks; i++ {
			var block TaskIds_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.ActiveTasks[i] = block
		}
	}

	var numStandbyTasks int
	if numStandbyTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numStandbyTasks > 0 {
		r.StandbyTasks = make([]TaskIds_StreamsGroupHeartbeatRequest, numStandbyTasks)
		for i := 0; i < numStandbyTasks; i++ {
			var block TaskIds_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.StandbyTasks[i] = block
		}
	}

	var numWarmupTasks int
	if numWarmupTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numWarmupTasks > 0 {
		r.WarmupTasks = make([]TaskIds_StreamsGroupHeartbeatRequest, numWarmupTasks)
		for i := 0; i < numWarmupTasks; i++ {
			var block TaskIds_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.WarmupTasks[i] = block
		}
	}

	if r.ProcessID, err = pd.getNullableString(); err != nil {
		return err
	}

	tmpUserEndpoint := Endpoint_StreamsGroupHeartbeatRequest{}
	if err := tmpUserEndpoint.decode(pd, r.Version); err != nil {
		return err
	}
	r.UserEndpoint = tmpUserEndpoint

	var numClientTags int
	if numClientTags, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numClientTags > 0 {
		r.ClientTags = make([]KeyValue_StreamsGroupHeartbeatRequest, numClientTags)
		for i := 0; i < numClientTags; i++ {
			var block KeyValue_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.ClientTags[i] = block
		}
	}

	var numTaskOffsets int
	if numTaskOffsets, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTaskOffsets > 0 {
		r.TaskOffsets = make([]TaskOffset_StreamsGroupHeartbeatRequest, numTaskOffsets)
		for i := 0; i < numTaskOffsets; i++ {
			var block TaskOffset_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.TaskOffsets[i] = block
		}
	}

	var numTaskEndOffsets int
	if numTaskEndOffsets, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTaskEndOffsets > 0 {
		r.TaskEndOffsets = make([]TaskOffset_StreamsGroupHeartbeatRequest, numTaskEndOffsets)
		for i := 0; i < numTaskEndOffsets; i++ {
			var block TaskOffset_StreamsGroupHeartbeatRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.TaskEndOffsets[i] = block
		}
	}

	if r.ShutdownApplication, err = pd.getBool(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *StreamsGroupHeartbeatRequest) GetKey() int16 {
	return 88
}

func (r *StreamsGroupHeartbeatRequest) GetVersion() int16 {
	return r.Version
}

func (r *StreamsGroupHeartbeatRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *StreamsGroupHeartbeatRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *StreamsGroupHeartbeatRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
