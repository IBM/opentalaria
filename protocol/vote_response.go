// protocol has been generated from message format json - DO NOT EDIT
package protocol

// PartitionData_VoteResponse contains the results for each partition.
type PartitionData_VoteResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition level error code.
	ErrorCode int16
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch.
	LeaderEpoch int32
	// VoteGranted contains a True if the vote was granted and false otherwise.
	VoteGranted bool
}

func (p *PartitionData_VoteResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	pe.putBool(p.VoteGranted)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_VoteResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.VoteGranted, err = pd.getBool(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_VoteResponse contains the results for each topic.
type TopicData_VoteResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the results for each partition.
	Partitions []PartitionData_VoteResponse
}

func (t *TopicData_VoteResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicData_VoteResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_VoteResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_VoteResponse
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

// NodeEndpoint_VoteResponse contains a Endpoints for all current-leaders enumerated in PartitionData.
type NodeEndpoint_VoteResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// NodeID contains the ID of the associated node.
	NodeID int32
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port uint16
}

func (n *NodeEndpoint_VoteResponse) encode(pe packetEncoder, version int16) (err error) {
	n.Version = version
	if n.Version >= 1 {
		pe.putInt32(n.NodeID)
	}

	if n.Version >= 1 {
		if err := pe.putString(n.Host); err != nil {
			return err
		}
	}

	if n.Version >= 1 {
		pe.putUint16(n.Port)
	}

	return nil
}

func (n *NodeEndpoint_VoteResponse) decode(pd packetDecoder, version int16) (err error) {
	n.Version = version
	if n.Version >= 1 {
		if n.NodeID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if n.Version >= 1 {
		if n.Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if n.Version >= 1 {
		if n.Port, err = pd.getUint16(); err != nil {
			return err
		}
	}

	return nil
}

type VoteResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top level error code.
	ErrorCode int16
	// Topics contains the results for each topic.
	Topics []TopicData_VoteResponse
	// NodeEndpoints contains a Endpoints for all current-leaders enumerated in PartitionData.
	NodeEndpoints []NodeEndpoint_VoteResponse
}

func (r *VoteResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt16(r.ErrorCode)

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

func (r *VoteResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_VoteResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_VoteResponse
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

func (r *VoteResponse) GetKey() int16 {
	return 52
}

func (r *VoteResponse) GetVersion() int16 {
	return r.Version
}

func (r *VoteResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *VoteResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *VoteResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *VoteResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
