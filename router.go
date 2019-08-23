package router

import (
	"errors"
)

var (
	ERR_UNINITIALIZED_ROUTER = errors.New("Uninitialized router.")
	ERR_UNSUPPORT_ROUTER     = errors.New("Unsupport router.")
)

type Router interface {
	Handler
	Get(string) Handler
	Regist(Handler) error
}
