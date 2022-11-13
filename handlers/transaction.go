package handlers

import (
	"encoding/json"
	"net/http"
	dto "waysbooks/dto/result"
	transactiondto "waysbooks/dto/transaction"
	"waysbooks/models"
	"waysbooks/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := transactiondto.CreateTransactionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{
		UserID: userId,
		Status: request.Status,
		Books: request.Book,
		Totalpayment: request.Totalpayment,
	}

	transaction, err = h.TransactionRepository.CreateTransaction(transaction)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)

}