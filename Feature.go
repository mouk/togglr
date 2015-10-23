package togglr

type Feature interface {
	IsEnabled() bool
}
