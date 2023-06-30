package models

import (
	"template/src/formatter"
	"time"
)

type User struct {
	Id        int                           `db:"id" json:"id"`
	Name      string                        `db:"name" json:"name"`
	UserName  string                        `db:"user_name" json:"userName"`
	Password  string                        `db:"password" json:"password"`
	Birthdate time.Time                     `db:"birthdate" json:"birthdate"`
	Age       int                           `db:"age" json:"age"`
	Flag      formatter.DataType[float32]   `db:"flag" json:"flag"`
	CreatedAt formatter.DataType[time.Time] `db:"created_at" json:"createdAt"`
	CreatedBy formatter.DataType[string]    `db:"created_by" json:"createdBy"`
	UpdatedAt formatter.DataType[time.Time] `db:"updated_at" json:"updatedAt"`
	UpdatedBy formatter.DataType[string]    `db:"updated_by" json:"updatedBy"`
	DeletedAt formatter.DataType[time.Time] `db:"deleted_at" json:"deletedAt"`
	DeletedBy formatter.DataType[string]    `db:"deleted_by" json:"deletedBy"`
}

type UserInput struct {
	Name      string    `db:"name"`
	Password  string    `db:"password" json:"password"`
	Birthdate time.Time `db:"birthdate" example:"2022-06-21T10:32:29Z"`
	Age       int       `db:"age"`
	Flag      int       `db:"flag" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	CreatedBy string    `db:"created_by" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	UpdatedBy string    `db:"updated_by" json:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"-"`
	DeletedBy string    `db:"deleted_by" json:"-"`
}
