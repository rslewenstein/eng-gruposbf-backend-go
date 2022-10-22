package controllers

import (
	"net/http"

	"go.mod/src/services"
)

// Faz a conversão de todas as moedas cadastradas
func GetConverterAllCurrency(w http.ResponseWriter, r *http.Request) {
	services.GetConverterCurrency()
	w.Write([]byte("Converting..."))
}

// Faz a conversão de uma moeda específica
// func GetConvertedCurrency(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Converting..."))
// }

// Cadastra uma moeda específica
// func CreateCurrency(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Creating..."))
// }
