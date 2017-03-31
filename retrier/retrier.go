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
		ret := work()
		switch r.classifier.Classify(ret) {
		case Succeed, Fail:
			return ret
		case Retry:
			if retries >= len(r.backoff) {
				return ret
			}
			time.Sleep(r.backoff[retries])
			retries++
		}
	}
}
