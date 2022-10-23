package models_test

import "testing"

type Symbol struct {
	ID      uint64 `json:"id,omitempty"`
	Symbol  string `json:"symbol,omitempty"`
	Country string `json:"country,omitempty"`
}

func TestSymbolModel(t *testing.T) {
	ret := MockSymbol()

	if ret.ID < 1 {
		t.Failed()
	}

	if ret.Symbol == "" {
		t.Failed()
	}

	if ret.Country == "" {
		t.Failed()
	}
}

func MockSymbol() Symbol {
	var symbol Symbol
	symbol.ID = 1000
	symbol.Symbol = "USD"
	symbol.Country = "Estados Unidos"

	return symbol
}
