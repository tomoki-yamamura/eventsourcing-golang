package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/event"
)

type EventStoreRepository interface {
	LoadEvents(ctx context.Context, aggregateID uuid.UUID) ([]event.Event, error)
	SaveEvents(ctx context.Context, aggregateID uuid.UUID, events []event.Event, expectedVersion int) error
}