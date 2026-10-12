package protocol

import (
	"errors"
	"testing"
)

// TestGetCompactString_OutOfBounds checks that getCompactString returns ErrInsufficientData
// when the wire length exceeds the remaining buffer bytes, rather than panicking with slice bounds out of range.
func TestGetCompactString_OutOfBounds(t *testing.T) {
	// 0xe8, 0x07 is uvarint 1000 -> length 999, but only 2 bytes ("AA") remain in buffer
	raw := []byte{0xe8, 0x07, 0x41, 0x41}
	rd := &realDecoder{raw: raw}

	str, err := rd.getCompactString()
	if err == nil {
		t.Errorf("expected error for out-of-bounds length, got string %q", str)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetCompactNullableString_OutOfBounds checks that getCompactNullableString returns ErrInsufficientData
// when wire length exceeds remaining buffer bytes, rather than panicking with slice bounds out of range.
func TestGetCompactNullableString_OutOfBounds(t *testing.T) {
	// uvarint 1000 -> length 999, with only 2 bytes remaining
	raw := []byte{0xe8, 0x07, 0x41, 0x41}
	rd := &realDecoder{raw: raw}

	str, err := rd.getCompactNullableString()
	if err == nil {
		t.Errorf("expected error for out-of-bounds length, got string %v", str)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetCompactArrayLength_OutOfBounds checks that getCompactArrayLength rejects array lengths
// exceeding remaining buffer bytes with ErrInsufficientData.
func TestGetCompactArrayLength_OutOfBounds(t *testing.T) {
	// 10-byte uvarint for 2^64-1
	raw := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	rd := &realDecoder{raw: raw}

	length, err := rd.getCompactArrayLength()
	if err == nil {
		t.Errorf("expected ErrInsufficientData for out-of-bounds array length, got length: %d", length)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetCompactStringArray_OutOfBounds checks that getCompactStringArray returns an error
// when array length is out of bounds, rather than panicking with makeslice len out of range.
func TestGetCompactStringArray_OutOfBounds(t *testing.T) {
	// 10-byte uvarint for 2^64-1
	raw := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	rd := &realDecoder{raw: raw}

	arr, err := rd.getCompactStringArray()
	if err == nil {
		t.Errorf("expected error for out-of-bounds array length, got %v", arr)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetCompactInt32Array_OutOfBounds checks that getCompactInt32Array returns an error
// when array length is out of bounds, rather than panicking with makeslice len out of range.
func TestGetCompactInt32Array_OutOfBounds(t *testing.T) {
	// 10-byte uvarint for 2^64-1
	raw := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	rd := &realDecoder{raw: raw}

	arr, err := rd.getCompactInt32Array()
	if err == nil {
		t.Errorf("expected error for out-of-bounds array length, got %v", arr)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetCompactInt16Array_OutOfBounds checks that getCompactInt16Array returns an error
// when array length is out of bounds, rather than panicking with makeslice len out of range.
func TestGetCompactInt16Array_OutOfBounds(t *testing.T) {
	raw := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	rd := &realDecoder{raw: raw}

	arr, err := rd.getCompactInt16Array()
	if err == nil {
		t.Errorf("expected error for out-of-bounds array length, got %v", arr)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetCompactInt8Array_OutOfBounds checks that getCompactInt8Array returns an error
// when array length is out of bounds, rather than panicking with makeslice len out of range.
func TestGetCompactInt8Array_OutOfBounds(t *testing.T) {
	raw := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	rd := &realDecoder{raw: raw}

	arr, err := rd.getCompactInt8Array()
	if err == nil {
		t.Errorf("expected error for out-of-bounds array length, got %v", arr)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestGetStringArray_OutOfBounds checks that getStringArray rejects unbounded lengths
// before slice allocation.
func TestGetStringArray_OutOfBounds(t *testing.T) {
	// 4 bytes encoding int32 0x7FFFFFFF (2147483647)
	raw := []byte{0x7f, 0xff, 0xff, 0xff}
	rd := &realDecoder{raw: raw}

	arr, err := rd.getStringArray()
	if err == nil {
		t.Errorf("expected error for out-of-bounds array length, got %v", arr)
	}
	if !errors.Is(err, ErrInsufficientData) {
		t.Errorf("expected ErrInsufficientData, got: %v", err)
	}
}

// TestFlexibleDecoder_ApiVersionsRequest_OutOfBounds checks that decoding an ApiVersionsRequest v3
// with an oversized compact string returns an error and does not panic.
func TestFlexibleDecoder_ApiVersionsRequest_OutOfBounds(t *testing.T) {
	// ApiVersionsRequest v3 payload with ClientSoftwareName having compact string length 1000 (0xe8, 0x07) followed by "AA"
	req := &ApiVersionsRequest{}
	payload := []byte{0xe8, 0x07, 0x41, 0x41}
	_, err := VersionedDecode(payload, req, 3)
	if err == nil {
		t.Error("expected decode error, got nil")
	}
}
