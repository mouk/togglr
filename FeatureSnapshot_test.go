package togglr

import (
	"fmt"
	"testing"
)

func TestFeatureSnapshotst(t *testing.T) {

	m := SnapshotFeature{
		staticValueFeature{false},
		staticValueFeature{true},
		"Nothing",
		nil,
	}
	snaptshots := GetFeatureSpanshots(&m)
	fmt.Println(snaptshots)

	if f, ok := snaptshots["FeatureOn"]; !ok || !f {
		t.Fatal("Snapshot FeatureOn wasn't set")
	}
	if f, ok := snaptshots["FeatureOff"]; !ok || f {
		t.Fatal("Snapshot Env wasn't set")
	}

	if len(snaptshots) != 2 {
		t.Fatal("number of snapshots wrong")
	}
}

type SnapshotFeature struct {
	FeatureOff Feature
	FeatureOn  Feature
	IgnoreIt   string
	NilFeature Feature
}
