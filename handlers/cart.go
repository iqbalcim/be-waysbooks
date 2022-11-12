package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	cartdto "waysbooks/dto/cart"
	dto "waysbooks/dto/result"
	"waysbooks/models"
	"waysbooks/repositories"

	"github.com/gorilla/mux"
)

type handlerCart struct {
	cartRepository repositories.CartRepository
}

func HandlerCart (cartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{cartRepository}
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	qty, _ := strconv.Atoi(r.FormValue("qty"))
	totalPayment, _ := strconv.Atoi(r.FormValue("totalPayment"))
	bookId, _ := strconv.Atoi(r.FormValue("book_id"))
	userId, _ := strconv.Atoi(r.FormValue("user_id"))

	request := cartdto.CreateCartRequest{
		Qty:       			qty,
		TotalPayment: 		totalPayment,
		BookID: 			bookId,
		UserID: 			userId,
	}

	cart := models.Cart{
		Qty:       					request.Qty,
		TotalPayment: 				request.TotalPayment,
		BookID: 					request.BookID,
		UserID: 					request.UserID,
	}

	// err := json.NewDecoder(r.Body).Decode(&cart)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	cartExist , err := h.cartRepository.GetCartExist(cart.UserID, cart.BookID)
	
	if err == nil {
		cartExist.TotalPayment = cartExist.TotalPayment + (cartExist.TotalPayment / cartExist.Qty)
		cartExist.Qty = cartExist.Qty + 1
		cart, err = h.cartRepository.UpdateCartQty(cartExist, cartExist.ID)
	} else {
		cart, err = h.cartRepository.AddToCart(cart)
	}

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


func (h *handlerCart) UpdateCartQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	updateType := r.URL.Query()["update"]
	cartId, _ := strconv.Atoi(mux.Vars(r)["cartId"])

	var cartUpdate models.Cart
	var err error
	cartUpdate, err = h.cartRepository.GetCart(cartUpdate, cartId)

	if len(updateType) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if updateType[0] == "add" {
		cartUpdate.TotalPayment = cartUpdate.TotalPayment + (cartUpdate.TotalPayment / cartUpdate.Qty)
		cartUpdate.Qty = cartUpdate.Qty + 1

	} else {
		cartUpdate.TotalPayment = cartUpdate.TotalPayment - (cartUpdate.TotalPayment / cartUpdate.Qty)
		cartUpdate.Qty = cartUpdate.Qty - 1
	}
	cartUpdate.UpdateAt = time.Now()

	cart, err := h.cartRepository.UpdateCartQty(cartUpdate, cartId)
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

