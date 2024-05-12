package repository

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type RepoUsersIF interface {
	GetByEmail(email string) (*config.Result, error)
	GetAllUser() (*config.Result, error)
	GetAuthData(email string) (*models.User, error)
	CreateUser(data *models.User) (*config.Result, error)
	Update(data *models.User, user_id string) (*config.Result, error)
	Delete(data *models.User) (*config.Result, error)
}

type RepoUsers struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db}
}

func (r *RepoUsers) GetByEmail(email string) (*config.Result, error) {
	q := `select * from users where email = $1`

	// var user = struct {
	// 	User_id  string `json:"user_id" db:"user_id"`
	// 	Password string `json:"password" db:"password"`
	// }{}

	var user models.User

	err := r.Get(&user, r.Rebind(q), email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("email not found")
		}
		return nil, err
	}

	return &config.Result{Data: user}, nil

}

func (r *RepoUsers) GetAllUser() (*config.Result, error) {
	var data models.Users
	q := `SELECT email, "role", created_at, updated_at FROM users ORDER BY created_at DESC`

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &config.Result{Data: data}, nil
}

func (r *RepoUsers) GetAuthData(email string) (*models.User, error) {
	var result models.User
	q := `SELECT user_id, email, role, password FROM public.users WHERE email = ?`

	if err := r.Get(&result, r.Rebind(q), email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("email not found")
		}

		return nil, err
	}

	return &result, nil
}

func (r *RepoUsers) CreateUser(data *models.User) (*config.Result, error) {
	q := `INSERT INTO public.users(
		email,
		phone,
		password,
		role)
	VALUES(
		:email,
		:phone,
		:password,
		:role
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 data user created"}, nil

}

func (r *RepoUsers) Update(data *models.User, user_id string) (*config.Result, error) {
	q := `UPDATE users SET
			first_name = $1,
			last_name = $2,
			email = $3,
			phone = $4,
			password = $5,
			birth = $6,
			gender = $7,
			image = $8,
			updated_at = now()
			WHERE user_id::text = $9`

	_, err := r.Exec(q, data.First_name, data.Last_name, data.Email, data.Phone, data.Password, data.Birth, data.Gender, data.Image, user_id)
	if err != nil {
		return nil, err
	}

	return &config.Result{}, nil

}

func (r *RepoUsers) Delete(data *models.User) (*config.Result, error) {
	q := `delete from users where user_id::text = $1`

	_, err := r.Exec(q, data.User_id)
	if err != nil {
		return nil, err
	}

	return &config.Result{}, nil

}
