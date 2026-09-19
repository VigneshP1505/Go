package streams

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP("localhost:9092"),
			Topic:    "order-created",
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Publish(ctx context.Context, event any) error {
	data, _ := json.Marshal(event)
	return p.writer.WriteMessages(ctx, kafka.Message{
		Value: data,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
