package router

import (
	"api-devbook/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()	

	return routes.SetupRoute(r)
}