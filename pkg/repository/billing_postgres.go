package repository

import (
	"balance"
	"database/sql"
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
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s as u  (id, amount) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE"+
		" SET amount = EXCLUDED.amount+u.amount", usersTable)
	_, err = tx.Exec(query, account.Id, account.Amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	account.Reason = "account replenishment: " + account.Reason
	query = fmt.Sprintf("INSERT INTO %s  (user_id, reason, amount, date) VALUES ($1, $2, $3, current_timestamp)", historyTable)
	_, err = tx.Exec(query, account.Id, account.Reason, account.Amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	return tx.Commit()

}

func (r *BillingPostgres) Reserve(ord balance.Order) error {
	currAmount := 0
	query := fmt.Sprintf("SELECT amount FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&currAmount, query, ord.UserId)
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
	_, err = tx.Exec(addQuery, ord.UserId, ord.ServiceId, ord.OrderId, ord.Amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}

	cutQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE id=$2",
		usersTable)
	_, err = tx.Exec(cutQuery, ord.Amount, ord.UserId)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	return tx.Commit()
}

func (r *BillingPostgres) WriteOff(ord balance.Order) error {
	err := r.checkRes(ord)
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
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}

	historyQuery := fmt.Sprintf("INSERT INTO %s ( user_id, reason, amount, date) VALUES ($1, $2, $3, current_timestamp)",
		historyTable)
	name, err := r.getName(ord.ServiceId)
	if err != nil {
		name = "without name"
	}
	reason := fmt.Sprintf("write-off for the service %s", name)
	_, err = tx.Exec(historyQuery, ord.UserId, reason, ord.Amount)

	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
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

func (r *BillingPostgres) checkRes(ord balance.Order) error {
	reserv := 0
	checkQuery := fmt.Sprintf("SELECT amount FROM %s WHERE order_id=$1 AND user_id=$2 AND service_id=$3", reservedTable)
	err := r.db.Get(&reserv, checkQuery, ord.OrderId, ord.UserId, ord.ServiceId)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("fields do not match")
		}
		return err
	}
	if reserv < ord.Amount {
		return errors.New("not enough money reserved")
	}
	return nil
}

func (r *BillingPostgres) Dereserve(ord balance.Order) error {

	amount := 0
	query := fmt.Sprintf("SELECT amount FROM %s WHERE order_id=$1")
	err := r.db.Get(&amount, query, ord.OrderId)
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	delQuery := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1 AND user_id=$2 RETURNING amount", reservedTable)
	row := tx.QueryRow(delQuery, ord.OrderId, ord.UserId)
	err = row.Scan(&amount)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}

	addQuery := fmt.Sprintf("UPDATE %s SET amount = amount+$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(addQuery, amount, ord.UserId)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	return tx.Commit()

}

func (r *BillingPostgres) getName(id int) (string, error) {
	query := fmt.Sprintf("SELECT name FROM %s WHERE id=$1")
	var name string
	err := r.db.Get(&name, query, id)
	return name, err
}
