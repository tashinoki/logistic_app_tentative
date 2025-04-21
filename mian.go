package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

type eventServiceClient struct {
	pubsubClient *pubsub.Client
	topic        *pubsub.Topic
}

func main() {

	http.HandleFunc("/api/v1/health", healthSendMessage)

	http.ListenAndServe(":8080", nil)
}

func healthSendMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, World!")

	ctx := r.Context()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	topicID := os.Getenv("PUBSUB_TOPIC")
	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		return
	}

	topic := client.Topic(topicID)

	svc := &eventServiceClient{
		pubsubClient: client,
		topic:        topic,
	}

	svc.PublishEvent(ctx, "Hello, World!")
}

func (s *eventServiceClient) PublishEvent(
	ctx context.Context,
	message string,
) error {

	log.Printf("Received request: %v", message)

	result := s.topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})

	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
