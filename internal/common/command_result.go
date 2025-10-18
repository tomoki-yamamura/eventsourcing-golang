package common

import "github.com/google/uuid"

// CommandResult represents the result of a command execution
// that allows giving feedback to the client to update
type CommandResult struct {
	Identifier        uuid.UUID `json:"identifier"`
	AggregateSequence int64     `json:"aggregateSequence"`
}