package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/middleware"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart/add/{bookID}", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/carts", h.FindCarts).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.GetCartByUser)).Methods("GET")
	r.HandleFunc("/cart/delete/{cartId}", middleware.Auth(h.DeleteCartByID)).Methods("DELETE")
	r.HandleFunc("/cart/clear/{userId}", middleware.Auth(h.DeleteCartByUser)).Methods("DELETE")

}