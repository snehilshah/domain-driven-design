package aggregate

import (
	"ddd/entity"
	"ddd/valueobject"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("A customer has to have a name")
)

// aggregates combines many entities into a full object
type Customer struct {
	// Person will be the root entity of the customer
	// which means person.ID is the main identifier for the customer
	person   *entity.Person
	products []*entity.Item
	// transactions are not pointers because it is not going to change hence, no point in making it as a pointer
	transactions []*valueobject.Transaction
}

// NewCustomer is a factory function for creating a new Customer
// It will also do validation like:
// name is not empty etc.
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{Name: name, ID: uuid.New()}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}
