// protocol has been generated from message format json - DO NOT EDIT
package protocol

type DescribeClusterRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// IncludeClusterAuthorizedOperations contains a Whether to include cluster authorized operations.
	IncludeClusterAuthorizedOperations bool
	// EndpointType contains the endpoint type to describe. 1=brokers, 2=controllers.
	EndpointType int8
	// IncludeFencedBrokers contains a Whether to include fenced brokers when listing brokers.
	IncludeFencedBrokers bool
}

func (r *DescribeClusterRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putBool(r.IncludeClusterAuthorizedOperations)

	if r.Version >= 1 {
		pe.putInt8(r.EndpointType)
	}

	if r.Version >= 2 {
		pe.putBool(r.IncludeFencedBrokers)
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeClusterRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.IncludeClusterAuthorizedOperations, err = pd.getBool(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.EndpointType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		if r.IncludeFencedBrokers, err = pd.getBool(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeClusterRequest) GetKey() int16 {
	return 60
}

func (r *DescribeClusterRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeClusterRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *DescribeClusterRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeClusterRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *DescribeClusterRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
