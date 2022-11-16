package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart models.Cart) (models.Cart, error)
	FindCarts(cart []models.Cart) ([]models.Cart, error)
	GetCart(cart models.Cart, ID int) (models.Cart, error)
	GetCartByUser(cart []models.Cart, userID int) ([]models.Cart, error)
	GetCartsByCurrentUser(userID int) ([]models.Cart, error)
	GetCartExist(userID int, BookID int) (models.Cart, error)
	UpdateCartQty(cart models.Cart, ID int) (models.Cart, error)
	DeleteCartByID(cart models.Cart, ID int) (models.Cart, error)
	DeleteCartByUser(cart models.Cart, userID int) error
	GetPrice(bookID int) (book models.Book, err error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddToCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Preload("User").Preload("Books").Error
	return cart, err
}

func (r *repository) GetPrice(bookID int) (book models.Book, err error) {
	err = r.db.First(&book, bookID).Error
	return book, err
}

func (r *repository) FindCarts(carts []models.Cart) ([]models.Cart, error) {
	
	err := r.db.Preload("User").Preload("Books").Find(&carts).Error
	return carts, err
}

func (r *repository) GetCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Preload("User").Preload("Books").First(&cart, ID).Error
	return cart, err
}

func (r *repository) GetCartByUser(cart []models.Cart, userID int) ([]models.Cart, error) {
	err := r.db.Preload("User").Preload("Books").Find(&cart, "user_id=?", userID).Error

	return cart, err
}

func (r *repository) GetCartsByCurrentUser(userID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("User").Preload("Books").Where("user_id = ?", userID).Find(&carts).Error

	return carts, err
}

func (r *repository) GetCartExist(userID int, BookID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("User").Preload("Books").Where("user_id = ? AND book_id = ?", userID, BookID).First(&cart).Error
	return cart, err
}

func (r *repository) UpdateCartQty(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Model(&cart).Where("id = ?", ID).Preload("User").Preload("Books").Updates(&cart).Error
	return cart, err
}

func (r *repository) DeleteCartByID(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart, "id=?", ID).Preload("Books").Preload("User").Error
	return cart, err
}

func (r *repository) DeleteCartByUser(cart models.Cart, userID int) error {
	err := r.db.Preload("Books").Preload("User").Delete(&cart, "user_id=?", userID).Error
	return err
}
