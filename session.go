package router

import (
	"io"
)

type Session interface {
	io.ReadWriteCloser
	From() string
	Request() Request
	SetRequest(Request)
	Response() Response
	SetResponse(Response)
}
