package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router){
	authRepository := repositories.RepositoryAuth(mysql.DB)

	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

}