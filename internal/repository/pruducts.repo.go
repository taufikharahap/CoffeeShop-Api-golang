package repository

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/jmoiron/sqlx"
)

type RepoProducts struct {
	*sqlx.DB
}

func NewPruduct(db *sqlx.DB) *RepoProducts {
	return &RepoProducts{db}
}

func (r *RepoProducts) GetProd(data *models.Product, page int, limit int) ([]models.Product, error) {
	offset := (page - 1) * limit
	q := `select * from products order by name asc limit $1 offset $2`

	rows, err := r.Queryx(q, limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var products []models.Product
	if rows != nil {
		for rows.Next() {
			var (
				product_id  string
				name        string
				category    string
				price       int
				discount    float64
				image_url   string
				description *string
				created_at  *time.Time
				updated_at  *time.Time
			)
			err := rows.Scan(&product_id, &name, &category, &price, &discount, &image_url, &description, &created_at, &updated_at)
			if err != nil {
				log.Println(err)
			}
			product := models.Product{Product_id: product_id, Name: name, Category: category, Price: price, Discount: discount, Image_url: image_url, Description: description, CreatedAt: created_at, UpdatedAt: updated_at}
			products = append(products, product)
		}
	}
	return products, nil

}

func (r *RepoProducts) GetProdBy(params models.Meta) (*config.Result, error) {
	var data models.Products
	var metas config.Metas
	var filterQuery string
	var metaQuery string
	var count int
	var args []interface{}
	var filter []interface{}

	if params.Name != "" {
		filterQuery = "AND name = ?"
		args = append(args, params.Name)
		filter = append(filter, params.Name)
	}

	offset := (params.Page - 1) * params.Limit
	metaQuery = "LIMIT ? OFFSET ? "
	args = append(args, params.Limit, offset)

	m := fmt.Sprintf(`SELECT COUNT(product_id) as count FROM public.products WHERE true %s`, filterQuery)
	err := r.Get(&count, r.Rebind(m), filter...)
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf(`SELECT * FROM public.products WHERE true %s %s`, filterQuery, metaQuery)

	err = r.Select(&data, r.Rebind(q), args...)
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

	return &config.Result{Data: data, Meta: metas}, nil
}

func (r *RepoProducts) CreateProd(data *models.Product) (*config.Result, error) {
	q := `INSERT INTO public.products(
		name,
		category,
		price,
		discount,
		image_url,
		description)
	VALUES(
		:name,
		:category,
		:price,
		:discount,
		:image_url,
		:description
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	return &config.Result{}, nil

}

func (r *RepoProducts) UpdateProd(data *models.Product, product_id string) (*config.Result, error) {
	q := `UPDATE products SET
			name = $1,
			category = $2,
			price = $3,
			discount = $4,
			image_url = $5,
			description = $6,
			updated_at = now()
			WHERE product_id::text = $7`

	_, err := r.Exec(q, data.Name, data.Category, data.Price, data.Discount, data.Image_url, data.Description, product_id)
	if err != nil {
		return nil, err
	}

	return &config.Result{}, nil

}

func (r *RepoProducts) DeleteProd(data *models.Product) (*config.Result, error) {
	q := `delete from products where product_id::text = $1`

	_, err := r.Exec(q, data.Product_id)
	if err != nil {
		return nil, err
	}

	return &config.Result{}, nil

}
