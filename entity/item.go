package entity

import "github.com/google/uuid"

// package entity holds all entries that are shared across all sub-domains

// Item is an entity that represents an item in all domains
type Item struct {
	// ID is the unique identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
