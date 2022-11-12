package cartdto

type CreateCartRequest struct {
	Qty          int `json:"qty"`
	TotalPayment int `json:"totalPayment"`
	BookID       int `json:"bookId"`
	UserID       int `json:"userId"`
}