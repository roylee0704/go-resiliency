package retrier

import (
	"errors"
	"testing"
)

var (
	errFoo = errors.New("FOO")
	errBar = errors.New("BAR")
)

func TestDefaultClassifier(t *testing.T) {
	c := DefaultClassifier{}
	if c.Classify(nil) != Succeed {
		t.Error("default misclassified nil")
	}

	if c.Classify(errFoo) != Retry {
		t.Error("default misclassified errFoo")
	}
}

func TestWhitelistClassifier(t *testing.T) {
	c := WhitelistClassifier{errBar, errFoo}
	if c.Classify(nil) != Succeed {
		t.Error("whitelist misclassied nil")
	}
	if c.Classify(errBar) != Retry {
		t.Error("whitelist misclassified bar")
	}
	if c.Classify(errFoo) != Retry {
		t.Error("whitelist misclassified foo")
	}
}

func TestBlacklistClassifier(t *testing.T) {
	c := BlacklistClassifier{errBar, errFoo}
	if c.Classify(nil) != Succeed {
		t.Error("blacklist misclassied nil")
	}
	if c.Classify(errBar) != Fail {
		t.Error("blacklist misclassied bar")
	}
	if c.Classify(errFoo) != Fail {
		t.Error("blaclist misclassied foo")
	}
}
