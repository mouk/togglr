package togglr

import (
	"testing"
)

func TestInitWithJson(t *testing.T) {
	Init("data/feature.json")

	m := FeaturesJSON{}
	Read(&m)

	if !m.FeatureJson.IsEnabled() {
		t.Fatal("FeatureJson was not set")
	}

	if !m.FeatureWithTag.IsEnabled() {
		t.Fatal("FeatureWithTag was not set")
	}

	if m.FeatureJsonFalse.IsEnabled() {
		t.Fatal("FeatureJsonFalse was not set correctly")
	}
}

type FeaturesJSON struct {
	FeatureJson      Feature
	FeatureJsonFalse Feature
	FeatureWithTag   Feature `json:"fwt"`
}
