package togglr

func DefaulFactory(val interface{}) (Feature, bool) {
	if val == true {
		return staticValueFeature{true}, true
	} else if val == false {
		return staticValueFeature{false}, true
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
