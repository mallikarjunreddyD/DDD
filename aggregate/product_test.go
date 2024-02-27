package aggregate_test

import (
	"errors"
	"testing"

	"github.com/mallikarjunreddyD/DDD/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		desc        string
		price       float64
		expectedErr error
	}

	testcases := []testCase{
		{
			test:        "Empty product data validation",
			name:        "",
			desc:        "",
			price:       10,
			expectedErr: aggregate.ErrInvalidProductData,
		}, {
			test:        "Valid product data validation",
			name:        "Laptop",
			desc:        "Electornic Item",
			price:       10,
			expectedErr: nil,
		},
		{
			test:        "Valid product data validation",
			name:        "",
			desc:        "Electornic Item",
			price:       10,
			expectedErr: aggregate.ErrInvalidProductData,
		},
		{
			test:        "Valid product data validation",
			name:        "Laptop",
			desc:        "",
			price:       10,
			expectedErr: aggregate.ErrInvalidProductData,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.desc, tc.price)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
