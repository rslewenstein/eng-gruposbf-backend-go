package controllers

import "net/http"

// Faz a conversão de todas as moedas cadastradas
func GetConverterAllCurrency(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Converting..."))
}

// Faz a conversão de uma moeda específica
// func GetConverterCurrency(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Converting..."))
// }

// Cadastra uma moeda específica
// func CreateCurrency(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Creating..."))
// }
