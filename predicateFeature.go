package togglr

type predicateFeature struct {
	predicate (func() bool)
}

func newPredicateFeature(pred func() bool) Feature {
	return &predicateFeature{pred}
}

func (f *predicateFeature) IsEnabled() bool {
	return f.predicate()
}
