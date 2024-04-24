package models

import (
	"time"
)

type Favorite struct {
	Favorite_id string     `db:"favorite_id" form:"favorite_id" json:"favorite_id" uri:"favorite_id"`
	User_id     string     `db:"user_id" form:"user_id" json:"user_id" uri:"user_id"`
	Product_id  string     `db:"product_id" form:"product_id" json:"product_id"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

type FavoriteUser struct {
	Product_id  string  `db:"product_id" form:"product_id" json:"product_id" uri:"product_id"`
	Name        string  `db:"name" form:"name" json:"name"`
	Category    string  `db:"category" form:"category" json:"category"`
	Price       int     `db:"price" form:"price" json:"price"`
	Discount    float64 `db:"discount" form:"discount" json:"discount"`
	Image       string  `db:"image" form:"image" json:"image"`
	Description string  `db:"description" form:"description" json:"description"`
}
