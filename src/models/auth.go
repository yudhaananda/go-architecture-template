package models

import (
	"time"
)

const (
	UserKey = "currentUser"
)

type Login struct {
	UserName string `json:"username" form:"username" name:"Username" type:"text"`
	Password string `json:"password" form:"password" name:"Password" type:"password"`
}

type Register struct {
	Name            string    `db:"name" json:"name" form:"name" name:"Name" type:"text"`
	UserName        string    `db:"user_name" json:"username" form:"username" name:"Username" type:"text"`
	Password        string    `db:"password" json:"password" form:"password" name:"Password" type:"password"`
	ConfirmPassword string    `db:"password" json:"confirmPassword" form:"confirmPassword" name:"Confirm Password" type:"password"`
	Birthdate       time.Time `db:"birthdate" json:"birthdate" form:"birthdate" name:"Birthdate" type:"text"`
	Age             int       `db:"age" json:"age" form:"age" name:"Age" type:"number"`
}

func (m *Register) ToUserInput() UserInput {
	return UserInput{
		Name:      m.Name,
		UserName:  m.UserName,
		Password:  m.Password,
		Birthdate: m.Birthdate,
		Age:       int64(m.Age),
	}
}
