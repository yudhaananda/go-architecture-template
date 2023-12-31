package models

import (
	"template/src/formatter"
	"time"
)

type User struct {
	Id        int64                                 `db:"id" json:"id" header:"Id" type:"-"`
	UserName  string                                `db:"user_name" json:"username" header:"Username" type:"text"`
	Password  string                                `db:"password" json:"password" header:"Password" type:"password"`
	Name      formatter.NullableDataType[string]    `db:"name" json:"name" header:"Name" type:"text"`
	Birthdate formatter.NullableDataType[time.Time] `db:"birthdate" json:"birthdate" header:"Birthdate" type:"text"`
	Age       formatter.NullableDataType[int64]     `db:"age" json:"age" header:"Age" type:"number"`
	Status    int64                                 `db:"status" json:"status" header:"-" type:"-"`
	CreatedAt formatter.NullableDataType[time.Time] `db:"created_at" json:"createdAt" header:"-" type:"-"`
	CreatedBy formatter.NullableDataType[int64]     `db:"created_by" json:"createdBy" header:"-" type:"-"`
	UpdatedAt formatter.NullableDataType[time.Time] `db:"updated_at" json:"updatedAt" header:"-" type:"-"`
	UpdatedBy formatter.NullableDataType[int64]     `db:"updated_by" json:"updatedBy" header:"-" type:"-"`
	DeletedAt formatter.NullableDataType[time.Time] `db:"deleted_at" json:"deletedAt" header:"-" type:"-"`
	DeletedBy formatter.NullableDataType[int64]     `db:"deleted_by" json:"deletedBy" header:"-" type:"-"`
}

type UserInput struct {
	Name      string    `db:"name" json:"name" form:"name"`
	UserName  string    `db:"user_name" json:"username" form:"username"`
	Password  string    `db:"password" json:"password" form:"password"`
	Birthdate time.Time `db:"birthdate" json:"birthdate" form:"birthdate" example:"2022-06-21T10:32:29Z"`
	Age       int64     `db:"age" json:"age" form:"age"`
	Status    int64     `db:"status" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	CreatedBy int64     `db:"created_by" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	UpdatedBy int64     `db:"updated_by" json:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"-"`
	DeletedBy int64     `db:"deleted_by" json:"-"`
}
