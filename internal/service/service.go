package service

import (
	"context"

	"github.com/sonochiwa/avito-shop/api/rest/models"
	"github.com/sonochiwa/avito-shop/internal/converter"
	"github.com/sonochiwa/avito-shop/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetInfo(ctx context.Context, userID int) (models.GetInfoHandlerResponse, error) {
	data, err := s.repo.GetInfo(ctx, userID)
	if err != nil {
		return models.GetInfoHandlerResponse{}, err
	}

	return converter.ConvertGetInfoToResponse(data), nil
}

func (s Service) SendCoin() {
	//TODO implement me
	panic("implement me")
}

func (s Service) BuyItem() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Auth() {
	//TODO implement me
	panic("implement me")
}
