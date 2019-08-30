package router

import (
	"io"
)

type Response interface {
	io.WriteCloser
	To() string
	SetTo(string)
	Session() Session
	SetSession(Session)
}
