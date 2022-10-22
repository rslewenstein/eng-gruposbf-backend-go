package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"go.mod/src/connection"
	"go.mod/src/models"
	"go.mod/src/repository"
	"go.mod/src/responses"
	"go.mod/src/services"
)

// Faz a conversão de uma moeda passada como parâmetro
func GetConvertedCurrency(w http.ResponseWriter, r *http.Request) {
	var c models.Converter
	c.CurrencyFrom = strings.ToLower(r.URL.Query().Get("symbol"))
	c.Amount = strings.ToLower(r.URL.Query().Get("amount"))

	// buscar no banco as siglas
	// c.CurrencyTo, a cada sigla, pesquisar o preço e converter

	teste, err := services.GetConverterCurrency(c)
	if err != nil {
		log.Fatal(err)
	}

	responses.JSON(w, http.StatusOK, teste)
}

// Cadastra uma moeda específica (sigla e país)
func CreateCurrency(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var symbol models.Symbol
	if err = json.Unmarshal(bodyReq, &symbol); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := connection.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	symbol.ID, err = repository.CreateSymbol(symbol)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, symbol)
}
