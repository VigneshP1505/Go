package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/vignesh/online-ordering/internal/events"
)

func Consumer() {
	fmt.Println("Consumer running")
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			"localhost:9092",
		},
		Topic:   "order-created",
		GroupID: "payment-workers",
	})
	defer reader.Close()
	log.Printf("Payment worker started...")
	log.Printf("Waiting for orders...")

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("kafka error:", err)
			continue
		}

		log.Printf("message received: partition %d and offset %d", message.Partition, message.Offset)

		var event events.OrderCreatedEvent
		err = json.Unmarshal(message.Value, &event)

		if err != nil {
			log.Printf("failed to deserialize:", err)
			continue
		}

		time.Sleep(2 * time.Second)
		log.Printf("order finished processing")

	}
}
