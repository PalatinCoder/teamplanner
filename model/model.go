package model

// Model is the common interface for every type in the domain model
type Model interface {
	// Key provides the database friendly key
	Key() string
	// ID provides a short identifier for the entity
	ID() string
}
