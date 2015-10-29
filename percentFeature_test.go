package togglr

import (
	"testing"
)

func TestPercent100(t *testing.T) {
	Init("data/percent_feature.json")

	m := FeaturePercent{}
	Read(&m)

	if !m.AlwayOn.IsEnabled() {
		t.Fatal("AlwayOn")
	}
}

func TestPercent0(t *testing.T) {
	Init("data/percent_feature.json")

	m := FeaturePercent{}
	Read(&m)

	if m.AlwaysOff.IsEnabled() {
		t.Fatal("AlwaysOff")
	}
}

type FeaturePercent struct {
	AlwayOn    Feature
	AlwaysOff  Feature
	FiftyFifty Feature
}
