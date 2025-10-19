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
	Version     int
}

func NewCartAggregate() *CartAggregate {
	return &CartAggregate{
		CartItems: make([]uuid.UUID, 0),
	}
}

func (c *CartAggregate) ExecuteAddItemCommand(cmd command.AddItemCommand) ([]event.Event, error) {
	var events []event.Event
	nextVersion := c.Version + 1

	if c.AggregateID == uuid.Nil {
		cartCreatedEvent := event.NewCartCreatedEvent(cmd.AggregateID, nextVersion)
		events = append(events, cartCreatedEvent)
		nextVersion++
	}

	if err := c.validateCartItemLimit(); err != nil {
		return nil, err
	}

	itemAddedEvent := event.NewItemAddedEvent(
		cmd.AggregateID,
		nextVersion,
		cmd.Description,
		cmd.Image,
		cmd.Price,
		cmd.ItemID,
		cmd.ProductID,
	)
	events = append(events, itemAddedEvent)

	return events, nil
}

func (c *CartAggregate) validateCartItemLimit() error {
	if len(c.CartItems) >= 3 {
		return errors.New("can only add 3 items")
	}
	return nil
}
