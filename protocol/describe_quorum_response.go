// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// ReplicaState_DescribeQuorumResponse contains a
type ReplicaState_DescribeQuorumResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ReplicaID contains the ID of the replica.
	ReplicaID int32
	// ReplicaDirectoryID contains the replica directory ID of the replica.
	ReplicaDirectoryID uuid.UUID
	// LogEndOffset contains the last known log end offset of the follower or -1 if it is unknown.
	LogEndOffset int64
	// LastFetchTimestamp contains the last known leader wall clock time time when a follower fetched from the leader. This is reported as -1 both for the current leader or if it is unknown for a voter.
	LastFetchTimestamp int64
	// LastCaughtUpTimestamp contains the leader wall clock append time of the offset for which the follower made the most recent fetch request. This is reported as the current time for the leader and -1 if unknown for a voter.
	LastCaughtUpTimestamp int64
}

func (r *ReplicaState_DescribeQuorumResponse) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	pe.putInt32(r.ReplicaID)

	if r.Version >= 2 {
		if err := pe.putUUID(r.ReplicaDirectoryID); err != nil {
			return err
		}
	}

	pe.putInt64(r.LogEndOffset)

	if r.Version >= 1 {
		pe.putInt64(r.LastFetchTimestamp)
	}

	if r.Version >= 1 {
		pe.putInt64(r.LastCaughtUpTimestamp)
	}

	pe.putUVarint(0)
	return nil
}

func (r *ReplicaState_DescribeQuorumResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ReplicaID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if r.ReplicaDirectoryID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if r.LogEndOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.LastFetchTimestamp, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		if r.LastCaughtUpTimestamp, err = pd.getInt64(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// PartitionData_DescribeQuorumResponse contains the partition data.
type PartitionData_DescribeQuorumResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition error code.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch.
	LeaderEpoch int32
	// HighWatermark contains the high water mark.
	HighWatermark int64
	// CurrentVoters contains the current voters of the partition.
	CurrentVoters []ReplicaState_DescribeQuorumResponse
	// Observers contains the observers of the partition.
	Observers []ReplicaState_DescribeQuorumResponse
}

func (p *PartitionData_DescribeQuorumResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if p.Version >= 2 {
		if err := pe.putNullableString(p.ErrorMessage); err != nil {
			return err
		}
	}

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	pe.putInt64(p.HighWatermark)

	if err := pe.putArrayLength(len(p.CurrentVoters)); err != nil {
		return err
	}
	for _, block := range p.CurrentVoters {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(p.Observers)); err != nil {
		return err
	}
	for _, block := range p.Observers {
		if err := block.encode(pe, p.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_DescribeQuorumResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if p.ErrorMessage, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.HighWatermark, err = pd.getInt64(); err != nil {
		return err
	}

	var numCurrentVoters int
	if numCurrentVoters, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numCurrentVoters > 0 {
		p.CurrentVoters = make([]ReplicaState_DescribeQuorumResponse, numCurrentVoters)
		for i := 0; i < numCurrentVoters; i++ {
			var block ReplicaState_DescribeQuorumResponse
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.CurrentVoters[i] = block
		}
	}

	var numObservers int
	if numObservers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numObservers > 0 {
		p.Observers = make([]ReplicaState_DescribeQuorumResponse, numObservers)
		for i := 0; i < numObservers; i++ {
			var block ReplicaState_DescribeQuorumResponse
			if err := block.decode(pd, p.Version); err != nil {
				return err
			}
			p.Observers[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_DescribeQuorumResponse contains the response from the describe quorum API.
type TopicData_DescribeQuorumResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partition data.
	Partitions []PartitionData_DescribeQuorumResponse
}

func (t *TopicData_DescribeQuorumResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.TopicName); err != nil {
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

func (t *TopicData_DescribeQuorumResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_DescribeQuorumResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_DescribeQuorumResponse
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

// Listener_DescribeQuorumResponse contains the listeners of this controller.
type Listener_DescribeQuorumResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the hostname.
	Host string
	// Port contains the port.
	Port uint16
}

func (l *Listener_DescribeQuorumResponse) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 2 {
		if err := pe.putString(l.Name); err != nil {
			return err
		}
	}

	if l.Version >= 2 {
		if err := pe.putString(l.Host); err != nil {
			return err
		}
	}

	if l.Version >= 2 {
		pe.putUint16(l.Port)
	}

	pe.putUVarint(0)
	return nil
}

func (l *Listener_DescribeQuorumResponse) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 2 {
		if l.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.Version >= 2 {
		if l.Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.Version >= 2 {
		if l.Port, err = pd.getUint16(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Node contains the nodes in the quorum.
type Node struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// NodeID contains the ID of the associated node.
	NodeID int32
	// Listeners contains the listeners of this controller.
	Listeners []Listener_DescribeQuorumResponse
}

func (n *Node) encode(pe packetEncoder, version int16) (err error) {
	n.Version = version
	if n.Version >= 2 {
		pe.putInt32(n.NodeID)
	}

	if n.Version >= 2 {
		if err := pe.putArrayLength(len(n.Listeners)); err != nil {
			return err
		}
		for _, block := range n.Listeners {
			if err := block.encode(pe, n.Version); err != nil {
				return err
			}
		}
	}

	pe.putUVarint(0)
	return nil
}

func (n *Node) decode(pd packetDecoder, version int16) (err error) {
	n.Version = version
	if n.Version >= 2 {
		if n.NodeID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if n.Version >= 2 {
		var numListeners int
		if numListeners, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numListeners > 0 {
			n.Listeners = make([]Listener_DescribeQuorumResponse, numListeners)
			for i := 0; i < numListeners; i++ {
				var block Listener_DescribeQuorumResponse
				if err := block.decode(pd, n.Version); err != nil {
					return err
				}
				n.Listeners[i] = block
			}
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeQuorumResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top level error code.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// Topics contains the response from the describe quorum API.
	Topics []TopicData_DescribeQuorumResponse
	// Nodes contains the nodes in the quorum.
	Nodes []Node
}

func (r *DescribeQuorumResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt16(r.ErrorCode)

	if r.Version >= 2 {
		if err := pe.putNullableString(r.ErrorMessage); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		if err := pe.putArrayLength(len(r.Nodes)); err != nil {
			return err
		}
		for _, block := range r.Nodes {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeQuorumResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if r.ErrorMessage, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_DescribeQuorumResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_DescribeQuorumResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 2 {
		var numNodes int
		if numNodes, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numNodes > 0 {
			r.Nodes = make([]Node, numNodes)
			for i := 0; i < numNodes; i++ {
				var block Node
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.Nodes[i] = block
			}
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeQuorumResponse) GetKey() int16 {
	return 55
}

func (r *DescribeQuorumResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeQuorumResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *DescribeQuorumResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeQuorumResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *DescribeQuorumResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
