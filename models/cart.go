package models

import "time"

type Cart struct {
	ID         		int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Qty        		int          `json:"qty"`
	TotalPayment 	int          `json:"totalPayment"`
	BookID  		int          `json:"bookId"`
	Books	  		Book     	`json:"books" gorm:"foreignKey:BookID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    		int          `json:"userId" `
	User       		UserResponse `json:"user"`
	CreatedAt  		time.Time    `json:"createdAT" gorm:"default:Now()"`
	UpdateAt   		time.Time    `json:"updatedAt" gorm:"default:null"`
}

type CartUpdateRequest struct {
	ID         int `json:"id"`
	Qty        int `json:"qty"`
	TotalPayment int `json:"totalPayment"`
}

func (CartUpdateRequest) TableName() string {
	return "carts"
}