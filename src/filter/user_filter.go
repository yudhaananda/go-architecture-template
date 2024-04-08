package filter

import (
	"time"
)

type UserFilter struct {
	Id        int       `db:"id" json:"id" form:"id" type:"number" name:"ID"`
	Name      string    `db:"name" json:"name" form:"name" type:"text" name:"Name"`
	UserName  string    `db:"user_name" json:"userName" form:"user_name" type:"text" name:"Username"`
	Birthdate time.Time `db:"birthdate" json:"birthdate" form:"birthdate" type:"text" name:"Birthdate"`
	Age       int       `db:"age" json:"age" form:"age" type:"number" name:"Age"`
}
