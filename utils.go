package router

import (
	"strings"
)

const (
	SEPARATOR = "/"
	CUTSET    = "/ "
)

func Split(path string) []string {
	return strings.Split(path, SEPARATOR)
}

func Join(pieces []string) string {
	return strings.Join(pieces, SEPARATOR)
}

func Prefix(path string) (prefix string, left string) {
	path = strings.Trim(path, CUTSET)

	i := strings.Index(path, SEPARATOR)
	if i == -1 {
		return path, ""
	}

	return path[:i], path[i+1:]
}
