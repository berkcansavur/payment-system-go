package usecase

import (
	"context"
	"fmt"
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
	"strconv"
)

type BasketUseCase struct {
	BasketRepo interfaces.IBasketRepository
	ClientRepo interfaces.IClientRepository
}
func (b *BasketUseCase) Create(ctx context.Context, createBasketDto entity.BasketDto, clientId string)(*entity.BasketDto, error){
	client, err := b.ClientRepo.GetById(ctx, clientId)
	if err!= nil {
        return nil, err
    }
	
	totalPrice, err := b.CalculateTotalPrice(&createBasketDto)
	if err!= nil {
        return nil, err
    }

	createBasketDto.Client = *client
	createBasketDto.IsActive = true
	createBasketDto.TotalPrice = totalPrice

	return b.BasketRepo.Create(ctx, createBasketDto)
}
func (b *BasketUseCase) CalculateTotalPrice(basket *entity.BasketDto) (string, error) {
	totalPrice := 0.0
	for _, item := range basket.BasketItems {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return "", fmt.Errorf("invalid price format for item %s: %v", item.Id, err)
		}
		totalPrice += price
	}
	basket.TotalPrice = fmt.Sprintf("%.2f", totalPrice)
	return basket.TotalPrice, nil
}
// func (b *BasketUseCase) Abort(ctx context.Context, id string)(*entity.BasketDto, error){
//     return b.BasketRepo.UpdateById(ctx, id,)
// }