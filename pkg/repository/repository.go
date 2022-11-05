package repository

import (
	"balance"
	"github.com/jmoiron/sqlx"
)

type Billing interface {
	AddMoney(account balance.User) error
	Reserve(userId int, ord balance.Order) error
	WriteOff(userId int, ord balance.Order) error
	GetBalance(id int) (int, error)
	Dereserve(orderId, userId int) error
}

type Repository struct {
	Billing
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Billing: NewBillingPostgres(db)}
}
