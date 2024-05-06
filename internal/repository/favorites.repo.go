package repository

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"
	"fmt"
	"log"
	"math"

	"github.com/jmoiron/sqlx"
)

type RepoFavorites struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorites {
	return &RepoFavorites{db}
}

func (r *RepoFavorites) GetByUserId(data *models.Favorite, page int, limit int) ([]models.FavoriteUser, error) {
	offset := (page - 1) * limit

	q := `select p.product_id, p.name, p.category, p.price, p.discount, p.image_url, p.description
			from public.favorites f
			join public.users u on u.user_id::text = f.user_id::text
			join public.products p on p.product_id::text = f.product_id::text
			where u.user_id::text = $1
			limit $2 offset $3`

	rows, err := r.Queryx(q, data.User_id, limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var favoritesUser []models.FavoriteUser
	if rows != nil {
		for rows.Next() {
			var (
				product_id  string
				name        string
				category    string
				price       int
				discount    float64
				image_url   string
				description string
			)
			err := rows.Scan(&product_id, &name, &category, &price, &discount, &image_url, &description)
			if err != nil {
				log.Println(err)
			}
			favorite := models.FavoriteUser{Product_id: product_id, Name: name, Category: category, Price: price, Discount: discount, Image_url: image_url, Description: &description}
			favoritesUser = append(favoritesUser, favorite)
		}
	}

	return favoritesUser, nil

}

func (r *RepoFavorites) GetUserIdBy(params models.Meta, userId string) (*config.Result, error) {
	var result models.FavoriteUsers
	var metas config.Metas
	var filterQuery string
	var metaQuery string
	var count int
	var args []interface{}
	var filter []interface{}

	if params.Name != "" {
		filterQuery = "AND p.name = ?"
		args = append(args, params.Name)
		filter = append(filter, params.Name)
	}

	offset := (params.Page - 1) * params.Limit
	metaQuery = "LIMIT ? OFFSET ? "
	args = append(args, params.Limit, offset)

	m := fmt.Sprintf(`select COUNT(p.product_id) as count 
						from public.favorites f join public.users u on u.user_id = f.user_id join public.products p on p.product_id = f.product_id 
						where true and u.user_id = %s %s`, userId, filterQuery)
	err := r.Get(&count, r.Rebind(m), filter...)
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf(`select p.product_id, p.name, p.category, p.price, p.discount, p.image_url, p.description 
	from public.favorites f join public.users u on u.user_id = f.user_id 
	join public.products p on p.product_id = f.product_id where true and u.user_id = %s %s %s`, userId, filterQuery, metaQuery)

	err = r.Select(&result, r.Rebind(q), args...)
	if err != nil {
		return nil, err
	}

	check := math.Ceil(float64(count) / float64(params.Limit))
	metas.Total = count
	if count > 0 && params.Page != int(check) {
		metas.Next = params.Page + 1
	}

	if params.Page != 1 {
		metas.Prev = params.Page - 1
	}

	return &config.Result{Data: result, Meta: metas}, nil
}

func (r *RepoFavorites) CreateFavorite(data *models.Favorite) (string, error) {
	q := `INSERT INTO public.favorites(
		user_id,
		product_id
		)
	VALUES(
		:user_id,
		:product_id
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data favorite created", nil

}

func (r *RepoFavorites) UpdateFav(data *models.Favorite, favorite_id string) (string, error) {
	q := `UPDATE favorites SET
			user_id = $1,
			product_id = $2,
			update_at = now()
			WHERE favorite_id::text = $3`

	_, err := r.Exec(q, data.User_id, data.Product_id, favorite_id)
	if err != nil {
		return "", err
	}

	return "1 data favorite updated", nil

}

func (r *RepoFavorites) DeleteFav(data *models.Favorite) (string, error) {
	q := `delete from favorites where favorite_id::text = $1`

	_, err := r.Exec(q, data.Favorite_id)
	if err != nil {
		return "", err
	}

	return "1 data favorite deleted", nil

}
