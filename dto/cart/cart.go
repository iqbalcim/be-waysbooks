package cartdto

type CreateCartRequest struct {
	Qty    int `json:"qty"`
	Price  int `json:"price"`
	BookID int `json:"bookId"`
	UserID int `json:"userId"`
}