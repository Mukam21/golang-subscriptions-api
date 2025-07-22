package service

import (
	"errors"
	"time"

	"golang-subscriptions-api/internal/model"
	"golang-subscriptions-api/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(sub *model.Subscription) error {
	return s.repo.Create(sub)
}

func (s *Service) GetAll() ([]model.Subscription, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id string) (*model.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id string, input *model.Subscription) error {
	sub, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if sub == nil {
		return errors.New("subscription not found")
	}

	sub.ServiceName = input.ServiceName
	sub.Price = input.Price
	sub.UserID = input.UserID
	sub.StartDate = input.StartDate
	sub.EndDate = input.EndDate

	return s.repo.Update(sub)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) TotalSum(userID, serviceName string, from, to *time.Time) (int64, error) {
	return s.repo.SumPrice(userID, serviceName, from, to)
}
