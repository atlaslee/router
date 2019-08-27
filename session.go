package router

import (
	"io"
)

type Session interface {
	io.ReadWriteCloser
	From() string
	Request() Request
	Response() Response
}
