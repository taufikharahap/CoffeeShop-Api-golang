package models

import (
	"time"
)

type User struct {
	User_id    string     `db:"user_id" form:"user_id" json:"user_id" uri:"user_id"`
	First_name string     `db:"first_name" form:"first_name" json:"first_name"`
	Last_name  string     `db:"last_name" form:"last_name" json:"last_name"`
	Email      string     `db:"email" form:"email" json:"email"`
	Phone      string     `db:"phone" form:"phone" json:"phone"`
	Password   string     `db:"password" form:"password" json:"password"`
	Birth      string     `db:"birth" form:"birth" json:"birth"`
	Gender     string     `db:"gender" form:"gender" json:"gender"`
	Image      string     `db:"image" form:"image" json:"image"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}
