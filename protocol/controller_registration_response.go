// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

type ControllerRegistrationResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the response error code.
	ErrorCode int16
	// ErrorMessage contains the response error message, or null if there was no error.
	ErrorMessage *string
}

func (r *ControllerRegistrationResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *ControllerRegistrationResponse) decode(pd packetDecoder, version int16) (err error) {
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

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ControllerRegistrationResponse) GetKey() int16 {
	return 70
}

func (r *ControllerRegistrationResponse) GetVersion() int16 {
	return r.Version
}

func (r *ControllerRegistrationResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ControllerRegistrationResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ControllerRegistrationResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ControllerRegistrationResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
