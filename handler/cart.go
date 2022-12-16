package handler

import (
	"context"

	"github.com/Tamamusiiro/go-micro-cart-service/domian/model"
	"github.com/Tamamusiiro/go-micro-cart-service/domian/service"
	pb "github.com/Tamamusiiro/go-micro-cart-service/proto/cart"
	common "github.com/Tamamusiiro/go-micro-common"
)

type CartHandler struct {
	CartService service.ICartService
}

func (h *CartHandler) AddCart(ctx context.Context, req *pb.CartInfo, res *pb.CartId) error {
	cart := model.Cart{}
	err := common.Transform(req, &cart)
	if err != nil {
		return err
	}
	res.CartId, err = h.CartService.CreateCart(&cart)
	if err != nil {
		return err
	}
	return nil
}

func (h *CartHandler) CleanCart(ctx context.Context, req *pb.UserId, res *pb.Response) error {
	return h.CartService.CleanCart(req.UserId)
}

func (h *CartHandler) DeleteCart(ctx context.Context, req *pb.CartId, res *pb.Response) error {
	return h.CartService.DeleteCart(req.CartId)
}

func (h *CartHandler) GetAll(ctx context.Context, req *pb.UserId, res *pb.CartList) error {
	carts, err := h.CartService.FindAll(req.UserId)
	if err != nil {
		return err
	}
	return cartsToResponse(carts, res)
}

func (h *CartHandler) Incr(ctx context.Context, req *pb.IncrRequest, res *pb.Response) error {
	return h.CartService.Incr(req.CartId, req.Num)
}

func (h *CartHandler) Decr(ctx context.Context, req *pb.DecrRequest, res *pb.Response) error {
	return h.CartService.Decr(req.CartId, req.Num)
}

func cartsToResponse(carts []model.Cart, res *pb.CartList) error {
	for _, cart := range carts {
		to := pb.CartInfo{}
		if err := common.Transform(cart, &to); err != nil {
			return err
		}
		res.Carts = append(res.Carts, &to)
	}
	return nil
}
