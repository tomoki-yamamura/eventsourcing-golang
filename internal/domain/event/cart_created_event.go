package event

import (
	"time"

	"github.com/google/uuid"
)

type CartCreatedEvent struct {
	aggregateID      uuid.UUID
	aggregateVersion int
	occurredAt       time.Time
}

func NewCartCreatedEvent(aggregateID uuid.UUID, aggregateVersion int) *CartCreatedEvent {
	return &CartCreatedEvent{
		aggregateID:      aggregateID,
		aggregateVersion: aggregateVersion,
		occurredAt:       time.Now(),
	}
}

func (e *CartCreatedEvent) AggregateID() string {
	return e.aggregateID.String()
}

func (e *CartCreatedEvent) AggregateVersion() int {
	return e.aggregateVersion
}

func (e *CartCreatedEvent) OccurredAt() time.Time {
	return e.occurredAt
}
