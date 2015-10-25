package togglr

type AlwaysEnabledFeature struct {
}

func (f AlwaysEnabledFeature) IsEnabled() bool {
	return true
}
