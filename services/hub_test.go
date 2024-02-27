package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
)

func Test_Hub(t *testing.T) {
	products := init_products(t)
	shippings := init_shippings(t)

	os, err := NewOrderService(
		withMemoryCustomerRepository(),
		withMemoryProductRepository(),
	)
	if err != nil {
		t.Fatal(err)
	}

	ss, err := NewShippingService(
		withMemoryShippingRepository(),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Arjun")
	if err != nil {
		t.Fatal(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range products {
		err = os.products.Add(p)
		if err != nil {
			t.Fatal(err)
		}
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
	hs, err := NewHub(
		WithOrderService(os),
		WithShippingService(ss),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = hs.OrderAndShip(cust.GetID(), productsIDs, shippingIDs)
	if err != nil {
		t.Fatal(err)
	}
}
