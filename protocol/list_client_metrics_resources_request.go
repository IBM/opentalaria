// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ListClientMetricsResourcesRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
}

func (r *ListClientMetricsResourcesRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putUVarint(0)
	return nil
}

func (r *ListClientMetricsResourcesRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ListClientMetricsResourcesRequest) GetKey() int16 {
	return 74
}

func (r *ListClientMetricsResourcesRequest) GetVersion() int16 {
	return r.Version
}

func (r *ListClientMetricsResourcesRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *ListClientMetricsResourcesRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ListClientMetricsResourcesRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ListClientMetricsResourcesRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
