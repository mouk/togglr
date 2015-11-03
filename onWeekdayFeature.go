package togglr

import (
	"strings"
	"time"
)

func createWeekdaysFeature(val map[string]interface{}) (Feature, bool) {
	if daysRaw, ok := val[WeekdaysKey]; ok {
		activeOnWeekdays := make(map[time.Weekday]bool)
		days := strings.Split(daysRaw.(string), " | ")

		for i := 0; i < 7; i = i + 1 {
			activeOnWeekdays[time.Weekday(i)] = false
			for _, day := range days {
				if time.Weekday(i).String() == day {
					activeOnWeekdays[time.Weekday(i)] = true
				}
			}
		}
		pred := func() bool {
			return activeOnWeekdays[time.Now().Weekday()]
		}
		return newPredicateFeature(pred), true
	}
	return nil, false
}
