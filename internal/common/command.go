package common

import "github.com/google/uuid"

// Command represents the base interface for all commands in the system
type Command interface {
	GetAggregateID() uuid.UUID
}