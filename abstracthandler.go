package router

import (
	"strings"
)

type AbstractHandler struct {
	name, path string
}

func (this *AbstractHandler) Name() string {
	return this.name
}

func (this *AbstractHandler) Path() string {
	return this.path
}

func (this *AbstractHandler) SetPath(path string) {
	this.path = strings.Trim(path, CUTSET)
}
