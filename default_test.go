package router

import (
	"errors"
	"io"
	"testing"
)

var (
	ERR_TEST_ERR = errors.New("Test error.")
)

type TestRouter struct {
	Default
}

func (this *TestRouter) Handle(string, io.Reader, io.Writer) error {
	return ERR_TEST_ERR
}

func TestDefaultHandle(t *testing.T) {
	root := &Default{}
	child := &Default{name: "CLD"}
	root.Regist(child)
	if child.Path() != "CLD" {
		t.Fatal("child.Path failed.")
	}

	test := &TestRouter{Default{name: "TEST"}}
	child.Regist(test)
	if test.Path() != "CLD/TEST" {
		t.Fatal("test.Path failed.")
	}

	err := root.Handle("CLD/TEST", nil, nil)
	if err != ERR_TEST_ERR {
		t.Fatal("root.Handle CLD/TEST failed.")
	}

	err = root.Handle("CLD/TEST1", nil, nil)
	if err != ERR_UNSUPPORT_ROUTER {
		t.Fatal("root.Handle CLD/TEST1 failed.")
	}
}
