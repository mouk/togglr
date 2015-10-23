package togglr

import (
	"os"
	"testing"
)

func TestTimeConsuming(t *testing.T) {
	os.Setenv("FeatureA", "1")
	m := Features{}
	Read(&m)

	if !m.FeatureA.IsEnabled() {
		t.Fatal("FeatureA was not set")
	}

	if m.FeatureB.IsEnabled() {
		t.Fatal("FeatureB was not set correctly")
	}
}

type Features struct {
	FeatureA Feature
	FeatureB Feature
}
