package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Joaovitordebrito/wex-purchase-service/src/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestValidadeUser(t *testing.T) {
	testTable := []struct {
		name         string
		purchaseData model.Purchase
		expectedErr  bool
	}{
		{
			name: "valid payload",

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
			purchaseData: model.Purchase{
				UUID:            uuid.New().String(),
				Description:     "aaaaaaaaaa",
				PurchaseAmount:  1.9,
				TransactionDate: "1999-116-01",
			},
			expectedErr: true,
		},
		{
			name: "invalid description",
			purchaseData: model.Purchase{
				UUID:            uuid.New().String(),
				Description:     "",
				PurchaseAmount:  1.9,
				TransactionDate: "2023-11-01",
			},
			expectedErr: true,
		},
		{
			name: "invalid purchase amount",
			purchaseData: model.Purchase{
				UUID:            uuid.New().String(),
				Description:     "aasasdaaa",
				PurchaseAmount:  -1.9,
				TransactionDate: "2023-11-01",
			},
			expectedErr: true,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.purchaseData.Prepare()

			if err != nil && tc.expectedErr == false {
				t.Errorf("expected (%v), got (%v)", tc.expectedErr, true)
			}
		})
	}
}
func TestCreatePurchase(t *testing.T) {
	mockController := &MockPurchaseController{
		CreatePurchaseFunc: func(w http.ResponseWriter, r *http.Request) {

			purchase := model.Purchase{
				Description:     "Test",
				TransactionDate: "2023-10-09",
				PurchaseAmount:  100.0,
			}
			json.NewEncoder(w).Encode(purchase)
		},
	}

	reqBody, _ := json.Marshal(model.Purchase{
		Description:     "Test",
		TransactionDate: "2023-10-09",
		PurchaseAmount:  100.0,
	})
	req := httptest.NewRequest("POST", "/purchase", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	mockController.CreatePurchase(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestGetConvertedCurrency(t *testing.T) {
	mockController := &MockPurchaseController{
		GetConvertedCurrencyFunc: func(w http.ResponseWriter, r *http.Request) {
			convertedPurchase := model.ConvertedPurchase{
				PurchaseAmount:  100.0,
				TargetCurrency:  5.0,
				ConvertedAmount: 500.0,
			}
			json.NewEncoder(w).Encode(convertedPurchase)
		},
	}

	req := httptest.NewRequest("GET", "/converted/currency/some-uuid/some-country", nil)
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/converted/currency/{uuid}/{country}", mockController.GetConvertedCurrency)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
