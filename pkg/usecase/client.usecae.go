package usecase

import (
	"context"
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
)

type ClientUseCase struct {
	ClientRepo interfaces.IClientRepository
}

func (c *ClientUseCase) RegisterClient(ctx context.Context, createClientDto entity.CreateClientDto) (entity.CreateClientResponse, error) {
	return c.ClientRepo.Create(ctx, createClientDto)
}
