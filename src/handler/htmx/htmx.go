package htmx

import (
	"net/http"
	"os"
	"template/src/middleware"
	"template/src/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Htmx interface {
	RegisterPath(router *gin.Engine) *gin.Engine
}

type htmx struct {
	service    *services.Services
	middleware middleware.Interface
}

type InitParam struct {
	Service    *services.Services
	Middleware middleware.Interface
}

func Init(params InitParam) Htmx {
	htmx := &htmx{
		service:    params.Service,
		middleware: params.Middleware,
	}
	return htmx
}

func (h *htmx) RegisterPath(router *gin.Engine) *gin.Engine {
	login := router.Group("/login")
	{
		login.GET("", h.Login)
		login.POST("", h.LoginValidate)
	}
	register := router.Group("/register")
	{
		register.GET("", h.Register)
		register.POST("", h.RegisterValidate)
	}
	dashboard := router.Group("")
	{
		dashboard.GET("", h.GetDashboard)
		dashboard.GET("dashboard/content", h.middleware.AuthMiddleware, h.DashboardContent)
		dashboard.GET("/dashboard.css", func(ctx *gin.Context) {
			css, err := os.ReadFile("./src/view/index.css")
			if err != nil {
				ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
			}
			ctx.Data(http.StatusOK, "text/css", css)
		})
	}
	user := router.Group("/user")
	{
		user.GET("", h.GetUser)
		user.GET("/content", h.middleware.AuthMiddleware, h.UserContent)
		user.DELETE("/:id", h.middleware.AuthMiddleware, h.DeleteUser)
		user.GET("/edit-modal/:id", h.ModalEditUser)
		user.GET("/create-modal", h.ModalCreateUser)
		user.PUT("/:id", h.middleware.AuthMiddleware, h.EditUser)
		user.POST("", h.middleware.AuthMiddleware, h.CreateUser)
	}

	return router
}

func (h *htmx) BindParams(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindWith(obj, binding.Query)
	if err != nil {
		return err
	}

	err = ctx.ShouldBindWith(obj, binding.Form)
	if err != nil {
		return err
	}

	err = ctx.ShouldBindUri(obj)
	if err != nil {
		return err
	}
	return nil
}
