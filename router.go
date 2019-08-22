package router

type Router interface {
	Handler
	Regist(Handler) error
}
