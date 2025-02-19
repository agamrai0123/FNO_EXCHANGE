package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/ingest/models"
	"github.com/segmentio/kafka-go"
)

// Producer sends order messages to Kafka
func Producer(order *models.Order) error {
	// Validate order
	if order == nil {
		return errors.New("order cannot be nil")
	}

	// Determine topic based on product type
	var topic string
	switch order.ProductType {
	case 'O':
		topic = "Options"
	default:
		return errors.New("unsupported product type for topic selection")
	}

	// Create Kafka writer
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// Marshal order to JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return errors.New("failed to marshal order: " + err.Error())
	}

	// Write message to Kafka
	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(order.OrderReference), // Using OrderReference as key for message routing
			Value: orderJSON,
		},
	)
	if err != nil {
		return errors.New("failed to write message: " + err.Error())
	}

	return nil
}

func ProducerWithRetry(order *models.Order, maxRetries int) error {
	var err error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		err = Producer(order)
		if err == nil {
			return nil
		}

		// Exponential backoff
		backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
		time.Sleep(backoff)

		log.Printf("Retry attempt %d failed: %v. Retrying in %v",
			attempt, err, backoff)
	}
	return fmt.Errorf("failed after %d retries: %v", maxRetries, err)
}
