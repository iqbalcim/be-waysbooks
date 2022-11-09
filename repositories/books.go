package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	GetBook(ID int)(models.Book, error)
	CreateBook(book models.Book) (models.Book ,error)
	UpdateBook(book models.Book, ID int) (models.Book ,error)
	DeleteBook(book models.Book, ID int) (models.Book ,error)
}

func RepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBooks() ([]models.Book, error){
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) GetBook(ID int) (models.Book, error){
	var book models.Book
	err := r.db.First(&book, ID).Error

	return book ,err
}

func (r *repository) CreateBook(book models.Book) (models.Book, error){
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) UpdateBook(book models.Book, ID int) (models.Book, error){
	err := r.db.Where("id = ?", ID).Updates(&book).Error

	return book, err
}

func (r *repository) DeleteBook(book models.Book, ID int) (models.Book, error){
	err := r.db.Delete(&book).Error

	return book, err
}