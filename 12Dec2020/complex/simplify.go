package main

import (
	"bytes"
	"fmt"
	"io"
)

type getDBFunc func(string, string) string
type notifyFunc func(string)
type publishFunc func(io.Writer, string, string, string) error

func simplify(ss SomeStruct, get getDBFunc, ntfy notifyFunc, pub publishFunc) error {
	// get msg from database

	msg := get(ss.Field1, ss.Field2)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "This is: %d MB.", 85)
	// Publish to pubsub
	if err := pub(&buf, "SomeProjectID", "SomeTopicID", msg); err != nil {
		panic(err)
	}

	// some crazy logic
	if ss.Field3 == 1 {
		// update db with some crazy data
	} else {
		// notify error on slack
		ntfy("Some error msg")
	}

	// notify that event successful to slack
	ntfy("Some success msg")
	return nil
}
