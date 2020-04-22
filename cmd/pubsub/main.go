package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {

	projectID := flag.String("p", "my-project-id", "-p=my-project-id")
	topic := flag.String("t", "my-topic", "-t=my-topic")
	flag.Parse()

	log.Printf("ProjectID: %s", *projectID)
	log.Printf("Topic: %s", *topic)

	// Connect to PubSub.
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *projectID)
	if err != nil {
		log.Fatalf("Failed to create client for projectID: %s, error: %s", *projectID, err.Error())
	}
	// Creates the new topic.

	if _, err := client.CreateTopic(ctx, *topic); err != nil {
		log.Fatalf("Failed to create topic: %s, error: %s", *topic, err.Error())
	}
	log.Printf("Topic %v created.\n", topic)
}
