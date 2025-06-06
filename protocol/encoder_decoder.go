package protocol

import (
	"errors"
	"fmt"
)

const (
	// MaxRequestSize is the maximum size (in bytes) of any request that Sarama will attempt to send. Trying
	// to send a request larger than this will result in an PacketEncodingError. The default of 100 MiB is aligned
	// with Kafka's default `socket.request.max.bytes`, which is the largest request the broker will attempt
	// to process.
	MaxRequestSize int32 = 100 * 1024 * 1024

	// MaxResponseSize is the maximum size (in bytes) of any response that Sarama will attempt to parse. If
	// a broker returns a response message larger than this value, Sarama will return a PacketDecodingError to
	// protect the client from running out of memory. Please note that brokers do not have any natural limit on
	// the size of responses they send. In particular, they can send arbitrarily large fetch responses to consumers
	// (see https://issues.apache.org/jira/browse/KAFKA-2063).
	MaxResponseSize int32 = 100 * 1024 * 1024
)

// Encoder is the interface that wraps the basic Encode method.
// Anything implementing Encoder can be turned into bytes using Kafka's encoding rules.
type encoder interface {
	encode(pe packetEncoder) error
}

type encoderWithHeader interface {
	encoder
	headerVersion() int16
}

// Encode takes an Encoder and turns it into bytes while potentially recording metrics.
func Encode(e encoder) ([]byte, error) {
	if e == nil {
		return nil, nil
	}

	var prepEnc prepEncoder
	var realEnc realEncoder

	err := e.encode(&prepEnc)
	if err != nil {
		return nil, err
	}

	if prepEnc.length < 0 || prepEnc.length > int(MaxRequestSize) {
		return nil, fmt.Errorf("invalid request size (%d)", prepEnc.length)
	}

	realEnc.raw = make([]byte, prepEnc.length)
	err = e.encode(&realEnc)
	if err != nil {
		return nil, err
	}

	return realEnc.raw, nil
}

// decoder is the interface that wraps the basic Decode method.
// Anything implementing Decoder can be extracted from bytes using Kafka's encoding rules.
type decoder interface {
	decode(pd packetDecoder) error
}

type versionedDecoder interface {
	decode(pd packetDecoder, version int16) error
}

// decode takes bytes and a decoder and fills the fields of the decoder from the bytes,
// interpreted using Kafka's encoding rules.
func Decode(buf []byte, in decoder) error {
	if buf == nil {
		return nil
	}

	helper := realDecoder{
		raw: buf,
	}
	err := in.decode(&helper)
	if err != nil {
		return err
	}

	if helper.off != len(buf) {
		return errors.New("invalid length")
	}

	return nil
}

func VersionedDecode(buf []byte, in versionedDecoder, version int16) (int, error) {
	if buf == nil {
		return -1, nil
	}

	helper := realDecoder{
		raw: buf,
	}
	err := in.decode(&helper, version)
	if err != nil {
		return -1, err
	}

	return helper.off, nil
}
