package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type OutboxEvent struct {
	ID          uuid.UUID       `json:"id"`
	EventType   string          `json:"event_type"`
	AggregateId uuid.UUID       `json:"aggregate_id"`
	Payload     json.RawMessage `json:"payload"`
	Published   bool            `json:"published"`
	CreatedAt   time.Time       `json:"created_at"`
	PublishedAt time.Time       `json:"published_at"`
}
