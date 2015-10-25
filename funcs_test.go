package togglr

import (
	"os"
	"testing"
)

func TestEnvVarMatchingFieldName(t *testing.T) {
	os.Setenv("FeatureA", "1")
	os.Setenv("FeatureB", "0")
	m := Features{}
	Read(&m)

	if !m.FeatureA.IsEnabled() {
		t.Fatal("FeatureA was not set")
	}

	if m.FeatureB.IsEnabled() {
		t.Fatal("FeatureB was not set correctly")
	}
}

func TestEnvVarTagging(t *testing.T) {
	os.Setenv("feature_env", "1")
	os.Setenv("FeatureB", "0")
	m := Features{}
	Read(&m)

	if !m.FeatureEnv.IsEnabled() {
		t.Fatal("FeatureEnv was not set")
	}
}

type Features struct {
	FeatureA Feature `env:"FeatureA"`
	FeatureB Feature `env:"FeatureB"`

	FeatureEnv Feature `env:"feature_env"`
}

func TestStaticFeatures(t *testing.T) {
	m := StaticFeatures{}
	Read(&m)

	if !m.On.IsEnabled() {
		t.Fatal("On was not set")
	}
}

type StaticFeatures struct {
	On AlwaysEnabledFeature
}
