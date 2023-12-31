package htmx

import (
	"html/template"
	"net/http"
	"template/src/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	LoginMember = `
	<div class="form-outline mb-4">
    	<input type="%s" id="%s" name="%s" class="form-control" />
    	<label class="form-label" for="%s">%s</label>
	</div>`
)

func (h *htmx) Login(ctx *gin.Context) {
	models := models.HTMX[models.Login]{}
	member := models.GenerateHTML(LoginMember)
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/login.html"))
	tmpl.Execute(ctx.Writer, member)
}

func (h *htmx) Register(ctx *gin.Context) {
	models := models.HTMX[models.Register]{}
	member := models.GenerateHTML(LoginMember)
	tmpl := template.Must(template.ParseFiles(h.Path() + "view/register.html"))
	tmpl.Execute(ctx.Writer, member)
}

func (h *htmx) LoginValidate(ctx *gin.Context) {
	var input models.Login

	err := ctx.ShouldBindWith(&input, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	_, token, err := h.service.Auth.Login(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, "Bearer "+token)
}

func (h *htmx) RegisterValidate(ctx *gin.Context) {
	var input models.Register

	err := ctx.ShouldBindWith(&input, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	if input.Password != input.ConfirmPassword {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	err = h.service.Auth.Register(ctx, models.Query[models.UserInput]{
		Model: input.ToUserInput(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
