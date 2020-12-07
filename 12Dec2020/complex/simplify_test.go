package main

import (
	"io"
	"testing"
)

func TestSimplify(t *testing.T) {
	var (
		get  getDBFunc
		pub  publishFunc
		ntfy notifyFunc
	)
	s := SomeStruct{}
	get = func(str1, str2 string) string {
		return "hello"
	}
	pub = func(w io.Writer, x string, y string, z string) error {
		return nil
	}

	ntfy = func(msg string) {

	}
	simplify(s, get, ntfy, pub)
}
