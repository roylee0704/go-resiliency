package retrier

import (
	"testing"
	"time"
)

func TestRetrier(t *testing.T) {

	r := New(ExponentialBackoff(2, 1*time.Second), DefaultClassifier{})

	r.Run(func() error {
		return errBar
	})
}
