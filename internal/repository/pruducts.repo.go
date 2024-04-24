package repository

import (
	"coffeeshop-api-golang/internal/models"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type RepoProducts struct {
	*sqlx.DB
}

func NewPruduct(db *sqlx.DB) *RepoProducts {
	return &RepoProducts{db}
}

func (r *RepoProducts) GetBy(data *models.Product, page int, limit int) ([]models.Product, error) {
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
				image       string
				description string
				createdAt   *time.Time
				updatedAt   *time.Time
			)
			err := rows.Scan(&product_id, &name, &category, &price, &discount, &image, &description, &createdAt, &updatedAt)
			if err != nil {
				log.Println(err)
			}
			product := models.Product{Product_id: product_id, Name: name, Category: category, Price: price, Discount: discount, Image: image, Description: description, CreatedAt: createdAt, UpdatedAt: updatedAt}
			products = append(products, product)
		}
	}

	return products, nil

}

func (r *RepoProducts) CreateProduct(data *models.Product) (string, error) {
	q := `INSERT INTO public.products(
		name,
		category,
		price,
		discount,
		image,
		description)
	VALUES(
		:name,
		:category,
		:price,
		:discount,
		:image,
		:description
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data product created", nil

}

func (r *RepoProducts) UpdateProd(data *models.Product, product_id string) (string, error) {
	q := `UPDATE products SET
			name = $1,
			category = $2,
			price = $3,
			discount = $4,
			image = $5,
			description = $6,
			update_at = now()
			WHERE product_id::text = $7`

	_, err := r.Exec(q, data.Name, data.Category, data.Price, data.Discount, data.Image, data.Description, product_id)
	if err != nil {
		return "", err
	}

	return "1 data product updated", nil

}

func (r *RepoProducts) DeleteProd(data *models.Product) (string, error) {
	q := `delete from products where product_id::text = $1`

	_, err := r.Exec(q, data.Product_id)
	if err != nil {
		return "", err
	}

	return "1 data product deleted", nil

}
