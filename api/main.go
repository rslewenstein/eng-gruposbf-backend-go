package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mod/src/config"
)

func main() {
	config.LoadEnv()

	fmt.Printf("Listening on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))

	// myRoute := mux.NewRouter().StrictSlash(true)
	// myRoute.HandleFunc("/listar", controllers.ListUsers()).Methods("GET")
}