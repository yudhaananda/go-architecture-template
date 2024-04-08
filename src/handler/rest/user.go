package rest

import (
	"net/http"
	"strconv"
	"template/src/filter"
	"template/src/models"

	"github.com/gin-gonic/gin"
	"github.com/yudhaananda/go-common/formatter"
	"github.com/yudhaananda/go-common/paging"
	"github.com/yudhaananda/go-common/response"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags User
// @Security ApiKeyAuth
// @Param paging query paging.Paging[filter.UserFilter] false "paging"
// @Param filter query filter.UserFilter false "filter"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /user [GET]
func (h *rest) GetUser(ctx *gin.Context) {
	var filter paging.Paging[filter.UserFilter]
	filter.SetDefault()

	if err := h.BindParams(ctx, &filter); err != nil {
		response := response.APIResponse("Get User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, count, err := h.service.User.Get(ctx, filter)
	if err != nil {
		response := response.APIResponse("Get User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	paginatedItems := formatter.PaginatedItems{}
	paginatedItems.Format(filter.Page, float64(len(users)), float64(count), float64(filter.Take), users)

	response := response.APIResponse("Get User Success", http.StatusOK, "Success", paginatedItems, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags User
// @Security ApiKeyAuth
// @Param models body models.UserInput true "models"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /user [POST]
func (h *rest) CreateUser(ctx *gin.Context) {
	var input models.UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := response.APIResponse("Create User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.User.Create(ctx, input); err != nil {
		response := response.APIResponse("Create User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.APIResponse("Create User Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags User
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Param models body models.UserInput true "models"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /user/{id} [PUT]
func (h *rest) UpdateUser(ctx *gin.Context) {
	var input models.UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := response.APIResponse("Update User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := response.APIResponse("Update User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.User.Update(ctx, input, id); err != nil {
		response := response.APIResponse("Update User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.APIResponse("Update User Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags User
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /user/{id} [DELETE]
func (h *rest) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := response.APIResponse("Delete User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.User.Delete(ctx, id); err != nil {
		response := response.APIResponse("Delete User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.APIResponse("Delete User Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}
