package event

import "time"

type Event interface {
	AggregateID() string
	AggregateVersion() int
	OccurredAt() time.Time
}
