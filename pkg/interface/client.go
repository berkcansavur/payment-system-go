package interfaces

import (
	"context"
	"payment-system/domain/entity"

	"github.com/berkcansavur/iyzico-authorization"
)

type IClientRepository interface {
	Create(ctx context.Context, createClientDto entity.ClientDto) (*entity.ClientDto, error)
	GetById(ctx context.Context, id string) (*entity.ClientDto, error)
	RemoveById(ctx context.Context, id string) (*entity.ClientDto, error)
	UpdateById(ctx context.Context, id string, updateClientDto entity.ClientDto) (*entity.ClientDto, error)
	GetCards(ctx context.Context, id string) ([]iyzico.PaymentCard, error)
}