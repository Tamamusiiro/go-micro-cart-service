package respository

import (
	"errors"

	"github.com/Tamamusiiro/go-micro-cart-service/domian/model"
	"gorm.io/gorm"
)

type ICartRepository interface {
	AutoMigrate() error
	FindAll(int64) ([]model.Cart, error)
	FindCartById(int64) (model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	UpdateCart(*model.Cart) error
	DeleteCart(int64) error
	CleanCart(int64) error
	Incr(cartId int64, num int64) error
	Decr(cartId int64, num int64) error
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDB: db}
}

func (r *CartRepository) AutoMigrate() error {
	return r.mysqlDB.AutoMigrate(&model.Cart{})
}

func (r *CartRepository) FindAll(userId int64) (carts []model.Cart, err error) {
	return carts, r.mysqlDB.Model(&model.Cart{}).Where("user_id = ?", userId).Find(&carts).Error
}

func (r *CartRepository) FindCartById(id int64) (cart model.Cart, err error) {
	return cart, r.mysqlDB.First(&cart, id).Error
}

func (r *CartRepository) CreateCart(cart *model.Cart) (id int64, err error) {
	return cart.ID, r.mysqlDB.FirstOrCreate(cart, model.Cart{
		UserId: cart.UserId, ProductId: cart.ProductId, SizeId: cart.SizeId,
	}).Error
}

func (r *CartRepository) UpdateCart(cart *model.Cart) error {
	return r.mysqlDB.Save(cart).Error
}

func (r *CartRepository) DeleteCart(id int64) error {
	return r.mysqlDB.Delete(&model.Cart{}, id).Error
}

func (r *CartRepository) CleanCart(userId int64) error {
	return r.mysqlDB.Where("user_id", userId).Delete(&model.Cart{}).Error
}

func (r *CartRepository) Incr(cart_id int64, num int64) error {
	return r.mysqlDB.Model(&model.Cart{ID: cart_id}).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (r *CartRepository) Decr(cart_id int64, num int64) error {
	tx := r.mysqlDB.Model(&model.Cart{ID: cart_id}).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("decr failed")
	}
	return nil
}
