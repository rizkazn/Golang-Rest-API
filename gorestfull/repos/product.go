package repos

import (
	"database/sql"

	"github.com/rizkazn/gorestfull/models"
)

type initRepo struct {
	db *sql.DB
}

// constructor
func Products(dbms *sql.DB) *initRepo {
	return &initRepo{dbms}
}

func (r *initRepo) FindAll() (*models.Products, error) {
	query := `SELECT public.product.id, public.product.product_name, public.product.price, public.product.product_image, public.category.id, public.product.created_at, public.product.updated_at FROM public.product INNER JOIN public.category ON public.product.category_id = public.category.id`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var products models.Product

	for rows.Next() {
		err := rows.Scan(&products.Id, &products.Product_name, &products.Price, &products.Product_image, &products.Category_id, &products.CreatedAt, &products.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, products)

	}

	return &data, nil
}

func (r *initRepo) Save(product *models.Product) error {
	query := `INSERT INTO public.product("id", "product_name", "price", "product_image", "category_id", "created_at", "updated_at") VALUES($1, $2, $3, $4, $5, $6, $7)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(product.Id, product.Product_name, product.Price, product.Product_image, product.Category_id, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepo) Update(product *models.Product, id string) error {
	query := `UPDATE public.product SET product_name=$1, price=$2, category_id=$3, updated_at=$4 WHERE Id=$5`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(product.Product_name, product.Price, product.Category_id, product.UpdatedAt, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *initRepo) Remove(id string) error {
	query := `DELETE FROM public.product WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepo) SearchProductBytName(name string) (*models.Products, error) {
	// query := `SELECT public.product.id, public.product.product_name, public.product.price, public.product.product_image, public.category.id, public.product.created_at, public.product.updated_at FROM public.product INNER JOIN public.category ON public.product.category_id = public.category.id WHERE public.product.product_name ILIKE $1 ORDER BY created_at DESC`

	query := `SELECT * FROM public.product WHERE product_name ILIKE '%' || $1% || '%' ORDER BY category_id DESC`

	rows, err := r.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var products models.Product

	for rows.Next() {
		err := rows.Scan(&products.Id, &products.Product_name, &products.Price, &products.Product_image, &products.Category_id, &products.CreatedAt, &products.UpdatedAt)
		if err != nil {
			return nil, err
		}
		data = append(data, products)
	}

	return &data, nil
}

func (r *initRepo) OrderProdByName() (*models.Products, error) {
	query := `SELECT * FROM public.product ORDER BY public.product.product_name ASC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var products models.Product

	for rows.Next() {
		err := rows.Scan(&products.Id, &products.Product_name, &products.Price, &products.Product_image, &products.Category_id, &products.CreatedAt, &products.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, products)

	}

	return &data, nil
}

func (r *initRepo) OrderProdByCategory() (*models.Products, error) {
	query := `SELECT * FROM public.product ORDER BY public.product.category_id ASC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var products models.Product

	for rows.Next() {
		err := rows.Scan(&products.Id, &products.Product_name, &products.Price, &products.Product_image, &products.Category_id, &products.CreatedAt, &products.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, products)

	}

	return &data, nil
}

func (r *initRepo) OrderProdByNewest() (*models.Products, error) {
	query := `SELECT * FROM public.product ORDER BY public.product.updated_at DESC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var products models.Product

	for rows.Next() {
		err := rows.Scan(&products.Id, &products.Product_name, &products.Price, &products.Product_image, &products.Category_id, &products.CreatedAt, &products.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, products)

	}

	return &data, nil
}

func (r *initRepo) OrderProdByPrice() (*models.Products, error) {
	query := `SELECT * FROM public.product ORDER BY public.product.price DESC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var products models.Product

	for rows.Next() {
		err := rows.Scan(&products.Id, &products.Product_name, &products.Price, &products.Product_image, &products.Category_id, &products.CreatedAt, &products.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, products)

	}

	return &data, nil
}
