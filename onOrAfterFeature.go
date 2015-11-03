package togglr

import (
	"time"
)

func createEnabledOnOrAfter(val map[string]interface{}) (Feature, bool) {
	if date, ok := val[EnabledOnOrAfterKey]; ok {
		dateRaw := date.(string)
		//TODO find a better why to specified supported formats
		if t, err := time.Parse("2006-01-02 15:04:05", dateRaw); err == nil {
			return createEnabledOnOrAfterWithDate(t), true
		} else if t, err := time.Parse("2006-01-02", dateRaw); err == nil {
			return createEnabledOnOrAfterWithDate(t), true
		}
	}
	return nil, false
}
func createEnabledOnOrAfterWithDate(t time.Time) Feature {
	return newPredicateFeature(func() bool {
		now := time.Now()
		return now.After(t) || now.Equal(t)
	})
}
