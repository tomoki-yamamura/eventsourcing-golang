package command

import (
	"github.com/google/uuid"
)

type AddItemCommand struct {
	AggregateID uuid.UUID
	Description string
	Image       string
	Price       float64
	TotalPrice  float64
	ItemID      uuid.UUID
	ProductID   uuid.UUID
}

func (c AddItemCommand) GetAggregateID() uuid.UUID {
	return c.AggregateID
}
