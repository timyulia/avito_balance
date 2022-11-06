package repository

import (
	"balance"
	"github.com/jmoiron/sqlx"
)

type Billing interface {
	AddMoney(account balance.User) error
	Reserve(ord balance.Order) error
	WriteOff(ord balance.Order) error
	GetBalance(id int) (int, error)
	Dereserve(ord balance.Order) error
}

type Repository struct {
	Billing
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Billing: NewBillingPostgres(db)}
}
