package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Representa as rotas da API
type Router struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

func ConfigureRouter(r *mux.Router) *mux.Router {
	routers := productRouters

	for _, route := range routers {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
