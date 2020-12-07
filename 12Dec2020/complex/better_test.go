package main

import (
	"io"
	"testing"
)

func TestBetter(t *testing.T) {
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

	better := Better{
		get:     get,
		publish: pub,
		notify:  ntfy,
	}
	better.simplify(s)
}
