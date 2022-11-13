package main

import (
	"fmt"
	"net/http"
	"os"
	"waysbooks/database"
	"waysbooks/pkg/mysql"
	"waysbooks/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	var allowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})
	var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")

	fmt.Println("server running on port " + port)
	http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(r))
}