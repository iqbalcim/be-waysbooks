package models

import "time"

type Transaction struct {
	ID     		int          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID 		int          `json:"userId"`
	User   		UserResponse `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	Totalpayment int          `json:"totalpayment"`
	Status    	string     `json:"status"`
	Books  		[]Book `json:"books" gorm:"many2many:transaction_books;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt 	time.Time  `json:"createdAT" gorm:"default:Now()"`
	UpdateAt  	time.Time  `json:"updatedAt" gorm:"default:null"`
}

// gorm:"foreignKey:UserID;references:ID"