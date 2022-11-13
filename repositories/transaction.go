package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	GetTransactionByUser(userID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Books").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Books").Preload("User").Find(&transaction, "id =?", ID).Error

	return transaction, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Books").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Books").Preload("User").Find(&transactions, "user_id=?", userID).Error

	return transactions, err
}

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Products").First(&transaction, ID)
	var product models.Book

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var transactionItem models.Transaction
		for _, tp := range transactionItem.Books {
			r.db.First(&product, tp.ID)
		}
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}