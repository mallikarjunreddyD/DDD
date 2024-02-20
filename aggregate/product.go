package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/entity"
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
			ID:   uuid.UUID{},
			Name: name,
			Desc: desc,
		},
		price:    price,
		quantity: 0,
	}, nil
}
