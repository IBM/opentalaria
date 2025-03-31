// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// PartitionData_BeginQuorumEpochRequest contains the partitions.
type PartitionData_BeginQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// VoterDirectoryID contains the directory id of the receiving replica.
	VoterDirectoryID uuid.UUID
	// LeaderID contains the ID of the newly elected leader.
	LeaderID int32
	// LeaderEpoch contains the epoch of the newly elected leader.
	LeaderEpoch int32
}

func (p *PartitionData_BeginQuorumEpochRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	if p.Version >= 1 {
		if err := pe.putUUID(p.VoterDirectoryID); err != nil {
			return err
		}
	}

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	if p.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionData_BeginQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version >= 1 {
		if p.VoterDirectoryID, err = pd.getUUID(); err != nil {
			return err
		}
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

// TopicData_BeginQuorumEpochRequest contains the topics.
type TopicData_BeginQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partitions.
	Partitions []PartitionData_BeginQuorumEpochRequest
}

func (t *TopicData_BeginQuorumEpochRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicData_BeginQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_BeginQuorumEpochRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_BeginQuorumEpochRequest
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

// LeaderEndpoint_BeginQuorumEpochRequest contains a Endpoints for the leader.
type LeaderEndpoint_BeginQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port uint16
}

func (l *LeaderEndpoint_BeginQuorumEpochRequest) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 1 {
		if err := pe.putString(l.Name); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		if err := pe.putString(l.Host); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		pe.putUint16(l.Port)
	}

	if l.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (l *LeaderEndpoint_BeginQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.Version >= 1 {
		if l.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		if l.Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		if l.Port, err = pd.getUint16(); err != nil {
			return err
		}
	}

	if l.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type BeginQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the cluster id.
	ClusterID *string
	// VoterID contains the replica id of the voter receiving the request.
	VoterID int32
	// Topics contains the topics.
	Topics []TopicData_BeginQuorumEpochRequest
	// LeaderEndpoints contains a Endpoints for the leader.
	LeaderEndpoints []LeaderEndpoint_BeginQuorumEpochRequest
}

func (r *BeginQuorumEpochRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putNullableString(r.ClusterID); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putInt32(r.VoterID)
	}

	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		if err := pe.putArrayLength(len(r.LeaderEndpoints)); err != nil {
			return err
		}
		for _, block := range r.LeaderEndpoints {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *BeginQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.VoterID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_BeginQuorumEpochRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_BeginQuorumEpochRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 1 {
		var numLeaderEndpoints int
		if numLeaderEndpoints, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numLeaderEndpoints > 0 {
			r.LeaderEndpoints = make([]LeaderEndpoint_BeginQuorumEpochRequest, numLeaderEndpoints)
			for i := 0; i < numLeaderEndpoints; i++ {
				var block LeaderEndpoint_BeginQuorumEpochRequest
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.LeaderEndpoints[i] = block
			}
		}
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *BeginQuorumEpochRequest) GetKey() int16 {
	return 53
}

func (r *BeginQuorumEpochRequest) GetVersion() int16 {
	return r.Version
}

func (r *BeginQuorumEpochRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *BeginQuorumEpochRequest) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 2
	}
	return 1
}

func (r *BeginQuorumEpochRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *BeginQuorumEpochRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
