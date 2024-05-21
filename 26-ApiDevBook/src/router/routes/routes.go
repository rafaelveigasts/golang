package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct that represents a route
type Route struct {
	URI string
	Method string
	Function func(w http.ResponseWriter, r *http.Request)
	RequiresAuth bool 
}

func SetupRoute(r * mux.Router) *mux.Router {
	var routes = append(userRoutes, LoginRoute)
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}