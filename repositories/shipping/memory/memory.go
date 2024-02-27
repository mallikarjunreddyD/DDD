package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
	"github.com/mallikarjunreddyD/DDD3/repositories/shipping"
)

type MemoryShippingRepository struct {
	shippings map[uuid.UUID]aggregate.Shipping
	sync.Mutex
}

func New() *MemoryShippingRepository {
	return &MemoryShippingRepository{
		shippings: make(map[uuid.UUID]aggregate.Shipping),
	}
}

func (msr *MemoryShippingRepository) GetAll() ([]aggregate.Shipping, error) {
	var shippings []aggregate.Shipping
	for _, shipping := range msr.shippings {
		shippings = append(shippings, shipping)
	}
	return shippings, nil
}
func (msr *MemoryShippingRepository) Get(id uuid.UUID) (aggregate.Shipping, error) {
	if shipping, ok := msr.shippings[id]; ok {
		return shipping, nil
	}
	return aggregate.Shipping{}, shipping.ErrShippingNotFound
}
func (msr *MemoryShippingRepository) Add(p aggregate.Shipping) error {
	if msr.shippings == nil {
		msr.shippings = make(map[uuid.UUID]aggregate.Shipping)
	}
	if _, ok := msr.shippings[p.GetID()]; ok {
		return shipping.ErrShippingAlreadyExists
	}
	msr.Lock()
	msr.shippings[p.GetID()] = p
	msr.Unlock()
	return nil
}
func (msr *MemoryShippingRepository) Update(p aggregate.Shipping) error {
	if _, ok := msr.shippings[p.GetID()]; !ok {
		return shipping.ErrShippingDoesNotExists
	}
	msr.Lock()
	msr.shippings[p.GetID()] = p
	msr.Unlock()
	return nil
}
func (msr *MemoryShippingRepository) Delete(id uuid.UUID) error {
	msr.Lock()
	defer msr.Unlock()
	if _, ok := msr.shippings[id]; !ok {
		return shipping.ErrShippingDoesNotExists
	}
	delete(msr.shippings, id)
	return nil
}
