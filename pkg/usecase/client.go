package usecase

import (
	"context"
	"fmt"
	"log"
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
)

type ClientUseCase struct {
	ClientRepo interfaces.IClientRepository
	BasketRepo interfaces.IBasketRepository
	PaymentRepo interfaces.IPaymentRepository
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
	log.Println("Clients current cards",updatedClient.Cards)
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
	log.Println("Clients current cards",updatedClient.Cards)
	return updatedClient.Cards, nil
}
func (c *ClientUseCase) CheckoutClientsActiveBasket(ctx context.Context, clientId string, checkoutDto entity.CheckoutDto) (entity.CreatePaymentResult, error) {
	basket, err := c.BasketRepo.GetActiveBasketByClientId(ctx, clientId)
	fmt.Print(basket)
	if err != nil {
		return entity.CreatePaymentResult{}, err
	}
	client, err := c.Get(ctx, clientId)
	if err != nil {
		return entity.CreatePaymentResult{}, err
	}
	buyer := entity.Buyer{
		Id: client.Id,
		Name: client.Name,
        Surname: client.Surname,
		IdentityNumber: client.IdentityNumber,
        Email: client.Email,
        GsmNumber: client.GsmNumber,
		RegistrationDate: client.RegistrationDate,
		LastLoginDate: client.LastLoginDate,
        RegistrationAddress: client.RegistrationAddress,
        City: client.City,
		Country: client.Country,
        ZipCode: client.ZipCode,
		Ip: client.Ip,
	}
	createPaymentRequest := entity.CreatePaymentRequest{
		Locale: checkoutDto.Locale,
		ConversationID: checkoutDto.ConversationID,
        Price: checkoutDto.Price,
        PaidPrice: checkoutDto.PaidPrice,
        Installment: checkoutDto.Installment,
        PaymentChannel: checkoutDto.PaymentChannel,
		BasketID: basket.Id,
        PaymentGroup: checkoutDto.PaymentGroup,
        PaymentCard: client.Cards[0],
        Buyer: buyer,
        ShippingAddress: client.ShippingAddress,
		BillingAddress: client.BillingAddress,
		BasketItems: basket.BasketItems,
        Currency: checkoutDto.Currency,

	}
	fmt.Print(createPaymentRequest)
	log.Println("Create Payment Request",createPaymentRequest)
	return c.PaymentRepo.CreatePayment(createPaymentRequest)
}