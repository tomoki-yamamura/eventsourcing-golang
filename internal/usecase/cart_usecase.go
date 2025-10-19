package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/aggregate"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/command"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/domain/repository"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/usecase/input"
	"github.com/tomoki-yamamura/eventsourcing-golang/internal/usecase/output"
)

type CartUsecase interface {
	AddItem(ctx context.Context, input input.AddItemInput) (*output.AddItemResponse, error)
}

type cartUsecaseImpl struct {
	eventStoreRepository repository.EventStoreRepository
	transaction          repository.Transaction
}

func NewCartUsecase(eventStoreRepository repository.EventStoreRepository, transaction repository.Transaction) CartUsecase {
	return &cartUsecaseImpl{
		eventStoreRepository: eventStoreRepository,
		transaction:          transaction,
	}
}

func (u *cartUsecaseImpl) AddItem(ctx context.Context, input input.AddItemInput) (*output.AddItemResponse, error) {
	var response *output.AddItemResponse
	
	err := u.transaction.RWTx(ctx, func(ctx context.Context) error {
		var aggregateID uuid.UUID
		if input.AggregateID == nil {
			aggregateID = uuid.New()
		} else {
			aggregateID = *input.AggregateID
		}

		cmd := command.AddItemCommand{
			AggregateID: aggregateID,
			Description: input.Description,
			Image:       input.Image,
			Price:       input.Price,
			TotalPrice:  input.TotalPrice,
			ItemID:      input.ItemID,
			ProductID:   input.ProductID,
		}

		existingEvents, err := u.eventStoreRepository.LoadEvents(ctx, aggregateID)
		if err != nil {
			return err
		}

		var cart *aggregate.CartAggregate
		if len(existingEvents) == 0 {
			cart = aggregate.NewCartAggregate(aggregateID)
		} else {
			cart = aggregate.NewCartAggregate(aggregateID)
			for _, evt := range existingEvents {
				if err := cart.ApplyEvent(evt); err != nil {
					return err
				}
			}
		}

		newEvents, err := cart.ExecuteAddItemCommand(cmd)
		if err != nil {
			return err
		}

		expectedVersion := cart.Version
		if err := u.eventStoreRepository.SaveEvents(ctx, aggregateID, newEvents, expectedVersion); err != nil {
			return err
		}

		finalVersion := cart.Version + len(newEvents)
		now := time.Now()

		cartItemsReadModel := &output.CartItemsReadModel{
			AggregateID: aggregateID,
			Description: input.Description,
			Image:       input.Image,
			Price:       input.Price,
			TotalPrice:  input.TotalPrice,
			ItemID:      input.ItemID,
			ProductID:   input.ProductID,
			Version:     finalVersion,
			UpdatedAt:   now,
		}

		cartSummaryReadModel := &output.CartSummaryReadModel{
			CartID:      aggregateID,
			ItemCount:   len(cart.CartItems) + 1,
			TotalAmount: input.TotalPrice,
			Version:     finalVersion,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		response = &output.AddItemResponse{
			CommandResult: output.CommandResult{
				Identifier:        aggregateID,
				AggregateSequence: finalVersion,
			},
			CartItems:   cartItemsReadModel,
			CartSummary: cartSummaryReadModel,
		}

		return nil
	})
	
	if err != nil {
		return nil, err
	}

	return response, nil
}