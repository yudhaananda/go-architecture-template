package models

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"template/src/formatter"
	"time"
)

type User struct {
	Id        int64                                 `db:"id" json:"id" header:"Id"`
	UserName  string                                `db:"user_name" json:"userName" header:"Username"`
	Password  string                                `db:"password" json:"password" header:"Password"`
	Name      formatter.NullableDataType[string]    `db:"name" json:"name" header:"Name"`
	Birthdate formatter.NullableDataType[time.Time] `db:"birthdate" json:"birthdate" header:"Birthdate"`
	Age       formatter.NullableDataType[int64]     `db:"age" json:"age" header:"Age"`
	Status    int64                                 `db:"status" json:"status" header:"-"`
	CreatedAt formatter.NullableDataType[time.Time] `db:"created_at" json:"createdAt" header:"-"`
	CreatedBy formatter.NullableDataType[int64]     `db:"created_by" json:"createdBy" header:"-"`
	UpdatedAt formatter.NullableDataType[time.Time] `db:"updated_at" json:"updatedAt" header:"-"`
	UpdatedBy formatter.NullableDataType[int64]     `db:"updated_by" json:"updatedBy" header:"-"`
	DeletedAt formatter.NullableDataType[time.Time] `db:"deleted_at" json:"deletedAt" header:"-"`
	DeletedBy formatter.NullableDataType[int64]     `db:"deleted_by" json:"deletedBy" header:"-"`
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

func (m User) ToColumn() (result Column) {
	ref := reflect.ValueOf(m)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("header") == "-" {
			continue
		}
		member := template.HTML(fmt.Sprint(ref.Field(i).Interface()))
		if strings.Contains(ref.Field(i).Type().Name(), "NullableDataType") {
			switch s := ref.Field(i).Interface().(type) {
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
					member = template.HTML(s.Data.Format(time.DateOnly))
				} else {
					member = ""
				}
			case formatter.NullableDataType[bool]:
				if s.Valid {
					member = template.HTML(fmt.Sprint(s.Data))
				} else {
					member = ""
				}
			}

		}
		result.Row = append(result.Row, MemberStruct{
			Member: member,
		})
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
