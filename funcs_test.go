package togglr

import (
	"os"
	"testing"
)

func TestEnvVarMatchingFieldName(t *testing.T) {
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

func TestEnvVarTagging(t *testing.T) {
	os.Setenv("feature_env", "1")
	m := Features{}
	Read(&m)

	if !m.FeatureEnv.IsEnabled() {
		t.Fatal("FeatureEnv was not set")
	}
}

type Features struct {
	FeatureA Feature
	FeatureB Feature

	FeatureEnv Feature `env:"feature_env"`
}
