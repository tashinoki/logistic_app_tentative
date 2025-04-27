package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

func publishMessage(projectID, topicID, msg string) (string, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("pubsub.NewClient 作成に失敗しました: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
		Attributes: map[string]string{
			"origin":   "csv-parser",
			"username": "hogehoge",
		},
	})

	id, err := result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("publish結果の取得に失敗しました (Get): %w", err)
	}

	return id, nil
}

func main() {
	projectID := os.Getenv("GCP_PROJECT_ID")
	topicID := os.Getenv("GCP_PUBSUB_TOPIC_ID")
	message := fmt.Sprintf("Hello from Go! Sent at %s", time.Now().Format(time.RFC3339))

	if projectID == "" {
		log.Fatal("エラー: Google Cloud プロジェクトIDを 'your-gcp-project-id' から実際の値に設定してください。")
	}

	log.Printf("プロジェクト '%s' のトピック '%s' にメッセージを送信します...", projectID, topicID)
	log.Printf("メッセージ内容: %s", message)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msgID, err := publishMessage(projectID, topicID, message)
		if err != nil {
			log.Fatalf("メッセージの送信に失敗しました: %v", err)
		}

		// 成功した場合、公開されたメッセージのIDをログに出力
		log.Printf("メッセージが正常に送信されました。Message ID: %s\n", msgID)
	})

	log.Println("サーバーを起動中... ポート80で待機中")
	http.ListenAndServe(":80", nil)
}
