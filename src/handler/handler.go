package handler

import (
	"template/src/handler/htmx"
	"template/src/handler/rest"
	"template/src/middleware"
	"template/src/services"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Run()
}

type handler struct {
	rest rest.Rest
	htmx htmx.Htmx
}

type InitParam struct {
	Service    *services.Services
	Middleware middleware.Interface
}

func Init(params InitParam) Handler {
	rest := &handler{
		rest: rest.Init(rest.InitParam{Service: params.Service, Middleware: params.Middleware}),
		htmx: htmx.Init(htmx.InitParam{Service: params.Service, Middleware: params.Middleware}),
	}
	return rest
}

func (h *handler) Run() {
	router := gin.Default()
	router = h.htmx.RegisterPath(router)
	router = h.rest.RegisterPath(router)
	if err := router.Run(); err != nil {
		panic(err)
	}
}
