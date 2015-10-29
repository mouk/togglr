package togglr

import (
	"testing"
)

func TestAllWeekdays(t *testing.T) {
	Init("data/weekday.json")

	m := WeekdaysFeature{}
	Read(&m)

	if !m.AlwaysOn.IsEnabled() {
		t.Fatal("AlwaysOn wasn't set")
	}
}

func TestNoneOfWeekdays(t *testing.T) {
	Init("data/weekday.json")

	m := WeekdaysFeature{}
	Read(&m)

	if m.AlwaysOff.IsEnabled() {
		t.Fatal("AlwaysOff wasn't set")
	}
}

type WeekdaysFeature struct {
	Sunday    Feature
	Monday    Feature
	Tuesday   Feature
	Wednesday Feature
	Thursday  Feature
	Friday    Feature
	Saturday  Feature
	AlwaysOn  Feature
	AlwaysOff Feature
}
