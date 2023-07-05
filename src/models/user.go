package models

import (
	"template/src/formatter"
	"time"
)

type User struct {
	Id        int                                   `db:"id" json:"id"`
	UserName  string                                `db:"user_name" json:"userName"`
	Password  string                                `db:"password" json:"password"`
	Name      formatter.NullableDataType[string]    `db:"name" json:"name"`
	Birthdate formatter.NullableDataType[time.Time] `db:"birthdate" json:"birthdate"`
	Age       formatter.NullableDataType[int]       `db:"age" json:"age"`
	Status    formatter.NullableDataType[float32]   `db:"status" json:"status"`
	CreatedAt formatter.NullableDataType[time.Time] `db:"created_at" json:"createdAt"`
	CreatedBy formatter.NullableDataType[int]       `db:"created_by" json:"createdBy"`
	UpdatedAt formatter.NullableDataType[time.Time] `db:"updated_at" json:"updatedAt"`
	UpdatedBy formatter.NullableDataType[int]       `db:"updated_by" json:"updatedBy"`
	DeletedAt formatter.NullableDataType[time.Time] `db:"deleted_at" json:"deletedAt"`
	DeletedBy formatter.NullableDataType[int]       `db:"deleted_by" json:"deletedBy"`
}

type UserInput struct {
	Name      string    `db:"name"`
	UserName  string    `db:"user_name" json:"userName"`
	Password  string    `db:"password" json:"password"`
	Birthdate time.Time `db:"birthdate" example:"2022-06-21T10:32:29Z"`
	Age       int       `db:"age"`
	Status    int       `db:"status" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	CreatedBy int       `db:"created_by" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	UpdatedBy int       `db:"updated_by" json:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"-"`
	DeletedBy int       `db:"deleted_by" json:"-"`
}
