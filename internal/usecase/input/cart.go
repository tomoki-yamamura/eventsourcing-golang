package input

import "github.com/google/uuid"

type AddItemInput struct {
	AggregateID *uuid.UUID `json:"cartId,omitempty"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Price       float64    `json:"price"`
	TotalPrice  float64    `json:"totalPrice"`
	ItemID      uuid.UUID  `json:"itemId"`
	ProductID   uuid.UUID  `json:"productId"`
}