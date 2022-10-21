package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mod/src/config"
)

func GetConverterCurrency() {
	const currencyFrom string = "BRL"
	var currencyTo string = "USD"
	var amount string = "529.00"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}