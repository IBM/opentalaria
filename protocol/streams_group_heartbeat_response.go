// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// Status_StreamsGroupHeartbeatResponse contains a
type Status_StreamsGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// StatusCode contains a A code to indicate that a particular status is active for the group membership
	StatusCode int8
	// StatusDetail contains a A string representation of the status.
	StatusDetail string
}

func (s *Status_StreamsGroupHeartbeatResponse) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	pe.putInt8(s.StatusCode)

	if err := pe.putString(s.StatusDetail); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (s *Status_StreamsGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.StatusCode, err = pd.getInt8(); err != nil {
		return err
	}

	if s.StatusDetail, err = pd.getString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicPartition_StreamsGroupHeartbeatResponse contains a
type TopicPartition_StreamsGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains a topic name
	Topic string
	// Partitions contains a partitions
	Partitions []int32
}

func (t *TopicPartition_StreamsGroupHeartbeatResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Topic); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicPartition_StreamsGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Topic, err = pd.getString(); err != nil {
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

// TaskIds_StreamsGroupHeartbeatResponse contains a
type TaskIds_StreamsGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// SubtopologyID contains the subtopology identifier.
	SubtopologyID string
	// Partitions contains the partitions of the input topics processed by this member.
	Partitions []int32
}

func (t *TaskIds_StreamsGroupHeartbeatResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TaskIds_StreamsGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
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

// Endpoint_StreamsGroupHeartbeatResponse contains a
type Endpoint_StreamsGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Host contains a host of the endpoint
	Host string
	// Port contains a port of the endpoint
	Port uint16
}

func (e *Endpoint_StreamsGroupHeartbeatResponse) encode(pe packetEncoder, version int16) (err error) {
	e.Version = version
	if err := pe.putString(e.Host); err != nil {
		return err
	}

	pe.putUint16(e.Port)

	pe.putUVarint(0)
	return nil
}

func (e *Endpoint_StreamsGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
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

// EndpointToPartitions contains a Global assignment information used for IQ. Null if unchanged since last heartbeat.
type EndpointToPartitions struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// UserEndpoint contains a User-defined endpoint to connect to the node
	UserEndpoint Endpoint_StreamsGroupHeartbeatResponse
	// Partitions contains a All partitions available on the node
	Partitions []TopicPartition_StreamsGroupHeartbeatResponse
}

func (p *EndpointToPartitions) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if err := p.UserEndpoint.encode(pe, p.Version); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(p.Partitions)); err != nil {
		return err
	}
	for _, block := range p.Partitions {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (p *EndpointToPartitions) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	tmpUserEndpoint := Endpoint_StreamsGroupHeartbeatResponse{}
	if err := tmpUserEndpoint.decode(pd, p.Version); err != nil {
		return err
	}
	p.UserEndpoint = tmpUserEndpoint

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		p.Partitions = make([]TopicPartition_StreamsGroupHeartbeatResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block TopicPartition_StreamsGroupHeartbeatResponse
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.Partitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type StreamsGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level error code, or 0 if there was no error
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// MemberID contains the member id is always generated by the streams consumer.
	MemberID string
	// MemberEpoch contains the member epoch.
	MemberEpoch int32
	// HeartbeatIntervalMs contains the heartbeat interval in milliseconds.
	HeartbeatIntervalMs int32
	// AcceptableRecoveryLag contains the maximal lag a warm-up task can have to be considered caught-up.
	AcceptableRecoveryLag int32
	// TaskOffsetIntervalMs contains the interval in which the task changelog offsets on a client are updated on the broker. The offsets are sent with the next heartbeat after this time has passed.
	TaskOffsetIntervalMs int32
	// Status contains a Indicate zero or more status for the group.  Null if unchanged since last heartbeat.
	Status []Status_StreamsGroupHeartbeatResponse
	// ActiveTasks contains a Assigned active tasks for this client. Null if unchanged since last heartbeat.
	ActiveTasks []TaskIds_StreamsGroupHeartbeatResponse
	// StandbyTasks contains a Assigned standby tasks for this client. Null if unchanged since last heartbeat.
	StandbyTasks []TaskIds_StreamsGroupHeartbeatResponse
	// WarmupTasks contains a Assigned warm-up tasks for this client. Null if unchanged since last heartbeat.
	WarmupTasks []TaskIds_StreamsGroupHeartbeatResponse
	// PartitionsByUserEndpoint contains a Global assignment information used for IQ. Null if unchanged since last heartbeat.
	PartitionsByUserEndpoint []EndpointToPartitions
}

func (r *StreamsGroupHeartbeatResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putString(r.MemberID); err != nil {
		return err
	}

	pe.putInt32(r.MemberEpoch)

	pe.putInt32(r.HeartbeatIntervalMs)

	pe.putInt32(r.AcceptableRecoveryLag)

	pe.putInt32(r.TaskOffsetIntervalMs)

	if err := pe.putArrayLength(len(r.Status)); err != nil {
		return err
	}
	for _, block := range r.Status {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
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

	if err := pe.putArrayLength(len(r.PartitionsByUserEndpoint)); err != nil {
		return err
	}
	for _, block := range r.PartitionsByUserEndpoint {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *StreamsGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
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

	if r.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if r.MemberEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.HeartbeatIntervalMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.AcceptableRecoveryLag, err = pd.getInt32(); err != nil {
		return err
	}

	if r.TaskOffsetIntervalMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numStatus int
	if numStatus, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numStatus > 0 {
		r.Status = make([]Status_StreamsGroupHeartbeatResponse, numStatus)
		for i := 0; i < numStatus; i++ {
			var block Status_StreamsGroupHeartbeatResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Status[i] = block
		}
	}

	var numActiveTasks int
	if numActiveTasks, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numActiveTasks > 0 {
		r.ActiveTasks = make([]TaskIds_StreamsGroupHeartbeatResponse, numActiveTasks)
		for i := 0; i < numActiveTasks; i++ {
			var block TaskIds_StreamsGroupHeartbeatResponse
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
		r.StandbyTasks = make([]TaskIds_StreamsGroupHeartbeatResponse, numStandbyTasks)
		for i := 0; i < numStandbyTasks; i++ {
			var block TaskIds_StreamsGroupHeartbeatResponse
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
		r.WarmupTasks = make([]TaskIds_StreamsGroupHeartbeatResponse, numWarmupTasks)
		for i := 0; i < numWarmupTasks; i++ {
			var block TaskIds_StreamsGroupHeartbeatResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.WarmupTasks[i] = block
		}
	}

	var numPartitionsByUserEndpoint int
	if numPartitionsByUserEndpoint, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitionsByUserEndpoint > 0 {
		r.PartitionsByUserEndpoint = make([]EndpointToPartitions, numPartitionsByUserEndpoint)
		for i := 0; i < numPartitionsByUserEndpoint; i++ {
			var block EndpointToPartitions
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.PartitionsByUserEndpoint[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *StreamsGroupHeartbeatResponse) GetKey() int16 {
	return 88
}

func (r *StreamsGroupHeartbeatResponse) GetVersion() int16 {
	return r.Version
}

func (r *StreamsGroupHeartbeatResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *StreamsGroupHeartbeatResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *StreamsGroupHeartbeatResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *StreamsGroupHeartbeatResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
