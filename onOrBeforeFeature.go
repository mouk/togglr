package togglr

import (
	"time"
)

func createEnabledOnOrBefore(val map[string]interface{}) (Feature, bool) {
	if date, ok := val[EnabledOnOrBeforeKey]; ok {
		dateRaw := date.(string)
		//TODO find a better why to specified supported formats
		if t, err := time.Parse("2006-01-02 15:04:05", dateRaw); err == nil {
			return createEnabledOnOrBeforeWithDate(t), true
		} else if t, err := time.Parse("2006-01-02", dateRaw); err == nil {
			return createEnabledOnOrBeforeWithDate(t), true
		}
	}
	return nil, false
}
func createEnabledOnOrBeforeWithDate(t time.Time) Feature {
	return newPredicateFeature(func() bool {
		now := time.Now()
		return now.Before(t) || now.Equal(t)
	})
}
