// protocol has been generated from message format json - DO NOT EDIT
package protocol

// PartitionData_EndQuorumEpochResponse contains the partition data.
type PartitionData_EndQuorumEpochResponse struct {
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
}

func (p *PartitionData_EndQuorumEpochResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	if p.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionData_EndQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
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

	if p.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// TopicData_EndQuorumEpochResponse contains the topic data.
type TopicData_EndQuorumEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partition data.
	Partitions []PartitionData_EndQuorumEpochResponse
}

func (t *TopicData_EndQuorumEpochResponse) encode(pe packetEncoder, version int16) (err error) {
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

	if t.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *TopicData_EndQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_EndQuorumEpochResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_EndQuorumEpochResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// NodeEndpoint_EndQuorumEpochResponse contains a Endpoints for all leaders enumerated in PartitionData.
type NodeEndpoint_EndQuorumEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// NodeID contains the ID of the associated node.
	NodeID int32
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port uint16
}

func (n *NodeEndpoint_EndQuorumEpochResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (n *NodeEndpoint_EndQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
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

type EndQuorumEpochResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top level error code.
	ErrorCode int16
	// Topics contains the topic data.
	Topics []TopicData_EndQuorumEpochResponse
	// NodeEndpoints contains a Endpoints for all leaders enumerated in PartitionData.
	NodeEndpoints []NodeEndpoint_EndQuorumEpochResponse
}

func (r *EndQuorumEpochResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *EndQuorumEpochResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_EndQuorumEpochResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_EndQuorumEpochResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *EndQuorumEpochResponse) GetKey() int16 {
	return 54
}

func (r *EndQuorumEpochResponse) GetVersion() int16 {
	return r.Version
}

func (r *EndQuorumEpochResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *EndQuorumEpochResponse) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 1
	}
	return 0
}

func (r *EndQuorumEpochResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *EndQuorumEpochResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
