package models

// representa os dados que serão convertidos
type Converter struct {
	CurrencyFrom string `json:"currencyFrom,omitempty"`
	CurrencyTo   string `json:"currencyTo,omitempty"`
	Amount       string `json:"amount,omitempty"`
}
