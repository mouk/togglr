package togglr

import (
	"testing"
)

func TestFactory(t *testing.T) {
	Init("data/custom_feature.json")
	InjectFeatureFactory(customFactory)
	m := customFeatures{}
	Read(&m)

	if _, ok := m.CustomFeature.(randomFeature); !ok {
		t.Fatal("The Custom fatory wasn't used")
	}

}

func customFactory(data interface{}) (Feature, bool) {
	if m, ok := data.(map[string]interface{}); ok && m["customKey"] == "CustomValue" {
		return randomFeature{}, true
	}
	return nil, false

}

type customFeatures struct {
	CustomFeature Feature
}
