package output

import (
	"time"

	"github.com/google/uuid"
)

type AddItemResponse struct {
	CommandResult CommandResult         `json:"commandResult"`
	CartItems     *CartItemsReadModel   `json:"cartItems"`
	CartSummary   *CartSummaryReadModel `json:"cartSummary"`
}

type CommandResult struct {
	Identifier        uuid.UUID `json:"identifier"`
	AggregateSequence int       `json:"aggregateSequence"`
}

type CartItemsReadModel struct {
	AggregateID uuid.UUID `json:"aggregateId"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	TotalPrice  float64   `json:"totalPrice"`
	ItemID      uuid.UUID `json:"itemId"`
	ProductID   uuid.UUID `json:"productId"`
	Version     int       `json:"version"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CartSummaryReadModel struct {
	CartID      uuid.UUID `json:"cartId"`
	ItemCount   int       `json:"itemCount"`
	TotalAmount float64   `json:"totalAmount"`
	Version     int       `json:"version"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
