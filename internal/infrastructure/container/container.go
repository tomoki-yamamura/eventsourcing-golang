package container

import (
	"context"

	"github.com/tomoki-yamamura/eventsourcing-golang/internal/config"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/infrastructure/eventstore"
)

type Container struct {
	Cfg *config.Config

	EventStoreClient *eventstore.Client
}

func NewContainer(ctx context.Context, cfg *config.Config) (*Container, error) {
	eventStoreClient, err := eventstore.NewClient(cfg.DatabaseConfig)
	if err != nil {
		return nil, err
	}

	return &Container{
		Cfg:              cfg,
		EventStoreClient: eventStoreClient,
	}, nil
}
