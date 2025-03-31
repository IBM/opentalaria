// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ShareGroupDescribeRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupIds contains the ids of the groups to describe.
	GroupIds []string
	// IncludeAuthorizedOperations contains a Whether to include authorized operations.
	IncludeAuthorizedOperations bool
}

func (r *ShareGroupDescribeRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putStringArray(r.GroupIds); err != nil {
		return err
	}

	pe.putBool(r.IncludeAuthorizedOperations)

	pe.putUVarint(0)
	return nil
}

func (r *ShareGroupDescribeRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.GroupIds, err = pd.getStringArray(); err != nil {
		return err
	}

	if r.IncludeAuthorizedOperations, err = pd.getBool(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ShareGroupDescribeRequest) GetKey() int16 {
	return 77
}

func (r *ShareGroupDescribeRequest) GetVersion() int16 {
	return r.Version
}

func (r *ShareGroupDescribeRequest) SetVersion(version int16) {
	r.Version = version
}

func (r *ShareGroupDescribeRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ShareGroupDescribeRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ShareGroupDescribeRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
