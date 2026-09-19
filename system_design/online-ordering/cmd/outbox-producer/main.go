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
		tx, _events, err := repo.LockUnpublishedEvents(ctx, 100)
		if err != nil {
			log.Println(err)
			continue
		}

		if len(_events) == 0 {
			tx.Rollback(ctx)
			continue
		}

		success := true

		for _, _event := range _events {

			if err = producer.Publish(ctx, _event); err != nil {
				log.Println("kafka publish failed", err)
				success = false
				continue
			}
			if err = repo.MarkPublished(ctx, _event.ID); err != nil {
				log.Println("marking published events published failed")
				success = false
				continue
			}

			if success == false {
				tx.Rollback(ctx)
			} else {
				if err = tx.Commit(ctx); err != nil {
					log.Println("failed to commit")
				}
			}

		}

	}

}
