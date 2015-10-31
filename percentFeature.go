package togglr

import "math/rand"

func createPercentFeature(val map[string]interface{}) (Feature, bool) {
	if percentRaw, ok := val[PercentKey]; ok {
		percent := int(percentRaw.(float64))
		pred := func() bool {
			return rand.Intn(100) < percent
		}
		return predicateFeature{pred}, true
	}
	return nil, false
}
