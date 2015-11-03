package togglr

func DefaulFactory(val interface{}) (Feature, bool) {
	if val == true {
		return NewStaticFeature(true), true
	} else if val == false {
		return NewStaticFeature(false), true
	}
	prop := val.(map[string]interface{})
	if feature, ok := createEnabledOnOrAfter(prop); ok {
		return feature, ok
	}
	if feature, ok := createEnabledOnOrBefore(prop); ok {
		return feature, ok
	}
	if feature, ok := createPercentFeature(prop); ok {
		return feature, ok
	}
	if feature, ok := createWeekdaysFeature(prop); ok {
		return feature, ok
	}
	return nil, false
}
