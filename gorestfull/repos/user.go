package repos

import (
	"database/sql"

	"github.com/rizkazn/gorestfull/models"
)

type userRepo struct {
	db *sql.DB
}

func User(dbms *sql.DB) *userRepo {
	return &userRepo{dbms}
}

func (r *userRepo) FindAll() (*models.Users, error) {
	query := `SELECT * FROM public.user`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Users
	var users models.User

	for rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.Username, &users.Email, &users.Password, &users.Role, &users.CreatedAt, &users.UpdatedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, users)

	}

	return &data, nil
}

func (r *userRepo) Save(user *models.User) error {
	query := `INSERT INTO public.user("id","name", "username", "email", "password", "role", "created_at", "updated_at") VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(user.Id, user.Name, user.Username, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetPassword(username string) (string, error) {
	query := `SELECT "password" FROM public.user WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var password string

	for rows.Next() {
		err := rows.Scan(&password)

		if err != nil {
			return "", err
		}

	}

	return password, nil
}

func (r *userRepo) GetUsername(username string) (string, error) {
	query := `SELECT "username" FROM public.user WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var user_name string

	for rows.Next() {
		err := rows.Scan(&username)

		if err != nil {
			return "", err
		}

	}
	// fmt.Println(user_name)
	return user_name, nil
}

func (r *userRepo) GetEmail(username string) (string, error) {
	query := `SELECT "email" FROM public.user WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var email string

	for rows.Next() {
		err := rows.Scan(&email)

		if err != nil {
			return "", err
		}

	}
	// fmt.Println(email)
	return email, nil

}

func (r *userRepo) GetRole(username string) (string, error) {
	query := `SELECT "role" FROM public.user WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var role string

	for rows.Next() {
		err := rows.Scan(&role)

		if err != nil {
			return "", err
		}

	}
	// fmt.Println(email)
	return role, nil

}
