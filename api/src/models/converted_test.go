package models_test

import "testing"

type Converted struct {
	Symbol string `json:"symbol,omitempty"`
	Price  string `json:"price,omitempty"`
}

func TestConvertedModel(t *testing.T) {
	ret := MockConverted()

	if ret.Symbol == "" {
		t.Failed()
	}

	if ret.Price == "" {
		t.Failed()
	}

	if ret.Price == "0.00" {
		t.Failed()
	}
}

func MockConverted() Converted {
	var converted Converted
	converted.Symbol = "USD"
	converted.Price = "530.00"

	return converted
}
