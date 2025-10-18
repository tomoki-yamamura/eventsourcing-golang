package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/command"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/event"
)

type CartAggregate struct {
	AggregateID uuid.UUID
	CartItems   []uuid.UUID
}

func NewCartAggregate() *CartAggregate {
	return &CartAggregate{
		CartItems: make([]uuid.UUID, 0),
	}
}

func (c *CartAggregate) ApplyAddItemEvent(cmd command.AddItemCommand) ([]event.Event, error) {
	var events []event.Event

	// If this is a new aggregate, create it first
	if c.AggregateID == uuid.Nil {
		cartCreatedEvent := &event.CartCreatedEvent{
			AggregateID: cmd.AggregateID,
		}
		events = append(events, cartCreatedEvent)
	}

	// Business rule: maximum 3 items per cart
	if len(c.CartItems) >= 3 {
		return nil, errors.New("can only add 3 items")
	}

	itemAddedEvent := &event.ItemAddedEvent{
		AggregateID: cmd.AggregateID,
		Description: cmd.Description,
		Image:       cmd.Image,
		Price:       cmd.Price,
		ItemID:      cmd.ItemID,
		ProductID:   cmd.ProductID,
	}
	events = append(events, itemAddedEvent)

	return events, nil
}
