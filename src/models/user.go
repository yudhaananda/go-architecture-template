package models

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
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

type UserInputForHTMX struct {
	Name      string `db:"name" json:"name" form:"name"`
	UserName  string `db:"user_name" json:"username" form:"username"`
	Password  string `db:"password" json:"password" form:"password"`
	Birthdate string `db:"birthdate" json:"birthdate" form:"birthdate" example:"2022-06-21T10:32:29Z"`
	Age       string `db:"age" json:"age" form:"age"`
}

func (m UserInputForHTMX) ToUserInput() UserInput {
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

func (m User) ToColumn() (result Column) {
	ref := reflect.ValueOf(m)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("header") == "-" {
			continue
		}

		member := processValue(ref.Field(i).Type().Name(), time.DateOnly, ref.Field(i).Interface())

		result.Row = append(result.Row, MemberStruct{
			Member: member,
		})
	}
	result.Name = "user"
	result.Id = template.HTML(fmt.Sprint(m.Id))
	return
}

func processValue(name, timeFormat string, val interface{}) (member template.HTML) {
	if strings.Contains(name, "NullableDataType") {
		switch s := val.(type) {
		case formatter.NullableDataType[string]:
			if s.Valid {
				member = template.HTML(s.Data)
			} else {
				member = ""
			}
		case formatter.NullableDataType[int]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[int64]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[int32]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[int16]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[int8]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[float32]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[float64]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.NullableDataType[time.Time]:
			if s.Valid {
				member = template.HTML(s.Data.Format(timeFormat))
			} else {
				member = ""
			}
		case formatter.NullableDataType[bool]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		default:
			member = template.HTML(fmt.Sprint(val))
		}

	} else {
		member = template.HTML(fmt.Sprint(val))
	}

	return
}

func (m User) ToHeader() (result HTMXGet) {
	ref := reflect.ValueOf(m)
	tpe := ref.Type()
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("header") == "-" {
			continue
		}
		result.Header = append(result.Header, MemberStruct{
			Member: template.HTML(tpe.Field(i).Tag.Get("header")),
		})
	}
	return
}

func (m User) ToModalMember() (result []ModalMember) {
	ref := reflect.ValueOf(m)
	tpe := ref.Type()
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("type") == "-" {
			continue
		}
		result = append(result,
			ModalMember{
				Id:          template.HTML(tpe.Field(i).Tag.Get("json")),
				Type:        template.HTML(tpe.Field(i).Tag.Get("type")),
				Name:        template.HTML(tpe.Field(i).Tag.Get("json")),
				Value:       processValue(ref.Field(i).Type().Name(), "01/02/2006", ref.Field(i).Interface()),
				Placeholder: template.HTML(tpe.Field(i).Tag.Get("header")),
			})
	}
	return
}
