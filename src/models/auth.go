package models

import (
	"strconv"
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
	Name            string `db:"name" json:"name" form:"name" name:"Name" type:"text"`
	UserName        string `db:"user_name" json:"username" form:"username" name:"Username" type:"text"`
	Password        string `db:"password" json:"password" form:"password" name:"Password" type:"password"`
	ConfirmPassword string `db:"password" json:"confirmPassword" form:"confirmPassword" name:"Confirm Password" type:"password"`
	Birthdate       string `db:"birthdate" json:"birthdate" form:"birthdate" name:"Birthdate" type:"text"`
	Age             string `db:"age" json:"age" form:"age" name:"Age" type:"number"`
}

func (m *Register) ToUserInput() UserInput {
	date, _ := time.Parse("01/02/2006", m.Birthdate)
	age, _ := strconv.ParseInt(m.Age, 10, 64)
	return UserInput{
		Name:      m.Name,
		UserName:  m.UserName,
		Password:  m.Password,
		Birthdate: date,
		Age:       age,
	}
}
