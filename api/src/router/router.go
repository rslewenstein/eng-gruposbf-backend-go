package router

import (
	"github.com/gorilla/mux"
	"go.mod/src/router/routers"
)

func GenerateRouters() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigureRouter(r)
}
