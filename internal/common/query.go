package common

// Query represents the base interface for all queries in the system
type Query interface{}

// QueryHandler handles queries and returns results
type QueryHandler[T Query, U any] interface {
	HandleQuery(query T) U
}