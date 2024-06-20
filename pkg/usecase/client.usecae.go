package usecase

import (
	"context"
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
)

type ClientUseCase struct {
	ClientRepo interfaces.IClientRepository
}

func (c *ClientUseCase) Register(ctx context.Context, createClientDto entity.ClientDto) (*entity.ClientDto, error) {
	return c.ClientRepo.Create(ctx, createClientDto)
}
func (c *ClientUseCase) Get(ctx context.Context, id string) (*entity.ClientDto, error) {
	return c.ClientRepo.GetById(ctx, id)
}
func (c *ClientUseCase) Update(ctx context.Context, id string, updateClientDto entity.ClientDto) (*entity.ClientDto, error) {
	return c.ClientRepo.UpdateById(ctx, id, updateClientDto)
}
func (c *ClientUseCase) Delete(ctx context.Context, id string) (*entity.ClientDto, error) {
	return c.ClientRepo.RemoveById(ctx, id)
}