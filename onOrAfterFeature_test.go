package togglr

import (
	"testing"
)

func TestOnOrAfterOff(t *testing.T) {
	Init("data/min_date_off.json")

	m := FeaturesOff{}
	Read(&m)

	if m.OffLong.IsEnabled() {
		t.Fatal("OffLong wasn't set")
	}

	if m.OffShort.IsEnabled() {
		t.Fatal("OffLong wasn't set")
	}
}

type FeaturesOff struct {
	OffLong  Feature
	OffShort Feature
}

func TestOnOrAfterOn(t *testing.T) {
	Init("data/min_date_on.json")

	m := FeaturesOn{}
	Read(&m)

	if !m.OnLong.IsEnabled() {
		t.Fatal("OnLong wasn't set")
	}

	if !m.OnShort.IsEnabled() {
		t.Fatal("OnLong wasn't set")
	}
}

type FeaturesOn struct {
	OnLong  Feature
	OnShort Feature
}
