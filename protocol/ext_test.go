package protocol

import (
	"testing"
)

// TestRequestHeader_String_NilClientID checks that RequestHeader.String()
// safely handles a nil ClientID.
func TestRequestHeader_String_NilClientID(t *testing.T) {
	header := &RequestHeader{
		RequestApiKey:     3,
		RequestApiVersion: 8,
		CorrelationID:     34,
		ClientID:          nil,
	}

	str := header.String()
	if str == "" {
		t.Error("expected non-empty string representation")
	}
}
