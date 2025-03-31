// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// SnapshotId_FetchSnapshotResponse contains the snapshot endOffset and epoch fetched.
type SnapshotId_FetchSnapshotResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// EndOffset contains the snapshot end offset.
	EndOffset int64
	// Epoch contains the snapshot epoch.
	Epoch int32
}

func (s *SnapshotId_FetchSnapshotResponse) encode(pe packetEncoder, version int16) (err error) {
	s.Version = version
	pe.putInt64(s.EndOffset)

	pe.putInt32(s.Epoch)

	pe.putUVarint(0)
	return nil
}

func (s *SnapshotId_FetchSnapshotResponse) decode(pd packetDecoder, version int16) (err error) {
	s.Version = version
	if s.EndOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if s.Epoch, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// LeaderIdAndEpoch_FetchSnapshotResponse contains the leader of the partition at the time of the snapshot.
type LeaderIdAndEpoch_FetchSnapshotResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// LeaderID contains the ID of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch.
	LeaderEpoch int32
}

func (c *LeaderIdAndEpoch_FetchSnapshotResponse) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	pe.putInt32(c.LeaderID)

	pe.putInt32(c.LeaderEpoch)

	return nil
}

func (c *LeaderIdAndEpoch_FetchSnapshotResponse) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if c.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	return nil
}

// PartitionSnapshot_FetchSnapshotResponse contains the partitions to fetch.
type PartitionSnapshot_FetchSnapshotResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Index contains the partition index.
	Index int32
	// ErrorCode contains the error code, or 0 if there was no fetch error.
	ErrorCode int16
	// SnapshotID contains the snapshot endOffset and epoch fetched.
	SnapshotID SnapshotId_FetchSnapshotResponse
	// CurrentLeader contains the leader of the partition at the time of the snapshot.
	CurrentLeader LeaderIdAndEpoch_FetchSnapshotResponse
	// Size contains the total size of the snapshot.
	Size int64
	// Position contains the starting byte position within the snapshot included in the Bytes field.
	Position int64
	// UnalignedRecords contains a Snapshot data in records format which may not be aligned on an offset boundary.
	UnalignedRecords RecordBatch
}

func (p *PartitionSnapshot_FetchSnapshotResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Index)

	pe.putInt16(p.ErrorCode)

	if err := p.SnapshotID.encode(pe, p.Version); err != nil {
		return err
	}

	pe.putInt64(p.Size)

	pe.putInt64(p.Position)

	if err := p.UnalignedRecords.encode(pe, p.Version); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionSnapshot_FetchSnapshotResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Index, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	tmpSnapshotId := SnapshotId_FetchSnapshotResponse{}
	if err := tmpSnapshotId.decode(pd, p.Version); err != nil {
		return err
	}
	p.SnapshotID = tmpSnapshotId

	if p.Size, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Position, err = pd.getInt64(); err != nil {
		return err
	}

	tmpUnalignedRecords := RecordBatch{}
	if err := tmpUnalignedRecords.decode(pd, p.Version); err != nil {
		return err
	}
	p.UnalignedRecords = tmpUnalignedRecords

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TopicSnapshot_FetchSnapshotResponse contains the topics to fetch.
type TopicSnapshot_FetchSnapshotResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the topic to fetch.
	Name string
	// Partitions contains the partitions to fetch.
	Partitions []PartitionSnapshot_FetchSnapshotResponse
}

func (t *TopicSnapshot_FetchSnapshotResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
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

func (t *TopicSnapshot_FetchSnapshotResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]PartitionSnapshot_FetchSnapshotResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionSnapshot_FetchSnapshotResponse
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

// NodeEndpoint_FetchSnapshotResponse contains a Endpoints for all current-leaders enumerated in PartitionSnapshot.
type NodeEndpoint_FetchSnapshotResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// NodeID contains the ID of the associated node.
	NodeID int32
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port uint16
}

func (n *NodeEndpoint_FetchSnapshotResponse) encode(pe packetEncoder, version int16) (err error) {
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

func (n *NodeEndpoint_FetchSnapshotResponse) decode(pd packetDecoder, version int16) (err error) {
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

type FetchSnapshotResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top level response error code.
	ErrorCode int16
	// Topics contains the topics to fetch.
	Topics []TopicSnapshot_FetchSnapshotResponse
	// NodeEndpoints contains a Endpoints for all current-leaders enumerated in PartitionSnapshot.
	NodeEndpoints []NodeEndpoint_FetchSnapshotResponse
}

func (r *FetchSnapshotResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

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

func (r *FetchSnapshotResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		r.Topics = make([]TopicSnapshot_FetchSnapshotResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicSnapshot_FetchSnapshotResponse
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

func (r *FetchSnapshotResponse) GetKey() int16 {
	return 59
}

func (r *FetchSnapshotResponse) GetVersion() int16 {
	return r.Version
}

func (r *FetchSnapshotResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *FetchSnapshotResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *FetchSnapshotResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *FetchSnapshotResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *FetchSnapshotResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
