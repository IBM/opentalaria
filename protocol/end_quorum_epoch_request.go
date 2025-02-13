// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// ReplicaInfo contains a A sorted list of preferred candidates to start the election.
type ReplicaInfo struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// CandidateID contains the ID of the candidate replica.
	CandidateID int32
	// CandidateDirectoryID contains the directory ID of the candidate replica.
	CandidateDirectoryID uuid.UUID
}

func (p *ReplicaInfo) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 1 {
		pe.putInt32(p.CandidateID)
	}

	if p.Version >= 1 {
		if err := pe.putUUID(p.CandidateDirectoryID); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *ReplicaInfo) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Version >= 1 {
		if p.CandidateID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if p.CandidateDirectoryID, err = pd.getUUID(); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// PartitionData_EndQuorumEpochRequest contains the partitions.
type PartitionData_EndQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// LeaderID contains the current leader ID that is resigning.
	LeaderID int32
	// LeaderEpoch contains the current epoch.
	LeaderEpoch int32
	// PreferredSuccessors contains a A sorted list of preferred successors to start the election.
	PreferredSuccessors []int32
	// PreferredCandidates contains a A sorted list of preferred candidates to start the election.
	PreferredCandidates []ReplicaInfo
}

func (p *PartitionData_EndQuorumEpochRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt32(p.LeaderID)

	pe.putInt32(p.LeaderEpoch)

	if p.Version == 0 {
		if err := pe.putInt32Array(p.PreferredSuccessors); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		if err := pe.putArrayLength(len(p.PreferredCandidates)); err != nil {
			return err
		}
		for _, block := range p.PreferredCandidates {
			if err := block.encode(pe, p.Version); err != nil {
				return err
			}
		}
	}

	if p.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionData_EndQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if p.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Version == 0 {
		if p.PreferredSuccessors, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if p.Version >= 1 {
		var numPreferredCandidates int
		if numPreferredCandidates, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numPreferredCandidates > 0 {
			p.PreferredCandidates = make([]ReplicaInfo, numPreferredCandidates)
			for i := 0; i < numPreferredCandidates; i++ {
				var block ReplicaInfo
				if err := block.decode(pd, p.Version); err != nil {
					return err
				}
				p.PreferredCandidates[i] = block
			}
		}
	}

	if p.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// TopicData_EndQuorumEpochRequest contains the topics.
type TopicData_EndQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// Partitions contains the partitions.
	Partitions []PartitionData_EndQuorumEpochRequest
}

func (t *TopicData_EndQuorumEpochRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicData_EndQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicName, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_EndQuorumEpochRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_EndQuorumEpochRequest
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

// LeaderEndpoint_EndQuorumEpochRequest contains a Endpoints for the leader.
type LeaderEndpoint_EndQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port uint16
}

func (l *LeaderEndpoint_EndQuorumEpochRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (l *LeaderEndpoint_EndQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
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

type EndQuorumEpochRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClusterID contains the cluster id.
	ClusterID *string
	// Topics contains the topics.
	Topics []TopicData_EndQuorumEpochRequest
	// LeaderEndpoints contains a Endpoints for the leader.
	LeaderEndpoints []LeaderEndpoint_EndQuorumEpochRequest
}

func (r *EndQuorumEpochRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putNullableString(r.ClusterID); err != nil {
		return err
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

func (r *EndQuorumEpochRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ClusterID, err = pd.getNullableString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicData_EndQuorumEpochRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_EndQuorumEpochRequest
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
			r.LeaderEndpoints = make([]LeaderEndpoint_EndQuorumEpochRequest, numLeaderEndpoints)
			for i := 0; i < numLeaderEndpoints; i++ {
				var block LeaderEndpoint_EndQuorumEpochRequest
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

func (r *EndQuorumEpochRequest) GetKey() int16 {
	return 54
}

func (r *EndQuorumEpochRequest) GetVersion() int16 {
	return r.Version
}

func (r *EndQuorumEpochRequest) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 2
	}
	return 1
}

func (r *EndQuorumEpochRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *EndQuorumEpochRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
