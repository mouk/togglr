package togglr

import (
	"testing"
)

func TestPredicate(t *testing.T) {

	predOn := newPredicateFeature(func() bool { return true })
	predOff := newPredicateFeature(func() bool { return false })

	if !predOn.IsEnabled() {
		t.Fatal("predOn")
	}

	if predOff.IsEnabled() {
		t.Fatal("predOff")
	}
}
