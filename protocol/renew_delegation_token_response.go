// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type RenewDelegationTokenResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ExpiryTimestampMs contains the timestamp in milliseconds at which this token expires.
	ExpiryTimestampMs int64
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
}

func (r *RenewDelegationTokenResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	pe.putInt64(r.ExpiryTimestampMs)

	pe.putInt32(r.ThrottleTimeMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *RenewDelegationTokenResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ExpiryTimestampMs, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *RenewDelegationTokenResponse) GetKey() int16 {
	return 39
}

func (r *RenewDelegationTokenResponse) GetVersion() int16 {
	return r.Version
}

func (r *RenewDelegationTokenResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *RenewDelegationTokenResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *RenewDelegationTokenResponse) IsValidVersion() bool {
	return r.Version >= 1 && r.Version <= 2
}

func (r *RenewDelegationTokenResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *RenewDelegationTokenResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
