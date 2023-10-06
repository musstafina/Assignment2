package decorators

type PlaylistDecorator interface{
	Decorate(playlist string) string
}