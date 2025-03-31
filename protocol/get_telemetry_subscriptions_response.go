// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	uuid "github.com/google/uuid"
	"time"
)

type GetTelemetrySubscriptionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ClientInstanceID contains a Assigned client instance id if ClientInstanceId was 0 in the request, else 0.
	ClientInstanceID uuid.UUID
	// SubscriptionID contains a Unique identifier for the current subscription set for this client instance.
	SubscriptionID int32
	// AcceptedCompressionTypes contains a Compression types that broker accepts for the PushTelemetryRequest.
	AcceptedCompressionTypes []int8
	// PushIntervalMs contains a Configured push interval, which is the lowest configured interval in the current subscription set.
	PushIntervalMs int32
	// TelemetryMaxBytes contains the maximum bytes of binary data the broker accepts in PushTelemetryRequest.
	TelemetryMaxBytes int32
	// DeltaTemporality contains a Flag to indicate monotonic/counter metrics are to be emitted as deltas or cumulative values.
	DeltaTemporality bool
	// RequestedMetrics contains a Requested metrics prefix string match. Empty array: No metrics subscribed, Array[0] empty string: All metrics subscribed.
	RequestedMetrics []string
}

func (r *GetTelemetrySubscriptionsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putUUID(r.ClientInstanceID); err != nil {
		return err
	}

	pe.putInt32(r.SubscriptionID)

	if err := pe.putInt8Array(r.AcceptedCompressionTypes); err != nil {
		return err
	}

	pe.putInt32(r.PushIntervalMs)

	pe.putInt32(r.TelemetryMaxBytes)

	pe.putBool(r.DeltaTemporality)

	if err := pe.putStringArray(r.RequestedMetrics); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *GetTelemetrySubscriptionsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ClientInstanceID, err = pd.getUUID(); err != nil {
		return err
	}

	if r.SubscriptionID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.AcceptedCompressionTypes, err = pd.getInt8Array(); err != nil {
		return err
	}

	if r.PushIntervalMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.TelemetryMaxBytes, err = pd.getInt32(); err != nil {
		return err
	}

	if r.DeltaTemporality, err = pd.getBool(); err != nil {
		return err
	}

	if r.RequestedMetrics, err = pd.getStringArray(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *GetTelemetrySubscriptionsResponse) GetKey() int16 {
	return 71
}

func (r *GetTelemetrySubscriptionsResponse) GetVersion() int16 {
	return r.Version
}

func (r *GetTelemetrySubscriptionsResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *GetTelemetrySubscriptionsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *GetTelemetrySubscriptionsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *GetTelemetrySubscriptionsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *GetTelemetrySubscriptionsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
