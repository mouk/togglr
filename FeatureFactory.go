package togglr

type FeatureFactory (func(data interface{}) (Feature, bool))

var factories = make([]FeatureFactory, 0, 2)

//Add new FeatureFactory
func InjectFeatureFactory(factory FeatureFactory) {
	factories = append(factories, factory)
}
