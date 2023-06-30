package filter

import "time"

type UserFilter struct {
	Id        int       `db:"id" json:"id" form:"id"`
	Name      string    `db:"name" json:"name" form:"name"`
	UserName  string    `db:"user_name" json:"userName" form:"user_name"`
	Birthdate time.Time `db:"birthdate" json:"birthdate" form:"birthdate"`
	Age       int       `db:"age" json:"age" form:"age"`
}
