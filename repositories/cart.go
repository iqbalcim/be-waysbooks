package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart models.Cart) (models.Cart, error)
	FindCarts(cart []models.Cart) ([]models.Cart, error)
	GetCart(cart models.Cart, ID int) (models.Cart, error)
	GetCartExist(userID int, productId int) (models.Cart, error)
	UpdateCartQty(cart models.Cart, ID int) (models.Cart, error)
	DeleteCartByID(cart models.Cart, ID int) (models.Cart, error)
	DeleteCartByUser(cart models.Cart, userID int) error
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddToCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Preload("User").Preload("Books").Error
	return cart, err
}

func (r *repository) FindCarts(carts []models.Cart) ([]models.Cart, error) {
	
	err := r.db.Preload("User").Preload("Books").Find(&carts).Error
	return carts, err
}

func (r *repository) GetCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Preload("User").Preload("Books").First(&cart, ID).Error
	return cart, err
}

func (r *repository) GetCartExist(userID int, productId int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("User").Preload("Books").Where("user_id = ? AND book_id = ?", userID, productId).First(&cart).Error
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
	err := r.db.Where("user_id = ?", userID).Delete(&cart).Error
	return err
}
