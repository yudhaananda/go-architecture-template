package htmx

import (
	"html/template"
	"net/http"
	"template/src/filter"
	"template/src/models"

	"github.com/gin-gonic/gin"
)

func (h *htmx) Index(ctx *gin.Context) {
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/index.html"))
	tmpl.Execute(ctx.Writer, nil)
}

func (h *htmx) GetUser(ctx *gin.Context) {
	user := models.User{}
	htmxGet := user.ToHeader()
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
