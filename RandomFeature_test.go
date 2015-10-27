package togglr

import (
	"testing"
)

func TestRandom(t *testing.T) {
	Init("data/feature.json")

	m := FeaturesRandom{}
	Read(&m)
	on := false
	off := false

	for i := 0; i < 10000; i++ {
		e := m.F1.IsEnabled()
		if e {
			on = true
		} else {
			off = true
		}
		if on && off {
			break
		}
	}

	if !(on && off) {
		t.Fatal("Not all cases were found")
	}

}

type FeaturesRandom struct {
	F1 RandomFeature
}
