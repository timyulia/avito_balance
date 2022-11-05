package repository

import (
	"balance"
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	"github.com/jmoiron/sqlx"
)

type BillingPostgres struct {
	db *sqlx.DB
}

func NewBillingPostgres(db *sqlx.DB) *BillingPostgres {
	return &BillingPostgres{db: db}
}

func (r *BillingPostgres) AddMoney(account balance.User) error {
	query := fmt.Sprintf("INSERT INTO %s as u  (id, amount) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE"+
		" SET amount = EXCLUDED.amount+u.amount", usersTable)
	_, err := r.db.Exec(query, account.Id, account.Amount)
	return err
}

func (r *BillingPostgres) Reserve(userId int, ord balance.Order) error {
	currAmount := 0
	query := fmt.Sprintf("SELECT amount FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&currAmount, query, userId)
	if err != nil {
		return err
	}
	if currAmount < ord.Amount {
		return errors.New("not enough money")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	addQuery := fmt.Sprintf("INSERT INTO %s (user_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4)",
		reservedTable)
	_, err = tx.Exec(addQuery, userId, ord.ServiceId, ord.OrderId, ord.Amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}

	cutQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE id=$2",
		usersTable)
	_, err = tx.Exec(cutQuery, ord.Amount, userId)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	return tx.Commit()
}

func (r *BillingPostgres) WriteOff(userId int, ord balance.Order) error {
	err := r.checkRes(userId, ord)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	reportQuery := fmt.Sprintf("INSERT INTO %s ( service_id, amount, date) VALUES ($1, $2, current_timestamp)",
		reportTable)
	_, err = tx.Exec(reportQuery, ord.ServiceId, ord.Amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}

	cutQuery := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1",
		reservedTable)
	_, err = tx.Exec(cutQuery, ord.OrderId)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	historyQuery := fmt.Sprintf("INSERT INTO %s ( user_id, reason, amount, date) VALUES ($1, $2, $3, current_timestamp)",
		historyTable)
	reason := fmt.Sprintf("write-off for the service %d", ord.ServiceId)
	_, err = tx.Exec(historyQuery, userId, reason, ord.Amount)

	if err != nil {
		err3 := tx.Rollback()
		if err3 != nil {
			return err3
		}
		return err
	}
	return tx.Commit()
}

func (r *BillingPostgres) GetBalance(id int) (int, error) {
	query := fmt.Sprintf("SELECT amount FROM %s WHERE id=$1", usersTable)
	amount := 0
	err := r.db.Get(&amount, query, id)
	return amount, err
}

func (r *BillingPostgres) checkRes(userId int, ord balance.Order) error {
	reserv := 0
	checkQuery := fmt.Sprintf("SELECT amount FROM %s WHERE order_id=$1 AND user_id=$2 AND service_id=$3", reservedTable)
	err := r.db.Get(&reserv, checkQuery, ord.OrderId, userId, ord.ServiceId)
	if err != nil {
		return err
	}
	if reserv < ord.Amount {
		return errors.New("not enough money reserved")
	}
	return nil
}

func (r *BillingPostgres) Dereserve(orderId, userId int) error {

	amount := 0
	query := fmt.Sprintf("SELECT amount FROM %s WHERE order_id=$1")
	err := r.db.Get(&amount, query, orderId)
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	delQuery := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1 AND user_id=$2 RETURNING amount", reservedTable)
	row := tx.QueryRow(delQuery, orderId, userId)
	err = row.Scan(&amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}

	addQuery := fmt.Sprintf("UPDATE %s SET amount = amount+$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(addQuery, amount, userId)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	return tx.Commit()

}
