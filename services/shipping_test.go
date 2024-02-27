package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
	valueobjects "github.com/mallikarjunreddyD/DDD3/valueObjects"
)

func init_shippings(t *testing.T) []aggregate.Shipping {
	products := init_products(t)

	shippings := make([]aggregate.Shipping, 0)
	for _, product := range products {
		shipping, err := aggregate.NeWShipping("Air", time.Date(2023, time.March, 22, 12, 12, 12, 12, time.UTC), product.GetItem(), valueobjects.Address{})

		if err != nil {
			return shippings
		}
		shippings = append(shippings, shipping)
	}
	return shippings
}

func TestOrder_NewShippingService(t *testing.T) {
	shippings := init_shippings(t)
	products := init_products(t)
	ss, err := NewShippingService(
		withMemoryShippingRepository(),
	)

	if err != nil {
		t.Fatal(err)
	}
	for _, ship := range shippings {

		err = ss.shippings.Add(ship)
		if err != nil {
			t.Fatal(err)
		}
	}
	shippingIDs := []uuid.UUID{
		shippings[0].GetID(),
		shippings[1].GetID(),
		shippings[2].GetID(),
	}
	productsIDs := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
		products[2].GetID(),
	}
	_, err = ss.CreateShipping(productsIDs, shippingIDs)
	if err != nil {
		t.Fatal(err)
	}

}
