// Package retrier implements the "retriable" resilliency pattern for Go.
package retrier

import "time"

// Retrier implements the "retriable" resiliency pattern, abstracting out the process of retrying a failed
// action a certain number of times with an optional back-off between each entry.
type Retrier struct {
	backoff    []time.Duration
	classifier Classifier
}

// New construct retrier with a certain number of backoffs
func New(backoff []time.Duration, classifier Classifier) *Retrier {
	return &Retrier{
		backoff:    backoff,
		classifier: classifier,
	}
}

// Run a
func (r *Retrier) Run(work func() error) error {
	retries := 0

	for {
		err := work()
		ret := r.classifier.Classify(err)

		switch ret {
		case Succeed:
			return nil
		case Fail:
			return err
		case Retry:
			if retries >= len(r.backoff) {
				return err
			}
			time.Sleep(r.backoff[retries])
			retries++
		}
	}
}
