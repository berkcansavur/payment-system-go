package usecase

import (
	"context"
	"fmt"
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
func (c *ClientUseCase) GetClientsCards(ctx context.Context, id string) ([]entity.PaymentCard, error) {
	return c.ClientRepo.GetCards(ctx, id)
}
func (c *ClientUseCase) AddCard(ctx context.Context, id string, card entity.PaymentCard) ([]entity.PaymentCard, error) {
	client, err := c.Get(ctx, id)
	if err!= nil {
        return nil, err
    }
	for _, existingCard := range client.Cards {
		if existingCard.CardNumber == card.CardNumber {
			return nil, fmt.Errorf("card with number %s already exists", card.CardNumber)
		}
	}
	cards := append(client.Cards, card)
	client.Cards = cards
	updatedClient, err := c.ClientRepo.UpdateById(ctx, id, *client)
	if err != nil {
		return nil, err
	}
	return updatedClient.Cards, nil
}
func (c *ClientUseCase) RemoveCard(ctx context.Context, id string, card entity.PaymentCard) ([]entity.PaymentCard, error) {
	client, err := c.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	cardFound := false
	updatedCards := []entity.PaymentCard{}
	for _, existingCard := range client.Cards {
		if existingCard.CardNumber == card.CardNumber {
			cardFound = true
		} else {
			updatedCards = append(updatedCards, existingCard)
		}
	}

	if !cardFound {
		return nil, fmt.Errorf("card with number %s not found", card.CardNumber)
	}

	client.Cards = updatedCards
	updatedClient, err := c.ClientRepo.UpdateById(ctx, id, *client)
	if err != nil {
		return nil, err
	}

	return updatedClient.Cards, nil
}
