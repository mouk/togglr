package togglr

import (
	"testing"
)

func TestOnOrBeforeOff(t *testing.T) {
	Init("data/max_date_off.json")

	m := FeaturesBeforeOff{}
	Read(&m)

	if m.OffLong.IsEnabled() {
		t.Fatal("OffLong wasn't set")
	}

	if m.OffShort.IsEnabled() {
		t.Fatal("OffLong wasn't set")
	}
}

type FeaturesBeforeOff struct {
	OffLong  Feature
	OffShort Feature
}

func TestOnOrBeforeOn(t *testing.T) {
	Init("data/max_date_on.json")

	m := FeaturesBeforeOn{}
	Read(&m)

	if !m.OnLong.IsEnabled() {
		t.Fatal("OnBeforeLong wasn't set")
	}

	if !m.OnShort.IsEnabled() {
		t.Fatal("OnBeforeShort wasn't set")
	}
}

type FeaturesBeforeOn struct {
	OnLong  Feature
	OnShort Feature
}
