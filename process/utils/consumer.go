package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/agamrai0123/FNO_EXCHANGE/process/models"
	"github.com/segmentio/kafka-go"
)

func Consumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "Options",
		StartOffset: kafka.LastOffset,
		GroupID:     "test-consumer",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		var order models.Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		log.Printf("Received message: %+v\n", order)
	}
}
