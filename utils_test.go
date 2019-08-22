package router

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	s := "ABC/DEF/"
	prefix, left := Prefix(s)

	if prefix != "ABC" {
		t.Fatal("Prefix incorrect:", prefix)
	}

	if left != "DEF" {
		t.Fatal("Left incorrect:", left)
	}

	s = "/ABC/"
	prefix, left = Prefix(s)

	if prefix != "ABC" {
		t.Fatal("Prefix incorrect:", prefix)
	}

	if left != "" {
		t.Fatal("Left incorrect:", left)
	}
}
