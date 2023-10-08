package model

import (
	"errors"
	"time"
)

type Purchase struct {
	UUID            string  `json:"UUID,omitempty"`
	Description     string  `json:"description,omitempty"`
	TransactionDate string  `json:"transactionDate,omitempty"`
	PurchaseAmount  float64 `json:"purchaseAmount,omitempty"`
}

type ConvertedPurchase struct {
	PurchaseAmount  float64 `json:"purchaseAmount,omitempty"`
	ExchangeRate    float64 `json:"exchangeRate,omitempty"`
	ConvertedAmount float64 `json:"ConvertedAmount,omitempty"`
}

func (purchase *Purchase) Prepare() error {
	err := purchase.validate()
	if err != nil {
		return err
	}

	err = isValidDate(purchase.TransactionDate)
	if err != nil {
		return err
	}
	return nil
}

func (user *Purchase) validate() error {
	if user.Description == "" {
		return errors.New("field Description is required")
	}
	if user.PurchaseAmount == 0.0 {
		return errors.New("field PurchaseAmount is required")
	}
	if user.TransactionDate == "" {
		return errors.New("field transactionDate is required")
	}
	if user.PurchaseAmount < 0.0 {
		return errors.New("field PurchaseAmount must be positive")
	}

	return nil
}

func isValidDate(inputDateStr string) error {
	_, err := time.Parse("2006-01-02", inputDateStr)
	if err != nil {
		return errors.New("field transactionDate must be a valid date")
	}
	return nil
}
