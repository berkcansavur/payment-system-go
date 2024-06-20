package interfaces

import (
	"context"
	"payment-system/domain/entity"
)

type IClientRepository interface {
	Create(ctx context.Context,createClientDto entity.CreateClientDto) (entity.CreateClientResponse, error)
}