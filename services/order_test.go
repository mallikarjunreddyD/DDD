package services

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	airConditioner, err := aggregate.NewProduct("Air Conditioner", "Samsung, 110, 5 star, long last cooling", 30100)
	if err != nil {
		log.Fatal(err)
	}

	mobile, err := aggregate.NewProduct("Mobile", "Samsung, 8 GB RAM, 64GB ROM", 15999)
	if err != nil {
		log.Fatal(err)
	}

	laptop, err := aggregate.NewProduct("Laptop", "Inter i5, 16 GB, 256 GB SSD", 45000.99)
	if err != nil {
		log.Fatal(err)
	}

	return []aggregate.Product{airConditioner, mobile, laptop}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		withMemoryCustomerRepository(),
		withMemoryProductRepository(),
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
	productIDs := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
		products[2].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), productIDs)
	if err != nil {
		t.Fatal(err)
	}

}
