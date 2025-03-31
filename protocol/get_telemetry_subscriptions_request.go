// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

type GetTelemetrySubscriptionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClientInstanceID contains a Unique id for this client instance, must be set to 0 on the first request.
	ClientInstanceID uuid.UUID
}

func (r *GetTelemetrySubscriptionsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putUUID(r.ClientInstanceID); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *GetTelemetrySubscriptionsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ClientInstanceID, err = pd.getUUID(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *GetTelemetrySubscriptionsRequest) GetKey() int16 {
	return 71
}

func (r *GetTelemetrySubscriptionsRequest) GetVersion() int16 {
	return r.Version
}

func (r *GetTelemetrySubscriptionsRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *GetTelemetrySubscriptionsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *GetTelemetrySubscriptionsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *GetTelemetrySubscriptionsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
