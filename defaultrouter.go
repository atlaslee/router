package router

type DefaultRouter struct {
	AbstractHandler
	handlers map[string]Handler
}

func (this *DefaultRouter) Get(path string) (handler Handler) {
	if this.handlers == nil {
		return nil
	}

	handler, _ = this.handlers[path]
	return
}

func (this *DefaultRouter) Handle(path string, request Request, response Response) error {
	if this.handlers == nil {
		return ERR_UNINITIALIZED_ROUTER
	}

	prefix, left := Prefix(path)
	if prefix == "" {
		return ERR_UNSUPPORT_ROUTER
	}

	handler, ok := this.handlers[prefix]
	if !ok {
		return ERR_UNSUPPORT_ROUTER
	}

	return handler.Handle(left, request, response)
}

func (this *DefaultRouter) Regist(handler Handler) error {
	if this.handlers == nil {
		this.handlers = map[string]Handler{}
	}

	if handler == nil {
		return ERR_UNINITIALIZED_ROUTER
	}

	this.handlers[handler.Name()] = handler
	handler.SetPath(this.Path() + SEPARATOR + handler.Name())
	return nil
}

func DefaultRouterNew(name string) *DefaultRouter {
	return &DefaultRouter{
		AbstractHandler: AbstractHandler{
			name: name},
		handlers: map[string]Handler{}}
}
