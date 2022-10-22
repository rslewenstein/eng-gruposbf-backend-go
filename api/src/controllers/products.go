package controllers

import (
	"net/http"
	"strings"

	"go.mod/src/services"
)

// Faz a conversão de uma moeda específica
func GetConvertedCurrency(w http.ResponseWriter, r *http.Request) {
	productId := strings.ToLower(r.URL.Query().Get("productid"))
	symbol := strings.ToLower(r.URL.Query().Get("symbol"))
	amount := strings.ToLower(r.URL.Query().Get("symbol"))
	// buscar o produto
	services.GetConverterCurrency()
	w.Write([]byte("ProdutoId: " + productId + ", symbol: " + symbol + ", amount: " + amount))
}

// Cadastra uma moeda específica
// func CreateCurrency(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Creating..."))
// }
