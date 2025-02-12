// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

type PushTelemetryRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClientInstanceID contains a Unique id for this client instance.
	ClientInstanceID uuid.UUID
	// SubscriptionID contains a Unique identifier for the current subscription.
	SubscriptionID int32
	// Terminating contains a Client is terminating the connection.
	Terminating bool
	// CompressionType contains a Compression codec used to compress the metrics.
	CompressionType int8
	// Metrics contains a Metrics encoded in OpenTelemetry MetricsData v1 protobuf format.
	Metrics []byte
}

func (r *PushTelemetryRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putUUID(r.ClientInstanceID); err != nil {
		return err
	}

	pe.putInt32(r.SubscriptionID)

	pe.putBool(r.Terminating)

	pe.putInt8(r.CompressionType)

	if err := pe.putBytes(r.Metrics); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *PushTelemetryRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClientInstanceID, err = pd.getUUID(); err != nil {
		return err
	}

	if r.SubscriptionID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Terminating, err = pd.getBool(); err != nil {
		return err
	}

	if r.CompressionType, err = pd.getInt8(); err != nil {
		return err
	}

	if r.Metrics, err = pd.getBytes(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *PushTelemetryRequest) GetKey() int16 {
	return 72
}

func (r *PushTelemetryRequest) GetVersion() int16 {
	return r.Version
}

func (r *PushTelemetryRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *PushTelemetryRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *PushTelemetryRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
