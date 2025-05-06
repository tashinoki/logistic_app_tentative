package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	connect "github.com/bufbuild/connect-go"
	"github.com/go-chi/chi"
	"github.com/tashinoki/logistic_app_tentative/proto/go-grpc/api/eventv1connect"
	eventv1 "github.com/tashinoki/logistic_app_tentative/proto/go/api"
)

type client struct{}

func (c *client) PublishEvent(
	ctx context.Context,
	req *connect.Request[eventv1.PublishEventRequest],
) (*connect.Response[eventv1.PublishEventResponse], error) {

	serverAddr := os.Getenv("SERVER_ADDR")

	if serverAddr == "" {
		serverAddr = "http://localhost:8000"
	}

	client := eventv1connect.NewEventReceiverServiceClient(
		&http.Client{
			Transport: &http.Transport{},
			Timeout:   15 * time.Second,
		},
		serverAddr,
		connect.WithGRPC(), // gRPC プロトコルを使用するためのオプション
		// connect.WithGRPCWeb(), // gRPC-Web を使いたい場合はこちら
	)

	servReq := connect.NewRequest(&eventv1.ReceiveEventRequest{
		Message: "from client",
	})
	client.ReceiveEvent(ctx, servReq)

	res := connect.NewResponse(&eventv1.PublishEventResponse{
		Success: true,
	})
	return res, nil
}

func main() {

	mux := chi.NewRouter()

	client := &client{}
	path, handler := eventv1connect.NewEventPublisherServiceHandler(client)
	log.Printf("Registering handler for path: %s", path)
	mux.Mount(path, handler)

	log.Println("Starting connect-gRPC server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
