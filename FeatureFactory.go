package togglr

type FeatureFactory (func(data interface{}) (Feature, bool))
