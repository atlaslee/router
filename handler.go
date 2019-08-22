package router

import (
	"io"
)

type Handler interface {
	Name() string
	Path() string
	SetPath(string)
	Handle(string, io.Reader, io.Writer) error
}
