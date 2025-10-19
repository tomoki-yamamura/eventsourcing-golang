package event

import (
	"time"

	"github.com/google/uuid"
)

type ItemAddedEvent struct {
	aggregateID      uuid.UUID
	aggregateVersion int
	occurredAt       time.Time
	Description      string
	Image            string
	Price            float64
	ItemID           uuid.UUID
	ProductID        uuid.UUID
}

func NewItemAddedEvent(aggregateID uuid.UUID, aggregateVersion int, description, image string, price float64, itemID, productID uuid.UUID) *ItemAddedEvent {
	return &ItemAddedEvent{
		aggregateID:      aggregateID,
		aggregateVersion: aggregateVersion,
		occurredAt:       time.Now(),
		Description:      description,
		Image:            image,
		Price:            price,
		ItemID:           itemID,
		ProductID:        productID,
	}
}

func (e *ItemAddedEvent) AggregateID() string {
	return e.aggregateID.String()
}

func (e *ItemAddedEvent) AggregateVersion() int {
	return e.aggregateVersion
}

func (e *ItemAddedEvent) OccurredAt() time.Time {
	return e.occurredAt
}
