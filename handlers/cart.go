package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	cartdto "waysbooks/dto/cart"
	dto "waysbooks/dto/result"
	"waysbooks/models"
	"waysbooks/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	cartRepository repositories.CartRepository
}

func HandlerCart (cartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{cartRepository}
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	bookId, _ := strconv.Atoi(mux.Vars(r)["bookID"])
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	price, _ := h.cartRepository.GetPrice(bookId)

	request := cartdto.CreateCartRequest{
		BookID: bookId,
		UserID: userId,
		Price:  price.Price,
		Qty: 1,
	}

	cart := models.Cart{
		BookID: request.BookID,
		UserID: request.UserID,
		Price:  request.Price,
		Qty:    request.Qty,
	}


	cart, err := h.cartRepository.AddToCart(cart)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCart) FindCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var carts []models.Cart

	carts, err := h.cartRepository.FindCarts(carts)
	if err == nil {
		if len(carts) == 0 {
			w.WriteHeader(http.StatusNotFound)
			errorMessage := "Carts not found!"
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: errorMessage}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: carts}
	json.NewEncoder(w).Encode(response)
	
}

func (h *handlerCart) GetCartByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var carts []models.Cart
	var err error

	carts, err = h.cartRepository.GetCartByUser(carts, userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: carts}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) GetCartsByCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	id := int(userInfo["id"].(float64))

	var carts []models.Cart
	carts, err := h.cartRepository.GetCartByUser(carts,id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: carts}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteCartByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cartId, _ := strconv.Atoi(mux.Vars(r)["cartId"])

	var cartDeleted models.Cart
	var err error
	cartDeleted, err = h.cartRepository.DeleteCartByID(cartDeleted, cartId)

	response := map[string]models.Cart{
		"cartDeleted": cartDeleted,
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	responsed := dto.SuccessResult{Code: "Success", Data: response}
	json.NewEncoder(w).Encode(responsed)

}

func (h *handlerCart) DeleteCartByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	var cartDeleted models.Cart

	 err := h.cartRepository.DeleteCartByUser(cartDeleted, userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: err}
	json.NewEncoder(w).Encode(response)
}

