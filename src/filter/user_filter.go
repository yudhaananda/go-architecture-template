package filter

import (
	"fmt"
	"html/template"
	"reflect"
	"template/src/models"
	"time"
)

type UserFilter struct {
	Id        int       `db:"id" json:"id" form:"id" type:"number" name:"ID"`
	Name      string    `db:"name" json:"name" form:"name" type:"text" name:"Name"`
	UserName  string    `db:"user_name" json:"userName" form:"user_name" type:"text" name:"Username"`
	Birthdate time.Time `db:"birthdate" json:"birthdate" form:"birthdate" type:"text" name:"Birthdate"`
	Age       int       `db:"age" json:"age" form:"age" type:"number" name:"Age"`
}

func (m UserFilter) ToHTMXFilter() (result []models.HTMXFilter) {
	ref := reflect.ValueOf(m)
	tpe := ref.Type()
	for i := 0; i < tpe.NumField(); i++ {
		value := template.HTML(fmt.Sprint(ref.Field(i).Interface()))
		if isEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			value = ""
		}
		result = append(result, models.HTMXFilter{
			Type:  template.HTML(tpe.Field(i).Tag.Get("type")),
			Id:    template.HTML(tpe.Field(i).Tag.Get("form")),
			Label: template.HTML(tpe.Field(i).Tag.Get("name")),
			Value: value,
		})
	}
	return
}
