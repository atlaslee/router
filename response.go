package router

import (
	"io"
)

type Response interface {
	io.WriteCloser
	To() string
	Session() Session
	SetSession(Session)
}
