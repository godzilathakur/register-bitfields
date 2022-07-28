package builder

type Builder interface {
	BuildHeader(interface{}) error
}
