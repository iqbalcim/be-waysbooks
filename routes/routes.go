package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	BookRoutes(r)
	CartRoutes(r)
}