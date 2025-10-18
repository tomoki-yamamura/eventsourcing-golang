package event

import (
	"github.com/google/uuid"
)

type CartCreatedEvent struct {
	AggregateID uuid.UUID
}
