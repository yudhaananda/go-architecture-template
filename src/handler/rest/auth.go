package rest

import (
	"net/http"
	"template/src/models"

	"github.com/gin-gonic/gin"
	"github.com/yudhaananda/go-common/formatter"
	"github.com/yudhaananda/go-common/response"
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
// @Success 200 {object} response.Response
// @Router /register [post]
func (h *rest) Register(ctx *gin.Context) {
	var input models.UserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := response.APIResponse("Register Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.Auth.Register(ctx, input)

	if err != nil {
		response := response.APIResponse("Register Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.APIResponse("Register Success", http.StatusOK, "Success", nil, nil)

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
// @Success 200 {object} response.Response
// @Router /login [post]
func (h *rest) Login(ctx *gin.Context) {
	var input models.Login

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := response.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, token, err := h.service.Auth.Login(ctx, input)
	if err != nil {
		errorMessage := err.Error()

		response := response.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Failed", nil, errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	auth := formatter.Auth{}
	auth.AuthFormat(loggedinUser, token)
	response := response.APIResponse("Loged In", http.StatusOK, "success", auth, nil)

	ctx.JSON(http.StatusOK, response)
}
