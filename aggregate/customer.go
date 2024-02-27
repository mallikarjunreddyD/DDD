package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD/entity"
	valueobjects "github.com/mallikarjunreddyD/DDD/valueObjects"
)

var ErrInvalidCustomerName = errors.New(" Customer must have an valid name")

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobjects.Transaction
	addresses    []valueobjects.Address
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidCustomerName
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobjects.Transaction, 0),
		addresses:    make([]valueobjects.Address, 0),
	}, nil
}

func (c Customer) GetID() uuid.UUID {
	return c.person.ID
}
func (c Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}
func (c Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
func (c Customer) GetName() string {
	return c.person.Name
}
func (c Customer) SetAge(age int) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Age = age
}
func (c Customer) GetAge() int {
	return c.person.Age
}
