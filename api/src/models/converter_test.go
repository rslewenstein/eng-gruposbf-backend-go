package models_test

import "testing"

type Converter struct {
	CurrencyFrom string `json:"currencyFrom,omitempty"`
	CurrencyTo   string `json:"currencyTo,omitempty"`
	Amount       string `json:"amount,omitempty"`
}

func TestConverterModel(t *testing.T) {
	ret := MockConverter()

	if ret.CurrencyFrom == "" {
		t.Failed()
	}

	if ret.CurrencyTo == "" {
		t.Failed()
	}

	if ret.Amount == "" {
		t.Failed()
	}

	if ret.Amount == "0.00" {
		t.Failed()
	}
}

func MockConverter() Converter {
	var converter Converter
	converter.CurrencyFrom = "BRL"
	converter.CurrencyTo = "USD"
	converter.Amount = "530.00"

	return converter
}
