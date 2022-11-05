package service

import (
	"balance"
	"balance/pkg/repository"
)

type Billing interface {
	AddMoney(account balance.User) error
	Reserve(userId int, ord balance.Order) error
	WriteOff(userId int, ord balance.Order) error
	GetBalance(id int) (int, error)
	Dereserve(orderId, userId int) error
}

type Service struct {
	Billing
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Billing: NewBillingService(r.Billing),
	}
}
