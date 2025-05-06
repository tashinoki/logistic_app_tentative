package main

import (
	"context"
	"log"
	"net/http"

	eventv1 "github.com/tashinoki/logistic_app_tentative/proto/go/api"

	"github.com/tashinoki/logistic_app_tentative/proto/go-grpc/api/eventv1connect"

	connect "github.com/bufbuild/connect-go"
	"github.com/go-chi/chi"
)

type client struct{}

func (c *client) ReceiveEvent(
	ctx context.Context,
	req *connect.Request[eventv1.ReceiveEventRequest],
) (*connect.Response[eventv1.ReceiveEventResponse], error) {

	res := connect.NewResponse(&eventv1.ReceiveEventResponse{
		Success: true,
	})
	return res, nil
}

func main() {

	mux := chi.NewRouter()

	client := &client{}
	path, handler := eventv1connect.NewEventReceiverServiceHandler(client)
	log.Printf("Registering handler for path: %s", path)
	mux.Mount(path, handler)

	log.Println("Starting connect-gRPC server on :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
