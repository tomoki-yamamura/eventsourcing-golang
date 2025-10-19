package eventstore

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/event"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/repository"
)

type Repository struct{}

func NewRepository() repository.EventStoreRepository {
	return &Repository{}
}

func (r *Repository) LoadEvents(ctx context.Context, aggregateID uuid.UUID) ([]event.Event, error) {
	tx, err := GetTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	query := `
		SELECT event_type, event_data, version 
		FROM event_store 
		WHERE aggregate_id = ? 
		ORDER BY version ASC
	`

	rows, err := tx.QueryContext(ctx, query, aggregateID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to load events: %w", err)
	}
	defer rows.Close()

	var eventList []event.Event
	for rows.Next() {
		var eventType string
		var eventData []byte
		var version int

		if err := rows.Scan(&eventType, &eventData, &version); err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		event, err := r.deserializeEvent(eventType, eventData, aggregateID, version)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize event: %w", err)
		}

		eventList = append(eventList, event)
	}

	return eventList, nil
}

func (r *Repository) SaveEvents(ctx context.Context, aggregateID uuid.UUID, eventList []event.Event, expectedVersion int) error {
	tx, err := GetTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to get transaction: %w", err)
	}

	// Check current version for optimistic locking
	var currentVersion int
	checkQuery := `
		SELECT COALESCE(MAX(version), 0) 
		FROM event_store 
		WHERE aggregate_id = ?
	`
	if err := tx.GetContext(ctx, &currentVersion, checkQuery, aggregateID.String()); err != nil {
		return fmt.Errorf("failed to check current version: %w", err)
	}

	if currentVersion != expectedVersion {
		return fmt.Errorf("concurrency conflict: expected version %d, but current version is %d", expectedVersion, currentVersion)
	}

	// Insert new events
	insertQuery := `
		INSERT INTO event_store (aggregate_id, aggregate_type, event_type, event_data, version, created_at)
		VALUES (?, ?, ?, ?, ?, NOW())
	`

	for _, event := range eventList {
		eventType, eventData, version, err := r.serializeEvent(event)
		if err != nil {
			return fmt.Errorf("failed to serialize event: %w", err)
		}

		if _, err := tx.ExecContext(ctx, insertQuery, aggregateID.String(), "Cart", eventType, eventData, version); err != nil {
			return fmt.Errorf("failed to insert event: %w", err)
		}
	}

	return nil
}

func (r *Repository) serializeEvent(e event.Event) (string, []byte, int, error) {
	switch evt := e.(type) {
	case *event.CartCreatedEvent:
		data, err := json.Marshal(map[string]interface{}{
			"aggregateID": evt.AggregateID(),
		})
		return "CartCreatedEvent", data, evt.AggregateVersion(), err

	case *event.ItemAddedEvent:
		data, err := json.Marshal(map[string]interface{}{
			"aggregateID": evt.AggregateID(),
			"description": evt.Description,
			"image":       evt.Image,
			"price":       evt.Price,
			"itemID":      evt.ItemID.String(),
			"productID":   evt.ProductID.String(),
		})
		return "ItemAddedEvent", data, evt.AggregateVersion(), err

	default:
		return "", nil, 0, fmt.Errorf("unknown event type: %T", e)
	}
}

func (r *Repository) deserializeEvent(eventType string, eventData []byte, aggregateID uuid.UUID, version int) (event.Event, error) {
	switch eventType {
	case "CartCreatedEvent":
		return event.NewCartCreatedEvent(aggregateID, version), nil

	case "ItemAddedEvent":
		var data map[string]interface{}
		if err := json.Unmarshal(eventData, &data); err != nil {
			return nil, err
		}

		itemID, _ := uuid.Parse(data["itemID"].(string))
		productID, _ := uuid.Parse(data["productID"].(string))

		return event.NewItemAddedEvent(
			aggregateID,
			version,
			data["description"].(string),
			data["image"].(string),
			data["price"].(float64),
			itemID,
			productID,
		), nil

	default:
		return nil, fmt.Errorf("unknown event type: %s", eventType)
	}
}