package service

import (
	"github.com/Tamamusiiro/go-micro-cart-service/domian/model"
	"github.com/Tamamusiiro/go-micro-cart-service/domian/respository"
)

type ICartService interface {
	FindAll(int64) ([]model.Cart, error)
	FindCartById(int64) (model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	UpdateCart(*model.Cart) error
	DeleteCart(int64) error
	CleanCart(int64) error
	Incr(cartId int64, num int64) error
	Decr(cartId int64, num int64) error
}

type CartService struct {
	CartRepository respository.ICartRepository
}

func NewCartService(repository respository.ICartRepository) ICartService {
	return &CartService{CartRepository: repository}
}

func (s *CartService) CreateCart(cart *model.Cart) (int64, error) {
	return s.CartRepository.CreateCart(cart)
}

func (s *CartService) UpdateCart(cart *model.Cart) error {
	return s.CartRepository.UpdateCart(cart)
}

func (s *CartService) DeleteCart(id int64) error {
	return s.CartRepository.DeleteCart(id)
}

func (s *CartService) FindCartById(id int64) (model.Cart, error) {
	return s.CartRepository.FindCartById(id)
}

func (s *CartService) FindAll(userId int64) ([]model.Cart, error) {
	return s.CartRepository.FindAll(userId)
}

func (s *CartService) CleanCart(userId int64) error {
	return s.CartRepository.CleanCart(userId)
}

func (s *CartService) Incr(cart_id int64, num int64) error {
	return s.CartRepository.Incr(cart_id, num)
}

func (s *CartService) Decr(cart_id int64, num int64) error {
	return s.CartRepository.Decr(cart_id, num)
}
