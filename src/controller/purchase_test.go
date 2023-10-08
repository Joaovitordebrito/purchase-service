package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCreatePurchase(t *testing.T) {
	requestBody := []byte(`{
		"description":"asasd",
		"purchaseAmount":0.9
	}`)

	req, err := http.NewRequest(http.MethodPost, "/purchase", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	CreatePurchase(w, req)

	res := w.Result()

	if !reflect.DeepEqual(http.StatusCreated, res.StatusCode) {
		t.Errorf("wanted: %d, got: %d", http.StatusCreated, res.StatusCode)
	}
}
