package togglr

import "math/rand"

type randomFeature struct {
}

func (random randomFeature) IsEnabled() bool {
	return rand.Int()%2 == 0
}
