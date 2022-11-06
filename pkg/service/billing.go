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

func (s *BillingService) Reserve(ord balance.Order) error {
	if ord.Amount < 0 {
		return errors.New("negative amount")
	}
	return s.repo.Reserve(ord)
}

func (s *BillingService) WriteOff(ord balance.Order) error {
	if ord.Amount < 0 {
		return errors.New("negative amount")
	}
	return s.repo.WriteOff(ord)
}

func (s *BillingService) GetBalance(id int) (int, error) {

	return s.repo.GetBalance(id)
}

func (s *BillingService) Dereserve(ord balance.Order) error {
	return s.repo.Dereserve(ord)
}
