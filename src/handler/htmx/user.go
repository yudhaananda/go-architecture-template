package htmx

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"template/src/filter"
	"template/src/formatter"
	"template/src/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	User     = "User"
	UserLink = "user"
)

func (h *htmx) GetUser(ctx *gin.Context) {
	name := make(map[string]string)
	name["Name"] = UserLink
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/middleware.html"))
	tmpl.Execute(ctx.Writer, name)
}

func (h *htmx) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	err = h.service.User.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	h.UserContent(ctx)
}

func (h *htmx) ModalCreateUser(ctx *gin.Context) {
	user := models.User{}
	modal := models.Modal{
		Name:   template.HTML("Create " + User),
		Link:   template.HTML(UserLink),
		Method: "post",
	}
	modal.Members = user.ToModalMember()
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/modal.html"))
	tmpl.Execute(ctx.Writer, modal)
}

func (h *htmx) ModalEditUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	users, _, err := h.service.User.Get(ctx, filter.Paging[filter.UserFilter]{
		Filter: filter.UserFilter{
			Id: id,
		},
	})
	if err != nil || len(users) < 1 {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	modal := models.Modal{
		Name:   template.HTML("Edit " + User),
		Link:   template.HTML(UserLink),
		Id:     template.HTML(fmt.Sprint(id)),
		Method: "put",
	}
	modal.Members = users[0].ToModalMember()
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/modal.html"))
	tmpl.Execute(ctx.Writer, modal)
}

func (h *htmx) CreateUser(ctx *gin.Context) {
	var input models.UserInput

	err := ctx.ShouldBindWith(&input, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	err = h.service.User.Create(ctx, models.Query[models.UserInput]{
		Model: input,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (h *htmx) EditUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	var input models.UserInput

	err = ctx.ShouldBindWith(&input, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	err = h.service.User.Update(ctx, models.Query[models.UserInput]{
		Model: input,
	}, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (h *htmx) UserContent(ctx *gin.Context) {
	user := models.User{}
	var filter filter.Paging[filter.UserFilter]
	filter.SetDefault()
	if err := h.BindParams(ctx, &filter); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	htmxGet := user.ToHeader()
	htmxGet.SectionName = User
	htmxGet.Link = UserLink
	htmxGet.Filter = filter.Filter.ToHTMXFilter()
	for key, values := range ctx.Request.URL.Query() {
		if key == "page" {
			continue
		}
		for _, value := range values {
			htmxGet.QueryPage += template.HTML(key + "=" + value + "&")
			if key == "take" {
				continue
			}
			htmxGet.QueryTake += template.HTML(key + "=" + value + "&")
		}
	}
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

	users, count, err := h.service.User.Get(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	paginatedItems := formatter.PaginatedItems{}
	paginatedItems.Format(filter.Page, float64(len(users)), float64(count), float64(filter.Take), users)
	if paginatedItems.PageCount == paginatedItems.PageIndex {
		htmxGet.IsLast = true
	}
	if paginatedItems.PageIndex == 1 {
		htmxGet.IsFirst = true
	}
	htmxGet.Take = template.HTML(fmt.Sprint(filter.Take))
	htmxGet.LastPage = template.HTML(fmt.Sprint(paginatedItems.PageCount))

	pageCount := 0

	for i := paginatedItems.PageIndex - 2; i <= paginatedItems.PageCount; i++ {
		if i < 1 {
			continue
		}
		if pageCount > 4 {
			break
		}
		active := ""
		if i == paginatedItems.PageIndex {
			active = "active"
			htmxGet.NextPage = template.HTML(fmt.Sprint(i + 1))
			htmxGet.PreviousPage = template.HTML(fmt.Sprint(i - 1))
		}
		htmxGet.Pagination = append(htmxGet.Pagination, models.HTMXPagination{
			Active:    template.HTML(active),
			Link:      UserLink,
			Page:      template.HTML(fmt.Sprint(i)),
			QueryPage: htmxGet.QueryPage,
		})
		pageCount++
	}

	for _, user := range users {
		htmxGet.Column = append(htmxGet.Column, user.ToColumn())
	}
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/index.html"))
	tmpl.Execute(ctx.Writer, htmxGet)
}
