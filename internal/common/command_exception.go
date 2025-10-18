package common

import "fmt"

// CommandException represents errors that occur during command processing
type CommandException struct {
	Message string
	Cause   error
}

func (e *CommandException) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func NewCommandException(message string, cause error) *CommandException {
	return &CommandException{
		Message: message,
		Cause:   cause,
	}
}