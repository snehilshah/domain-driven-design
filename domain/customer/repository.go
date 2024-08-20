package customer

import (
	"ddd/aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed to update customer")
)

type CustomerRepository interface {
	// It does not matter if the GET, ADD, UPDATE is done from mongo, mysql, postgre
	GET(uuid.UUID) (aggregate.Customer, error)
	ADD(aggregate.Customer) (aggregate.Customer, error)
	UPDATE(aggregate.Customer) error
}
