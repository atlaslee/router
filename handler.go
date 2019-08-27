package router

type Handler interface {
	Name() string
	Path() string
	SetPath(string)
	Handle(string, Request, Response) error
}
