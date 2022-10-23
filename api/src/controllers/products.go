package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.mod/src/connection"
	"go.mod/src/models"
	"go.mod/src/repository"
	"go.mod/src/responses"
	"go.mod/src/services"
)

// Faz a conversão de uma moeda passada como parâmetro
func GetConvertedCurrency(w http.ResponseWriter, r *http.Request) {
	var c models.Converter
	params := mux.Vars(r)
	c.CurrencyFrom = strings.ToUpper(params["symbol"])
	c.Amount = params["amount"]

	// buscar no banco as siglas
	// c.CurrencyTo, a cada sigla, pesquisar o preço e converter
	db, err := connection.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	symbols, err := repository.GetCurrencies()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	for _, item := range symbols {
		c.CurrencyTo = item.Symbol
		converted, err := services.GetConverterCurrency(c)
		if err != nil {
			log.Fatal(err)
		}

		responses.JSON(w, http.StatusOK, converted)
	}
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
