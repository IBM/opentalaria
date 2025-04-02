package utils

import (
	"testing"
)

func TestRef(t *testing.T) {
	type T int

	val := T(0)
	pointer := PtrTo(val)
	if *pointer != val {
		t.Errorf("expected %d, got %d", val, *pointer)
	}

	val = T(1)
	pointer = PtrTo(val)
	if *pointer != val {
		t.Errorf("expected %d, got %d", val, *pointer)
	}
}
