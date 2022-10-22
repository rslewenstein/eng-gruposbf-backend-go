package models

// representa os dados que serão persistidos
type Symbol struct {
	ID      uint64 `json:"id,omitempty"`
	Symbol  string `json:"symbol,omitempty"`
	Country string `json:"country,omitempty"`
}
