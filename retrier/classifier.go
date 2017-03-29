package retrier

// Action is the type returned by a Classifier to indicate how the Retrier
// should proceed.
type Action int

const (
	Succeed Action = iota
	Fail
	Retry
)

// Classifier is the interface implemented by anything that can classify Errors
// for a Retrier
type Classifier interface {
	Classify(error) Action
}

// DefaultClassifier classifies errors in the simplest way possible. If the
// error is nil, it returns Succeed, otherwise it returns Retry.
type DefaultClassifier struct{}

// Classify implements the Classifier interface
func (c DefaultClassifier) Classify(err error) Action {
	if err == nil {
		return Succeed
	}
	return Retry
}

// WhitelistClassifier classifies errors based on whitelist. If the error is nil, it
// returns Succeed; if the error is in the whitelist, it returns Retry; otherwise
// it returns Fail.
type WhitelistClassifier []error

// Classify implements the Classifier interface
func (list WhitelistClassifier) Classify(err error) Action {
	if err == nil {
		return Succeed
	}
	for _, c := range list {
		if c == err {
			return Retry
		}
	}
	return Fail
}

// BlacklistClassifier classifies errors based on blacklist. if the error is nil, it
// returns Succeed; if the error is on the blacklist, it returns Fail; otherwise
// it returns Fail.
type BlacklistClassifier []error

// Classify implements the Classifier interface
func (list BlacklistClassifier) Classify(err error) Action {
	if err == nil {
		return Succeed
	}
	for _, c := range list {
		if c == err {
			return Fail
		}
	}
	return Retry
}
