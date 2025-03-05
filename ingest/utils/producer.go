package utils

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/ingest/models"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaWriters = make(map[string]*kafka.Writer) // Map to store writers per topic
	writerMutex  sync.Mutex                       // Mutex for safe concurrent access
)

func getKafkaWriter(topic string) *kafka.Writer {
	writerMutex.Lock()
	defer writerMutex.Unlock()

	if writer, exists := kafkaWriters[topic]; exists {
		return writer
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond, // Buffer messages before sending
		Async:        true,                  // Non-blocking writes
	}

	kafkaWriters[topic] = writer
	return writer
}

// sendToKafka sends the order to the appropriate Kafka topic
func Producer(order *models.Order) error {
	start := time.Now()
	var topic string
	switch order.ProductType {
	case 'O':
		topic = "Options"
	case 'F':
		topic = "Futures"
	default:
		return errors.New("invalid product type")
	}

	writer := getKafkaWriter(topic)

	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// Set a short timeout for Kafka writes
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err = writer.WriteMessages(ctx, kafka.Message{Value: orderBytes})
	if err != nil {
		log.Printf("Kafka write error: %v (took %v)\n", err, time.Since(start))
		return err
	}

	log.Printf("Order pushed to Kafka topic: %s in %v\n", topic, time.Since(start)) // Log execution time
	return nil
}
