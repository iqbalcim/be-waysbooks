package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/middleware"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router){
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(UserRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", (middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}