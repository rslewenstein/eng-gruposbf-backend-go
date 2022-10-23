package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mod/src/config"
	"go.mod/src/router"
	"go.mod/src/services"
)

func main() {
	config.LoadEnv()

	r := router.GenerateRouters()
	services.LoggerSystem("sucesso", "running...", "sucesso")
	fmt.Printf("Listening on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
