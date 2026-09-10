package main

import (
	"context"
	"log"
	"time"

	"github.com/vignesh/online-ordering/internal/config"
	"github.com/vignesh/online-ordering/internal/db"
	"github.com/vignesh/online-ordering/internal/repository"
	"github.com/vignesh/online-ordering/internal/streams"
)

func main() {

	ctx := context.Background()
	cfg := config.Load()
	dbPool, _ := db.New(cfg)
	repo := repository.NewOrderRepository(dbPool)

	producer := streams.NewProducer()

	defer producer.Close()

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		tx, events, err := repo.LockUnpublishedEvents(ctx, 100)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, event := range events {
			if err = producer.Publish(ctx, event); err != nil {
				log.Println("kafka publish failed", err)
				continue
			}
			repo.MarkPublished(ctx, event.OrderId)
		}

	}

}
