package main

import (
	"context"
	"log"
	"net/http"
	"sampleClient/gen/go-grpc/api/eventv1connect"
	eventv1 "sampleClient/gen/go/api"

	connect "github.com/bufbuild/connect-go"
	"github.com/go-chi/chi"
)

type client struct{}

func (c *client) PublishEvent(
	ctx context.Context,
	req *connect.Request[eventv1.PublishEventRequest],
) (*connect.Response[eventv1.PublishEventResponse], error) {

	res := connect.NewResponse(&eventv1.PublishEventResponse{
		Success: true,
	})
	return res, nil
}

func main() {

	mux := chi.NewRouter()

	client := &client{}
	path, handler := eventv1connect.NewEventServiceHandler(client)
	log.Printf("Registering handler for path: %s", path)
	mux.Mount(path, handler)

	log.Println("Starting connect-gRPC server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
