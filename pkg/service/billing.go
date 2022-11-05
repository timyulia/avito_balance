package service

import (
	"balance"
	"balance/pkg/repository"
	"errors"
)

type BillingService struct {
	repo repository.Billing
}

func NewBillingService(repo repository.Billing) *BillingService {
	return &BillingService{repo: repo}
}

func (s *BillingService) AddMoney(account balance.User) error {
	if account.Amount < 0 {
		return errors.New("negative amount")
	}
	return s.repo.AddMoney(account)
}

func (s *BillingService) Reserve(userId int, ord balance.Order) error {
	if ord.Amount < 0 {
		return errors.New("negative amount")
	}
	return s.repo.Reserve(userId, ord)
}

func (s *BillingService) WriteOff(userId int, ord balance.Order) error {
	if ord.Amount < 0 {
		return errors.New("negative amount")
	}
	return s.repo.WriteOff(userId, ord)
}

func (s *BillingService) GetBalance(id int) (int, error) {

	return s.repo.GetBalance(id)
}

func (s *BillingService) Dereserve(orderId, userId int) error {
	return s.repo.Dereserve(orderId, userId)
}
