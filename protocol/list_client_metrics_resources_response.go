// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// ClientMetricsResource contains each client metrics resource in the response.
type ClientMetricsResource struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the resource name.
	Name string
}

func (c *ClientMetricsResource) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if err := pe.putString(c.Name); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (c *ClientMetricsResource) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Name, err = pd.getString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ListClientMetricsResourcesResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ClientMetricsResources contains each client metrics resource in the response.
	ClientMetricsResources []ClientMetricsResource
}

func (r *ListClientMetricsResourcesResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.ClientMetricsResources)); err != nil {
		return err
	}
	for _, block := range r.ClientMetricsResources {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *ListClientMetricsResourcesResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numClientMetricsResources int
	if numClientMetricsResources, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numClientMetricsResources > 0 {
		r.ClientMetricsResources = make([]ClientMetricsResource, numClientMetricsResources)
		for i := 0; i < numClientMetricsResources; i++ {
			var block ClientMetricsResource
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.ClientMetricsResources[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ListClientMetricsResourcesResponse) GetKey() int16 {
	return 74
}

func (r *ListClientMetricsResourcesResponse) GetVersion() int16 {
	return r.Version
}

func (r *ListClientMetricsResourcesResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ListClientMetricsResourcesResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ListClientMetricsResourcesResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ListClientMetricsResourcesResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
