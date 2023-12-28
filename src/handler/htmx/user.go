package htmx

import (
	"html/template"
	"net/http"
	"strings"
	"template/src/filter"
	"template/src/models"

	"github.com/gin-gonic/gin"
)

const (
	User = "User"
)

func (h *htmx) GetUser(ctx *gin.Context) {
	name := make(map[string]string)
	name["Name"] = strings.ToLower(User)
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/middleware.html"))
	tmpl.Execute(ctx.Writer, name)
}

func (h *htmx) UserContent(ctx *gin.Context) {
	user := models.User{}
	htmxGet := user.ToHeader()
	htmxGet.SectionName = User
	for _, feature := range models.Features {
		temp := models.SideBar{
			Name: template.HTML(feature.Name),
			Link: template.HTML(feature.Link),
		}
		if feature.Name == User {
			temp.Active = template.HTML("active")
		}
		htmxGet.SideBar = append(htmxGet.SideBar, temp)
	}
	users, _, err := h.service.User.Get(ctx, filter.Paging[filter.UserFilter]{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	for _, user := range users {
		htmxGet.Column = append(htmxGet.Column, user.ToColumn())
	}
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/index.html"))
	tmpl.Execute(ctx.Writer, htmxGet)
}
