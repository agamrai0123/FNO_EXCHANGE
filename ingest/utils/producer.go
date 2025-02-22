package utils

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/agamrai0123/FNO_EXCHANGE/ingest/models"
	"github.com/segmentio/kafka-go"
)

func Producer(order *models.Order) error {
	var topic string
	if order.ProductType == 'O' {
		topic = "Options"
	} else {
		return errors.New("topic not specified")
	}
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	jsonData, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: jsonData,
		},
	)
	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
		return err
	}

	log.Println("Message sent successfully!")
	return nil
}
