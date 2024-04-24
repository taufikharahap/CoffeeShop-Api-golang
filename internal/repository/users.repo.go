package repository

import (
	"coffeeshop-api-golang/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepoUsers struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db}
}

// func rowToStruct(rows *sqlx.Rows, dest interface{}) error {
// 	destv := reflect.ValueOf(dest).Elem()

// 	args := make([]interface{}, destv.Type().Elem().NumField())

// 	for rows.Next() {
// 		rowp := reflect.New(destv.Type().Elem())
// 		rowv := rowp.Elem()

// 		for i := 0; i < rowv.NumField(); i++ {
// 			args[i] = rowv.Field(i).Addr().Interface()
// 		}

// 		if err := rows.Scan(args...); err != nil {
// 			return err
// 		}

// 		destv.Set(reflect.Append(destv, rowv))
// 	}

// 	return nil
// }

func (r *RepoUsers) GetByEmail(data *models.User) (interface{}, error) {
	q := `select user_id, password from users where email = $1`

	var user = struct {
		User_id  string `json:"user_id" db:"user_id"`
		Password string `json:"password" db:"password"`
	}{}

	err := r.Get(&user, q, data.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil

}

func (r *RepoUsers) CreateUser(data *models.User) (string, error) {
	q := `INSERT INTO public.users(
		email,
		phone,
		password)
	VALUES(
		:email,
		:phone,
		:password
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data user created", nil

}

func (r *RepoUsers) Update(data *models.User, user_id string) (string, error) {
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
		return "", err
	}

	return "1 data user updated", nil

}

func (r *RepoUsers) Delete(data *models.User) (string, error) {
	q := `delete from users where user_id::text = $1`

	_, err := r.Exec(q, data.User_id)
	if err != nil {
		return "", err
	}

	return "1 data user deleted", nil

}
