package router

import (
	"errors"
	"io"
	"strings"
)

var (
	ERR_UNINITIALIZED_ROUTER = errors.New("Uninitialized router.")
	ERR_UNSUPPORT_ROUTER     = errors.New("Unsupport router.")
)

type Default struct {
	name, path string
	handlers   map[string]Handler
}

func (this *Default) Name() string {
	return this.name
}

func (this *Default) Path() string {
	return this.path
}

func (this *Default) SetPath(path string) {
	this.path = strings.Trim(path, CUTSET)
}

func (this *Default) Handle(path string, request io.Reader, response io.Writer) error {
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

func (this *Default) Regist(handler Handler) error {
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
