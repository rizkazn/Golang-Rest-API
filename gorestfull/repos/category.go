package repos

import (
	"database/sql"

	"github.com/rizkazn/gorestfull/models"
)

type catRepo struct {
	db *sql.DB
}

func Category(dbms *sql.DB) *catRepo {
	return &catRepo{dbms}
}

func (r *catRepo) FindAll() (*models.Categories, error) {
	query := `SELECT * FROM public.category`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Categories
	var categories models.Category

	for rows.Next() {
		err := rows.Scan(&categories.Id, &categories.Category_name, &categories.Category_image, &categories.CreatedAt, &categories.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, categories)

	}

	return &data, nil
}

func (r *catRepo) Save(category *models.Category) error {
	query := `INSERT INTO public.category("id", "category_name", "category_image", "created_at", "updated_at") VALUES($1, $2, $3, $4, $5)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(category.Id, category.Category_name, category.Category_image, category.CreatedAt, category.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *catRepo) Update(category *models.Category, id string) error {
	query := `UPDATE public.category SET category_name=$1, updated_at=$2 WHERE Id=$3`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(category.Category_name, category.UpdatedAt, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *catRepo) Remove(id string) error {
	query := `DELETE FROM public.category WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
