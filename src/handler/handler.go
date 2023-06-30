package handler

import (
	"net/http"
	"template/docs/swagger"
	"template/src/middleware"
	"template/src/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler interface {
	Run()
}

type handler struct {
	service    *services.Services
	middleware middleware.Interface
}

type InitParam struct {
	Service    *services.Services
	Middleware middleware.Interface
}

func Init(params InitParam) Handler {
	handler := &handler{
		service:    params.Service,
		middleware: params.Middleware,
	}
	return handler
}

func (h *handler) Run() {
	if err := h.register().Run(); err != nil {
		panic(err)
	}
}

func (h *handler) register() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	}))
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api/v1")
	swagger.SwaggerInfo.BasePath = "/api/v1"

	// API Route
	api.POST("/login", h.Login)
	api.POST("/register", h.Register)
	userApi := api.Group("/user").Use(h.middleware.AuthMiddleware)
	{
		userApi.GET("/", h.GetUser)
		userApi.POST("/", h.CreateUser)
		userApi.PUT("/:id", h.UpdateUser)
		userApi.DELETE("/:id", h.DeleteUser)
	}

	return router
}

func (h *handler) BindParams(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindWith(obj, binding.Query)
	if err != nil {
		return err
	}

	err = ctx.ShouldBindUri(obj)
	if err != nil {
		return err
	}

	return nil
}
