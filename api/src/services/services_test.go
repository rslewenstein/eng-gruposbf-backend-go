package services_test

import (
	"testing"

	"go.mod/src/models"
	"go.mod/src/services"
)

func TestGetConverterCurrency(t *testing.T) {
	var converter models.Converter
	converter.CurrencyFrom = "BRL"
	converter.CurrencyTo = "USD"
	converter.Amount = "500"

	resp, err := services.GetConverterCurrency(converter)

	if err != nil {
		t.Errorf("Fail")
	}
	if resp.Price == "0.00" {
		t.Errorf("Fail")
	}
}
