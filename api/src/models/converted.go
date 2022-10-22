package models

// representa os dados que serão retornados
type Converted struct {
	Symbol string `json:"symbol,omitempty"`
	Price  string `json:"price,omitempty"`
}
