package models

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"template/src/formatter"
	"time"
)

type HTMX[T comparable] struct {
	Model T
}

func (m HTMX[T]) GenerateHTML(html string) (result HTMXResult) {
	ref := reflect.ValueOf(m.Model)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if ref.Field(i).CanConvert(reflect.ValueOf(time.Time{}).Type()) {
			result.DateJQuery = append(result.DateJQuery, DateJQuery{Value: tpe.Field(i).Tag.Get("form")})
		}
		form := tpe.Field(i).Tag.Get("form")
		name := tpe.Field(i).Tag.Get("name")
		memberType := tpe.Field(i).Tag.Get("type")
		result.Members = append(result.Members, MemberStruct{
			Member: template.HTML(fmt.Sprintf(html, memberType, form, form, form, name)),
		})
	}
	return
}

func (m HTMX[T]) ToColumn(name string) (result Column) {
	ref := reflect.ValueOf(m.Model)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("header") == "-" {
			continue
		}

		if tpe.Field(i).Tag.Get("header") == "Id" {
			result.Id = template.HTML(fmt.Sprint(ref.Field(i).Interface()))
		}

		member := processValue(ref.Field(i).Type().Name(), time.DateOnly, ref.Field(i).Interface())

		result.Row = append(result.Row, MemberStruct{
			Member: member,
		})
	}
	result.Name = template.HTML(name)
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

func (m HTMX[T]) ToHeader() (result HTMXGet) {
	ref := reflect.ValueOf(m.Model)
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

func (m HTMX[T]) ToModalMember() (result []ModalMember) {
	ref := reflect.ValueOf(m.Model)
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
				Value:       processValue(ref.Field(i).Type().Name(), "2006-01-02T15:04:05Z", ref.Field(i).Interface()),
				Placeholder: template.HTML(tpe.Field(i).Tag.Get("header")),
			})
	}
	return
}

type HTMXResult struct {
	Members    []MemberStruct
	DateJQuery []DateJQuery
}

type MemberStruct struct {
	Member template.HTML
}

type HTMXGet struct {
	Header       []MemberStruct
	Column       []Column
	SideBar      []SideBar
	Link         template.HTML
	SectionName  template.HTML
	DateJQuery   []DateJQuery
	Filter       []HTMXFilter
	Pagination   []HTMXPagination
	IsFirst      bool
	IsLast       bool
	PreviousPage template.HTML
	NextPage     template.HTML
	LastPage     template.HTML
	Take         template.HTML
	QueryPage    template.HTML
	QueryTake    template.HTML
}

type DateJQuery struct {
	Value string
}

type HTMXPagination struct {
	Active    template.HTML
	Link      template.HTML
	Page      template.HTML
	QueryPage template.HTML
}

type HTMXFilter struct {
	Type  template.HTML
	Id    template.HTML
	Label template.HTML
	Value template.HTML
}

type SideBar struct {
	Active template.HTML
	Name   template.HTML
	Link   template.HTML
}

type Column struct {
	Row  []MemberStruct
	Id   template.HTML
	Name template.HTML
}

type Modal struct {
	Name    template.HTML
	Link    template.HTML
	Id      template.HTML
	Method  template.HTML
	Members []ModalMember
}

type ModalMember struct {
	Type        template.HTML
	Id          template.HTML
	Name        template.HTML
	Value       template.HTML
	Placeholder template.HTML
}
