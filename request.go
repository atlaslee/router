package router

import (
	"io"
)

type Request interface {
	io.ReadCloser
	From() string
	Session() Session
}
