package event

import (
	"github.com/google/uuid"
)

type ItemAddedEvent struct {
	AggregateID uuid.UUID
	Description string
	Image       string
	Price       float64
	ItemID      uuid.UUID
	ProductID   uuid.UUID
}
