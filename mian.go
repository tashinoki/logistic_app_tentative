package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

type eventServiceClient struct {
	pubsubClient *pubsub.Client
	topic        *pubsub.Topic
}

type PushRequest struct {
	Message struct {
		Data        string `json:"data"`
		MessageID   string `json:"messageId"`
		PublishTime string `json:"publishTime"`
	} `json:"message"`
	Subscription string `json:"subscription"` // サブスクリプション名
}

func main() {

	http.HandleFunc("/api/v1/health", healthSendMessage)
	http.HandleFunc("/api/v1/subscribe", subscribeHandler)

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

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTメソッドのみ許可されています", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディの読み込み
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("リクエストボディの読み込みに失敗しました: %v", err)
		http.Error(w, "リクエストボディの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// JSONのパース
	var pushReq PushRequest
	if err := json.Unmarshal(body, &pushReq); err != nil {
		log.Printf("JSONのパースに失敗しました: %v", err)
		http.Error(w, "JSONのパースに失敗しました", http.StatusBadRequest) // Pub/Subが再試行するように400を返す
		return
	}

	// Base64デコード
	data, err := base64.StdEncoding.DecodeString(pushReq.Message.Data)
	if err != nil {
		log.Printf("Base64デコードに失敗しました: %v", err)
		http.Error(w, "Base64デコードに失敗しました", http.StatusBadRequest) // Pub/Subが再試行するように400を返す
		return
	}

	// 受信したメッセージデータの処理
	log.Printf("Pub/Subからメッセージを受信しました:")
	log.Printf("  MessageID: %s", pushReq.Message.MessageID)
	log.Printf("  Subscription: %s", pushReq.Subscription)
	log.Printf("  Data: %s", string(data)) // デコードされたメッセージデータ

	// Pub/Subに成功を通知するために200 OKを返す
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "メッセージを受信し、処理しました。")
}
