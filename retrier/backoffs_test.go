package retrier

import (
	"testing"
	"time"
)

func TestConstantBackoff(t *testing.T) {
	b := ConstantBackoff(1, 10*time.Millisecond)
	if len(b) != 1 {
		t.Error("incorrect length")
	}
	for i := range b {
		if b[i] != 10*time.Millisecond {
			t.Error("incorrect value at", i)
		}
	}
}

func TestExponentialBackoff(t *testing.T) {
	b := ExponentialBackoff(1, 10*time.Millisecond)
	if len(b) != 1 {
		t.Error("incorrect length")
	}

	next := 10 * time.Millisecond
	for i := range b {
		if b[i] != next {
			t.Error("incorrect value at", i)
		}
		next *= 2
	}
}
