package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mod/src/config"
	"go.mod/src/models"
)

type Response struct {
	ResultRes float32  `json:"result"`
	QueryRes  QueryRes `json:"query"`
}

type QueryRes struct {
	SymbolRes string `json:"to"`
}

func GetConverterCurrency(c models.Converter) ([]models.Converted, error) {
	var currencyFrom string = c.CurrencyFrom //"BRL"
	var currencyTo string = c.CurrencyTo     //"USD"
	var amount string = c.Amount             //"529.00"
	url := "https://api.apilayer.com/exchangerates_data/convert?to=" + currencyTo + "&from=" + currencyFrom + "&amount=" + amount + ""

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", config.Api_key)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(body, &responseObject)
	fmt.Println(responseObject.QueryRes.SymbolRes)
	fmt.Println(responseObject.ResultRes)

	//Convert the body to type string
	// sb := string(body)
	// fmt.Println(sb)

	//fmt.Println(string(body))

	return nil, nil
}
