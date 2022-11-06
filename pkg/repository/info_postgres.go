package repository

import (
	"balance"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/sqltocsv"
	_ "time"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (r *InfoPostgres) MakeReport(year, month int) error {
	//var rows []balance.Report
	query := fmt.Sprintf("SELECT s.name, r.service_id, SUM(amount) FROM %s r LEFT JOIN %s s ON  r.service_id="+
		"s.service_id WHERE EXTRACT(YEAR FROM date)=$1 "+
		"AND EXTRACT(MONTH FROM date)=$2 GROUP BY s.name,  r.service_id", reportTable, serviceTable)
	rows, err := r.db.Query(query, year, month)
	if err != nil {
		return err
	}

	err = sqltocsv.WriteFile("report.csv", rows)
	return err
}

func (r *InfoPostgres) GiveName(serv balance.Report) error {
	query := fmt.Sprintf("INSERT INTO %s  (service_id, name) VALUES ($1, $2)", serviceTable)
	_, err := r.db.Exec(query, serv.ServiceId, serv.Name)
	return err
}
