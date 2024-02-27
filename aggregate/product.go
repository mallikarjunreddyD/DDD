package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD/entity"
)

var ErrInvalidProductData = errors.New("product must have valid name or description")

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, desc string, price float64) (Product, error) {
	if name == "" || desc == "" {
		return Product{}, ErrInvalidProductData
	}

	return Product{
		item: &entity.Item{
			ID:   uuid.New(),
			Name: name,
			Desc: desc,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}
func (p Product) SetID(id uuid.UUID) {
	if p.item == nil {
		p.item = &entity.Item{}
	}
	p.item.ID = id
}
func (p Product) GetName() string {
	return p.item.Name
}
func (p Product) SetName(name string) {
	if p.item == nil {
		p.item = &entity.Item{}
	}
	p.item.Name = name
}
func (p Product) GetDesc() string {
	return p.item.Desc
}
func (p Product) SetDesc(desc string) {
	if p.item == nil {
		p.item = &entity.Item{}
	}
	p.item.Desc = desc
}
func (p Product) GetItem() *entity.Item {
	return p.item
}
func (p Product) GetPrice() float64 {
	return p.price
}
