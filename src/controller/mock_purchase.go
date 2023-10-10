package controller

import (
	"net/http"
)

type MockPurchaseController struct {
	CreatePurchaseFunc       func(w http.ResponseWriter, r *http.Request)
	GetConvertedCurrencyFunc func(w http.ResponseWriter, r *http.Request)
}

func (m *MockPurchaseController) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	if m.CreatePurchaseFunc != nil {
		m.CreatePurchaseFunc(w, r)
	}
}

func (m *MockPurchaseController) GetConvertedCurrency(w http.ResponseWriter, r *http.Request) {
	if m.GetConvertedCurrencyFunc != nil {
		m.GetConvertedCurrencyFunc(w, r)
	}
}
