package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/middleware"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router){
	BookRepository := repositories.RepositoryBook(mysql.DB)
	h := handlers.HandlerBook(BookRepository)

	r.HandleFunc("/books", h.FindBooks).Methods("GET")
	r.HandleFunc("/book/{id}", h.GetBook).Methods("GET")
	r.HandleFunc("/book", middleware.Auth(middleware.UploadPdf(middleware.UploadFile(h.CreateBook)))).Methods("POST")
	r.HandleFunc("/book/{id}", middleware.Auth(middleware.UploadPdf(middleware.UploadFile(h.UpdateBook)))).Methods("PATCH")
	r.HandleFunc("/book/{id}", middleware.Auth(h.DeleteBook)).Methods("DELETE")
}