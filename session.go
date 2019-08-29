package router

import (
	"io"
)

type Session interface {
	io.ReadWriteCloser
	From() string
	Context() Context
	Request() Request
	SetRequest(Request)
	Response() Response
	SetResponse(Response)
}
