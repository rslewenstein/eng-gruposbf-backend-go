package models

// representa os dados que serão convertidos
type Converter struct {
	CurrencyTo string `json:"currencyTo,omitempty"`
	Amount     string `json:"amount,omitempty"`
}
