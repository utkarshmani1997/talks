package main

import (
	"bytes"
	"fmt"
)

type Better struct {
	get      getDBFunc
	notify   notifyFunc
	publish  publishFunc
	publish2 publishFunc
}

func (b *Better) simplify(ss SomeStruct) error {
	// get msg from database

	msg := b.get(ss.Field1, ss.Field2)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "This is: %d MB.", 85)
	// Publish to pubsub
	if err := b.publish(&buf, "SomeProjectID", "SomeTopicID", msg); err != nil {
		panic(err)
	}

	// some crazy logic
	if ss.Field3 == 1 {
		// update db with some crazy data
	} else {
		// notify error on slack
		b.notify("Some error msg")
	}

	// notify that event successful to slack
	b.notify("Some success msg")
	b.publish2()
	return nil
}
