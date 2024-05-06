package models

import (
	"time"
)

type User struct {
	User_id    string     `db:"user_id" form:"user_id" json:"user_id" uri:"user_id" valid:"-"`
	First_name string     `db:"first_name" form:"first_name" json:"first_name" valid:"-"`
	Last_name  string     `db:"last_name" form:"last_name" json:"last_name" valid:"-"`
	Email      string     `db:"email" form:"email" json:"email" valid:"required, email"`
	Phone      string     `db:"phone" form:"phone" json:"phone" valid:"required"`
	Password   string     `db:"password" form:"password" json:"password" valid:"required, stringlength(6|100)~password minimal 6 karakter"`
	Role       string     `db:"role" json:"role,omitempty" valid:"-"`
	Birth      string     `db:"birth" form:"birth" json:"birth" valid:"-"`
	Gender     string     `db:"gender" form:"gender" json:"gender" valid:"-"`
	Image      string     `db:"image" form:"image" json:"image" valid:"-"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type Users []User
