package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

func complex(ss SomeStruct) error {
	// get msg from database

	msg := getFromDB(ss.Field1, ss.Field2)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "This is: %d MB.", 85)
	// Publish to pubsub
	if err := publish(&buf, "SomeProjectID", "SomeTopicID", msg); err != nil {
		panic(err)
	}

	// some crazy logic
	if ss.Field3 == 1 {
		// update db with some crazy data
	} else {
		// notify error on slack
	}

	// notify that event successful to slack
	notify()
	return nil
}

func notify() {
	// notify to slack
	//
	// slackClient.Send()
}

func getFromDB(f1, f2 string) string {
	// Logic to get some info from db
	// db.Exec("Select name from student")
	return "some msg"
}

func publish(w io.Writer, projectID, topicID, msg string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
