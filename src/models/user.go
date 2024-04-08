package models

import (
	"time"

	"github.com/yudhaananda/go-common/formatter"
)

type User struct {
	Id        formatter.Null[int64]     `db:"id" json:"id" header:"Id" type:"-"`
	UserName  formatter.Null[string]    `db:"user_name" json:"username" header:"Username" type:"text"`
	Password  formatter.Null[string]    `db:"password" json:"password" header:"Password" type:"password"`
	Name      formatter.Null[string]    `db:"name" json:"name" header:"Name" type:"text"`
	Birthdate formatter.Null[time.Time] `db:"birthdate" json:"birthdate" header:"Birthdate" type:"text"`
	Age       formatter.Null[int64]     `db:"age" json:"age" header:"Age" type:"number"`
	Status    formatter.Null[int64]     `db:"status" json:"status" header:"-" type:"-"`
	CreatedAt formatter.Null[time.Time] `db:"created_at" json:"createdAt" header:"-" type:"-"`
	CreatedBy formatter.Null[int64]     `db:"created_by" json:"createdBy" header:"-" type:"-"`
	UpdatedAt formatter.Null[time.Time] `db:"updated_at" json:"updatedAt" header:"-" type:"-"`
	UpdatedBy formatter.Null[int64]     `db:"updated_by" json:"updatedBy" header:"-" type:"-"`
	DeletedAt formatter.Null[time.Time] `db:"deleted_at" json:"deletedAt" header:"-" type:"-"`
	DeletedBy formatter.Null[int64]     `db:"deleted_by" json:"deletedBy" header:"-" type:"-"`
}

func (u *User) ToDto() *UserDto {
	return &UserDto{
		Id:        u.Id.Data,
		UserName:  u.UserName.Data,
		Password:  u.Password.Data,
		Name:      u.Name.Data,
		Birthdate: u.Birthdate.Data,
		Age:       u.Age.Data,
		Status:    u.Status.Data,
		CreatedAt: u.CreatedAt.Data,
		CreatedBy: u.CreatedBy.Data,
		UpdatedAt: u.UpdatedAt.Data,
		UpdatedBy: u.UpdatedBy.Data,
		DeletedAt: u.DeletedAt.Data,
		DeletedBy: u.DeletedBy.Data,
	}
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

type UserDto struct {
	Id        int64     `db:"id" json:"id" header:"Id" type:"-"`
	UserName  string    `db:"user_name" json:"username" header:"Username" type:"text"`
	Password  string    `db:"password" json:"password" header:"Password" type:"password"`
	Name      string    `db:"name" json:"name" header:"Name" type:"text"`
	Birthdate time.Time `db:"birthdate" json:"birthdate" header:"Birthdate" type:"text"`
	Age       int64     `db:"age" json:"age" header:"Age" type:"number"`
	Status    int64     `db:"status" json:"status" header:"-" type:"-"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" header:"-" type:"-"`
	CreatedBy int64     `db:"created_by" json:"createdBy" header:"-" type:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt" header:"-" type:"-"`
	UpdatedBy int64     `db:"updated_by" json:"updatedBy" header:"-" type:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"deletedAt" header:"-" type:"-"`
	DeletedBy int64     `db:"deleted_by" json:"deletedBy" header:"-" type:"-"`
}
