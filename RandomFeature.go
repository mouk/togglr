package togglr

import "math/rand"

type RandomFeature struct {
}

func (f RandomFeature) IsEnabled() bool {
	return rand.Int()%2 == 0
}
