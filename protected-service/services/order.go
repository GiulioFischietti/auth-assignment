package services

import (
	"context"

	"protected-service/models"
	"protected-service/repositories"
)

type OrderService struct {
	repository *repositories.OrderRepository
}

func NewOrderService(
	repository *repositories.OrderRepository,
) *OrderService {

	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) GetOrders(
	ctx context.Context,
	userID int64,
) ([]models.Order, error) {

	return s.repository.FindByCustomerID(
		ctx,
		userID,
	)
}
