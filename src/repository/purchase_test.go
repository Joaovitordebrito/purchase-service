package repository

import (
	"reflect"
	"testing"

	"github.com/Joaovitordebrito/wex-purchase-service/src/model"
	"github.com/google/uuid"
)

func TestMakeHTTPCall(t *testing.T) {
	testTable := []struct {
		name             string
		expectedResponse model.ConvertedPurchase
		purchaseData     model.Purchase
		expectedErr      bool
	}{
		{
			name: "valid payload",
			expectedResponse: model.ConvertedPurchase{
				PurchaseAmount:  1.9,
				TargetCurrency:  5.03,
				ConvertedAmount: 9.56,
			},
			purchaseData: model.Purchase{
				UUID:            uuid.New().String(),
				Description:     "aaaaaaaaaa",
				PurchaseAmount:  1.9,
				TransactionDate: "2023-11-01",
			},
			expectedErr: false,
		},
		{
			name: "invalid date on payload",
			expectedResponse: model.ConvertedPurchase{
				PurchaseAmount:  0,
				TargetCurrency:  0,
				ConvertedAmount: 0,
			},
			purchaseData: model.Purchase{
				UUID:            uuid.New().String(),
				Description:     "aaaaaaaaaa",
				PurchaseAmount:  1.9,
				TransactionDate: "1999-11-01",
			},
			expectedErr: true,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := convertCurrency(tc.purchaseData, "Brazil")
			if !reflect.DeepEqual(resp, tc.expectedResponse) {
				t.Errorf("expected (%v), got (%v)", tc.expectedResponse, resp)
			}

			if err != nil && tc.expectedErr == false {
				t.Errorf("expected (%v), got (%v)", tc.expectedErr, true)
			}
		})
	}
}
