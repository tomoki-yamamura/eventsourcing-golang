package aggregate

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/command"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/event"
)

type CartAggregate struct {
	AggregateID uuid.UUID
	CartItems   []uuid.UUID
	Version     int
}

func NewCartAggregate(aggregateID uuid.UUID) *CartAggregate {
	return &CartAggregate{
		AggregateID: aggregateID,
		CartItems:   make([]uuid.UUID, 0),
		Version:     0,
	}
}

func (c *CartAggregate) ExecuteAddItemCommand(cmd command.AddItemCommand) ([]event.Event, error) {
	var eventList []event.Event
	nextVersion := c.Version + 1

	// If this is a new cart (Version == 0), create CartCreatedEvent first
	if c.Version == 0 {
		cartCreatedEvent := event.NewCartCreatedEvent(c.AggregateID, nextVersion)
		eventList = append(eventList, cartCreatedEvent)
		nextVersion++
	}

	if err := c.validateCartItemLimit(); err != nil {
		return nil, err
	}

	itemAddedEvent := event.NewItemAddedEvent(
		c.AggregateID,
		nextVersion,
		cmd.Description,
		cmd.Image,
		cmd.Price,
		cmd.ItemID,
		cmd.ProductID,
	)
	eventList = append(eventList, itemAddedEvent)

	return eventList, nil
}

func (c *CartAggregate) validateCartItemLimit() error {
	if len(c.CartItems) >= 3 {
		return errors.New("can only add 3 items")
	}
	return nil
}

// ApplyEvent applies event to the aggregate for hydration
func (c *CartAggregate) ApplyEvent(evt event.Event) error {
	switch e := evt.(type) {
	case *event.CartCreatedEvent:
		c.ApplyCartCreatedEvent(e)
	case *event.ItemAddedEvent:
		c.ApplyItemAddedEvent(e)
	default:
		return fmt.Errorf("unknown event type: %T", evt)
	}
	return nil
}

func (c *CartAggregate) ApplyCartCreatedEvent(event *event.CartCreatedEvent) {
	aggregateID, _ := uuid.Parse(event.AggregateID())
	c.AggregateID = aggregateID
	c.Version = event.AggregateVersion()
}

func (c *CartAggregate) ApplyItemAddedEvent(event *event.ItemAddedEvent) {
	c.CartItems = append(c.CartItems, event.ItemID)
	c.Version = event.AggregateVersion()
}
