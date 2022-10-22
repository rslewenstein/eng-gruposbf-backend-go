package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port                   = 0
	Api_key                = ""
	Symbol_country_default = ""
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	Api_key = os.Getenv("API_KEY")

	Symbol_country_default = os.Getenv("SYMBOL_COUNTRY_DEFAULT")
}
