package repos

import (
	"database/sql"

	"github.com/rizkazn/gorestfull/models"
)

type histRepo struct {
	db *sql.DB
}

func History(dbms *sql.DB) *histRepo {
	return &histRepo{dbms}
}

func (r *histRepo) FindAll() (*models.Histories, error) {
	query := `SELECT * FROM public.history`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Histories
	var histories models.History

	for rows.Next() {
		err := rows.Scan(&histories.Id, &histories.No_invoice, &histories.Cashier, &histories.Date, &histories.Orders, &histories.Amount)

		if err != nil {
			return nil, err
		}

		data = append(data, histories)

	}

	return &data, nil
}

func (r *histRepo) Save(history *models.History) error {
	query := `INSERT INTO public.history("id", "no_invoice", "cashier", "date", "orders", "amount") VALUES($1, $2, $3, $4, $5, $6)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(history.Id, history.No_invoice, history.Cashier, history.Date, history.Orders, history.Amount)

	if err != nil {
		return err
	}

	return nil
}

func (r *histRepo) Update(history *models.History, id string) error {
	query := `UPDATE public.history SET no_invoice=$1, cashier=$2, orders=$3, amount=$4 WHERE Id=$5`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(history.No_invoice, history.Cashier, history.Orders, history.Amount, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *histRepo) Remove(id string) error {
	query := `DELETE FROM public.history WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
