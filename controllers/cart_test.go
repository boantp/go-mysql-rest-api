package controllers

import (
	"testing"
)

func TestGetTaxAmount(t *testing.T) {
	getTaxAmountTest := []struct {
		name         string
		calculator   calculator
		hasTaxAmount float64
	}{
		{name: "Food", calculator: food{Amount: 1000}, hasTaxAmount: 100},
		{name: "Tobacco", calculator: tobacco{Amount: 1000}, hasTaxAmount: 50},
		{name: "Entertainment", calculator: entertainment{Amount: 150}, hasTaxAmount: 0.5},
	}

	for _, tt := range getTaxAmountTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.calculator.getTaxAmount()
			if got != tt.hasTaxAmount {
				t.Errorf("%#v got %.2f want %.2f", tt.calculator, got, tt.hasTaxAmount)
			}
		})
	}
}
