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

	db, err := connection.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		services.LoggerSystem("erro", "GetConvertedCurrency, connection.Connect()", string(err.Error()))
		return
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	symbols, err := repository.GetCurrencies()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		services.LoggerSystem("erro", "GetConvertedCurrency, repository.GetCurrencies()", string(err.Error()))
		return
	}

	for _, item := range symbols {
		c.CurrencyTo = item.Symbol
		converted, err := services.GetConverterCurrency(c)
		if err != nil {
			log.Fatal(err)
			services.LoggerSystem("erro", "GetConvertedCurrency, services.GetConverterCurrency(c)", string(err.Error()))
		}

		responses.JSON(w, http.StatusOK, converted)
		services.LoggerSystem("sucesso", "GetConvertedCurrency", "Conversão realizada")
	}
}

// Cadastra uma moeda específica (sigla e país)
func CreateCurrency(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		services.LoggerSystem("erro", "CreateCurrency, io.ReadAll(r.Body)", string(err.Error()))
		return
	}

	var symbol models.Symbol
	if err = json.Unmarshal(bodyReq, &symbol); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		services.LoggerSystem("erro", "CreateCurrency, json.Unmarshal", string(err.Error()))
		return
	}

	db, err := connection.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		services.LoggerSystem("erro", "GetConvertedCurrency, connection.Connect()", string(err.Error()))
		return
	}
	defer db.Close()

	repository := repository.NewRepository(db)

	symbol.ID, err = repository.CreateSymbol(symbol)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		services.LoggerSystem("erro", "CreateCurrency, repository.CreateSymbol(symbol)", string(err.Error()))
		return
	}

	responses.JSON(w, http.StatusCreated, symbol)
	services.LoggerSystem("sucesso", "CreateCurrency", "Moeda salva no BD")
}
