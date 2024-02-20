package aggregate_test

import (
	"errors"
	"testing"

	"github.com/mallikarjunreddyD/DDD3/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testcases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidCustomerName,
		}, {
			test:        "Valid name",
			name:        "Mallikarjun",
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
