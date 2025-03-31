// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// CurrentLeader contains a Details of the current Raft cluster leader.
type CurrentLeader struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// LeaderID contains the replica id of the current leader or -1 if the leader is unknown.
	LeaderID int32
	// LeaderEpoch contains the latest known leader epoch.
	LeaderEpoch int32
	// Host contains the node's hostname.
	Host string
	// Port contains the node's port.
	Port int32
}

func (c *CurrentLeader) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	pe.putInt32(c.LeaderID)

	pe.putInt32(c.LeaderEpoch)

	if err := pe.putString(c.Host); err != nil {
		return err
	}

	pe.putInt32(c.Port)

	return nil
}

func (c *CurrentLeader) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.LeaderID, err = pd.getInt32(); err != nil {
		return err
	}

	if c.LeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if c.Host, err = pd.getString(); err != nil {
		return err
	}

	if c.Port, err = pd.getInt32(); err != nil {
		return err
	}

	return nil
}

type UpdateRaftVoterResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// CurrentLeader contains a Details of the current Raft cluster leader.
	CurrentLeader CurrentLeader
}

func (r *UpdateRaftVoterResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	pe.putUVarint(0)
	return nil
}

func (r *UpdateRaftVoterResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *UpdateRaftVoterResponse) GetKey() int16 {
	return 82
}

func (r *UpdateRaftVoterResponse) GetVersion() int16 {
	return r.Version
}

func (r *UpdateRaftVoterResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *UpdateRaftVoterResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *UpdateRaftVoterResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *UpdateRaftVoterResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *UpdateRaftVoterResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
