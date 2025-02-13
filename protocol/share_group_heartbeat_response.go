// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

// TopicPartitions_ShareGroupHeartbeatResponse contains a
type TopicPartitions_ShareGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic ID.
	TopicID uuid.UUID
	// Partitions contains the partitions.
	Partitions []int32
}

func (t *TopicPartitions_ShareGroupHeartbeatResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putUUID(t.TopicID); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicPartitions_ShareGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TopicID, err = pd.getUUID(); err != nil {
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

// Assignment_ShareGroupHeartbeatResponse contains a null if not provided; the assignment otherwise.
type Assignment_ShareGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicPartitions contains the partitions assigned to the member.
	TopicPartitions []TopicPartitions_ShareGroupHeartbeatResponse
}

func (a *Assignment_ShareGroupHeartbeatResponse) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putArrayLength(len(a.TopicPartitions)); err != nil {
		return err
	}
	for _, block := range a.TopicPartitions {
		if err := block.encode(pe, a.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (a *Assignment_ShareGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	var numTopicPartitions int
	if numTopicPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopicPartitions > 0 {
		a.TopicPartitions = make([]TopicPartitions_ShareGroupHeartbeatResponse, numTopicPartitions)
		for i := 0; i < numTopicPartitions; i++ {
			var block TopicPartitions_ShareGroupHeartbeatResponse
			if err := block.decode(pd, a.Version); err != nil {
				return err
			}
			a.TopicPartitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ShareGroupHeartbeatResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// MemberID contains the member ID is generated by the consumer and provided by the consumer for all requests.
	MemberID *string
	// MemberEpoch contains the member epoch.
	MemberEpoch int32
	// HeartbeatIntervalMs contains the heartbeat interval in milliseconds.
	HeartbeatIntervalMs int32
	// Assignment contains a null if not provided; the assignment otherwise.
	Assignment Assignment_ShareGroupHeartbeatResponse
}

func (r *ShareGroupHeartbeatResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putNullableString(r.MemberID); err != nil {
		return err
	}

	pe.putInt32(r.MemberEpoch)

	pe.putInt32(r.HeartbeatIntervalMs)

	if err := r.Assignment.encode(pe, r.Version); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *ShareGroupHeartbeatResponse) decode(pd packetDecoder, version int16) (err error) {
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

	if r.MemberID, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.MemberEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.HeartbeatIntervalMs, err = pd.getInt32(); err != nil {
		return err
	}

	tmpAssignment := Assignment_ShareGroupHeartbeatResponse{}
	if err := tmpAssignment.decode(pd, r.Version); err != nil {
		return err
	}
	r.Assignment = tmpAssignment

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ShareGroupHeartbeatResponse) GetKey() int16 {
	return 76
}

func (r *ShareGroupHeartbeatResponse) GetVersion() int16 {
	return r.Version
}

func (r *ShareGroupHeartbeatResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ShareGroupHeartbeatResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ShareGroupHeartbeatResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ShareGroupHeartbeatResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
