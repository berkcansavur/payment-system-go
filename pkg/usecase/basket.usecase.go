package usecase

import (
	"context"
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
)

type BasketUseCase struct {
	BasketRepo interfaces.IBasketRepository
}
func (b *BasketUseCase) Create(ctx context.Context, createBasketDto entity.BasketDto)(*entity.BasketDto, error){
	return b.BasketRepo.Create(ctx, createBasketDto)
}