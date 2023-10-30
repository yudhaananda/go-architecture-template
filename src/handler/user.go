package handler

import (
	"net/http"
	"strconv"
	"template/src/filter"
	"template/src/formatter"
	"template/src/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags User
// @Security ApiKeyAuth
// @Param paging query filter.Paging[filter.UserFilter] false "paging"
// @Param filter query filter.UserFilter false "filter"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user/ [GET]
func (h *handler) GetUser(ctx *gin.Context) {
	var filter filter.Paging[filter.UserFilter]
	filter.SetDefault()

	if err := h.BindParams(ctx, &filter); err != nil {
		response := models.APIResponse("Get User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, count, err := h.service.User.Get(ctx, filter)
	if err != nil {
		response := models.APIResponse("Get User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	paginatedItems := formatter.PaginatedItems{}
	paginatedItems.Format(filter.Page, float64(len(users)), float64(count), float64(filter.Take), users)

	response := models.APIResponse("Get User Success", http.StatusOK, "Success", paginatedItems, nil)
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
// @Success 200 {object} models.Response
// @Router /user/ [POST]
func (h *handler) CreateUser(ctx *gin.Context) {
	var input models.Query[models.UserInput]

	if err := ctx.ShouldBindJSON(&input.Model); err != nil {
		response := models.APIResponse("Create User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.User.Create(ctx, input); err != nil {
		response := models.APIResponse("Create User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Create User Success", http.StatusOK, "Success", nil, nil)
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
// @Success 200 {object} models.Response
// @Router /user/{id} [PUT]
func (h *handler) UpdateUser(ctx *gin.Context) {
	var input models.Query[models.UserInput]

	if err := ctx.ShouldBindJSON(&input.Model); err != nil {
		response := models.APIResponse("Update User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := models.APIResponse("Update User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.User.Update(ctx, input, id); err != nil {
		response := models.APIResponse("Update User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Update User Success", http.StatusOK, "Success", nil, nil)
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
// @Success 200 {object} models.Response
// @Router /user/{id} [DELETE]
func (h *handler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := models.APIResponse("Delete User Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.User.Delete(ctx, id); err != nil {
		response := models.APIResponse("Delete User Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Delete User Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}
