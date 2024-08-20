package entity

import "github.com/google/uuid"

// package entity holds all entries that are shared across all sub-domains

// Person is an entity that represents a person in all domains
type Person struct {
	// ID is the unique identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
