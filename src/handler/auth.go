package handler

import (
	"net/http"
	"template/src/formatter"
	"template/src/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags Auth
// @Param registerInput body models.UserInput true "registerInput"
// @Accept json
// @Produce json
// @Success 200 {object} formatter.AuthFormatter
// @Router /register [post]
func (h *handler) Register(ctx *gin.Context) {
	var input models.Query[models.UserInput]

	err := ctx.ShouldBindJSON(&input.Model)
	if err != nil {
		response := models.APIResponse("Register Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.Auth.Register(ctx, input)

	if err != nil {
		response := models.APIResponse("Register Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.APIResponse("Register Success", http.StatusOK, "Success", nil, nil)

	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags Auth
// @Param loginInput body models.Login true "loginInput"
// @Accept json
// @Produce json
// @Success 200 {object} formatter.AuthFormatter
// @Router /login [post]
func (h *handler) Login(ctx *gin.Context) {
	var input models.Login

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		response := models.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.service.Auth.Login(ctx, input)
	if err != nil {
		response := models.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.service.Jwt.GenerateToken(loggedinUser.Id, loggedinUser.UserName)

	if err != nil {
		errorMessage := err.Error()

		response := models.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", nil, errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := formatter.Auth[models.User]{}
	formatter.Format(loggedinUser, token)

	response := models.APIResponse("Logedin", http.StatusOK, "success", formatter, nil)

	ctx.JSON(http.StatusOK, response)
}
