package transactiondto

import "waysbooks/models"

type CreateTransactionRequest struct {
	UserID   int               `Json:"userId"`
	Status   string            `json:"status"`
	Book 	 []models.Book 	   `json:"books"`
	Totalpayment int 		   `json:"totalpayment"`
}