package htmx

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	router.GET("/login", h.Login)
	router.POST("/login", h.LoginValidate)
	router.GET("/register", h.Register)
	router.POST("/register", h.RegisterValidate)
	router.GET("/", h.GetDashboard)
	router.GET("/dashboard-content", h.middleware.AuthMiddleware, h.DashboardContent)
	router.GET("/user", h.GetUser)
	router.GET("/user-content", h.middleware.AuthMiddleware, h.UserContent)
	router.GET("/dashboard.css", func(ctx *gin.Context) {
		css, err := os.ReadFile(h.Path() + "view/index.css")
		if err != nil {
			ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		}
		ctx.Data(http.StatusOK, "text/css", css)
	})

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

func (h *htmx) Path() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	dirnameArr := strings.Split(dirname, "/")

	path := ""
	for i := 0; i < len(dirnameArr)-2; i++ {
		path += dirnameArr[i] + "/"
	}
	return path
}
