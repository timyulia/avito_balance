package service

import (
	"balance"
	"balance/pkg/repository"
)

type Billing interface {
	AddMoney(account balance.User) error
	Reserve(ord balance.Order) error
	WriteOff(ord balance.Order) error
	GetBalance(id int) (int, error)
	Dereserve(ord balance.Order) error
}

type Service struct {
	Billing
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Billing: NewBillingService(r.Billing),
	}
}
