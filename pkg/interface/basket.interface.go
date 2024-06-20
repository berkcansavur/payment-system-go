package interfaces

import (
	"context"
	"payment-system/domain/entity"
)

type IBasketRepository interface {
	Create(ctx context.Context, createBasketDto entity.BasketDto) (*entity.BasketDto, error)
	GetById(ctx context.Context, id string) (*entity.BasketDto, error)
	RemoveById(ctx context.Context, id string) (*entity.BasketDto, error)
	UpdateById(ctx context.Context, id string, updateBasketDto entity.BasketDto) (*entity.BasketDto, error)
}