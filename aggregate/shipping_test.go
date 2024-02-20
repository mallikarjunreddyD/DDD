package aggregate_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mallikarjunreddyD/DDD3/aggregate"
	"github.com/mallikarjunreddyD/DDD3/entity"
	valueobjects "github.com/mallikarjunreddyD/DDD3/valueObjects"
)

func TestShipping_NewShipping(t *testing.T) {
	type testCase struct {
		test         string
		mode         string
		deliveryDate time.Time
		product      entity.Item
		address      valueobjects.Address
		expectedErr  error
	}

	testcases := []testCase{
		{
			test:         "valid",
			mode:         "Air",
			deliveryDate: time.Now(),
			product:      entity.Item{},
			address:      valueobjects.Address{},
			expectedErr:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NeWShipping(tc.mode, tc.deliveryDate, &tc.product, tc.address)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
