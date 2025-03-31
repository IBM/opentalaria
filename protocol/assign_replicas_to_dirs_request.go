// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// PartitionData_AssignReplicasToDirsRequest contains the partitions assigned to the directory.
type PartitionData_AssignReplicasToDirsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
}

func (p *PartitionData_AssignReplicasToDirsRequest) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_AssignReplicasToDirsRequest) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_AssignReplicasToDirsRequest contains the topics assigned to the directory.
type TopicData_AssignReplicasToDirsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the ID of the assigned topic.
	TopicID uuid.UUID
	// Partitions contains the partitions assigned to the directory.
	Partitions []PartitionData_AssignReplicasToDirsRequest
}

func (t *TopicData_AssignReplicasToDirsRequest) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putUUID(t.TopicID); err != nil {
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

func (t *TopicData_AssignReplicasToDirsRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_AssignReplicasToDirsRequest, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_AssignReplicasToDirsRequest
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

// DirectoryData_AssignReplicasToDirsRequest contains the directories to which replicas should be assigned.
type DirectoryData_AssignReplicasToDirsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ID contains the ID of the directory.
	ID uuid.UUID
	// Topics contains the topics assigned to the directory.
	Topics []TopicData_AssignReplicasToDirsRequest
}

func (d *DirectoryData_AssignReplicasToDirsRequest) encode(pe packetEncoder, version int16) (err error) {
	d.Version = version
	if err := pe.putUUID(d.ID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(d.Topics)); err != nil {
		return err
	}
	for _, block := range d.Topics {
		if err := block.encode(pe, d.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (d *DirectoryData_AssignReplicasToDirsRequest) decode(pd packetDecoder, version int16) (err error) {
	d.Version = version
	if d.ID, err = pd.getUUID(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		d.Topics = make([]TopicData_AssignReplicasToDirsRequest, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_AssignReplicasToDirsRequest
			if err := block.decode(pd, d.Version); err != nil {
				return err
			}
			d.Topics[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type AssignReplicasToDirsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the ID of the requesting broker.
	BrokerID int32
	// BrokerEpoch contains the epoch of the requesting broker.
	BrokerEpoch int64
	// Directories contains the directories to which replicas should be assigned.
	Directories []DirectoryData_AssignReplicasToDirsRequest
}

func (r *AssignReplicasToDirsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.BrokerID)

	pe.putInt64(r.BrokerEpoch)

	if err := pe.putArrayLength(len(r.Directories)); err != nil {
		return err
	}
	for _, block := range r.Directories {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *AssignReplicasToDirsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.BrokerEpoch, err = pd.getInt64(); err != nil {
		return err
	}

	var numDirectories int
	if numDirectories, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numDirectories > 0 {
		r.Directories = make([]DirectoryData_AssignReplicasToDirsRequest, numDirectories)
		for i := 0; i < numDirectories; i++ {
			var block DirectoryData_AssignReplicasToDirsRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Directories[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *AssignReplicasToDirsRequest) GetKey() int16 {
	return 73
}

func (r *AssignReplicasToDirsRequest) GetVersion() int16 {
	return r.Version
}

func (r *AssignReplicasToDirsRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *AssignReplicasToDirsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *AssignReplicasToDirsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AssignReplicasToDirsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
