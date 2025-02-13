// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// PartitionData_AssignReplicasToDirsResponse contains the list of assigned partitions.
type PartitionData_AssignReplicasToDirsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition level error code.
	ErrorCode int16
}

func (p *PartitionData_AssignReplicasToDirsResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	pe.putUVarint(0)
	return nil
}

func (p *PartitionData_AssignReplicasToDirsResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicData_AssignReplicasToDirsResponse contains the list of topics and their assigned partitions.
type TopicData_AssignReplicasToDirsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the ID of the assigned topic.
	TopicID uuid.UUID
	// Partitions contains the list of assigned partitions.
	Partitions []PartitionData_AssignReplicasToDirsResponse
}

func (t *TopicData_AssignReplicasToDirsResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicData_AssignReplicasToDirsResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionData_AssignReplicasToDirsResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionData_AssignReplicasToDirsResponse
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

// DirectoryData_AssignReplicasToDirsResponse contains the list of directories and their assigned partitions.
type DirectoryData_AssignReplicasToDirsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ID contains the ID of the directory.
	ID uuid.UUID
	// Topics contains the list of topics and their assigned partitions.
	Topics []TopicData_AssignReplicasToDirsResponse
}

func (d *DirectoryData_AssignReplicasToDirsResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (d *DirectoryData_AssignReplicasToDirsResponse) decode(pd packetDecoder, version int16) (err error) {
	d.Version = version
	if d.ID, err = pd.getUUID(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		d.Topics = make([]TopicData_AssignReplicasToDirsResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_AssignReplicasToDirsResponse
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

type AssignReplicasToDirsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code.
	ErrorCode int16
	// Directories contains the list of directories and their assigned partitions.
	Directories []DirectoryData_AssignReplicasToDirsResponse
}

func (r *AssignReplicasToDirsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

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

func (r *AssignReplicasToDirsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numDirectories int
	if numDirectories, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numDirectories > 0 {
		r.Directories = make([]DirectoryData_AssignReplicasToDirsResponse, numDirectories)
		for i := 0; i < numDirectories; i++ {
			var block DirectoryData_AssignReplicasToDirsResponse
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

func (r *AssignReplicasToDirsResponse) GetKey() int16 {
	return 73
}

func (r *AssignReplicasToDirsResponse) GetVersion() int16 {
	return r.Version
}

func (r *AssignReplicasToDirsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *AssignReplicasToDirsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *AssignReplicasToDirsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AssignReplicasToDirsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
