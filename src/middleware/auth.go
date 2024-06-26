package middleware

import (
	"errors"
	"net/http"
	"strings"
	"template/src/filter"
	"template/src/models"
	"template/src/services"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yudhaananda/go-common/env"
	"github.com/yudhaananda/go-common/paging"
	"github.com/yudhaananda/go-common/response"
)

type Interface interface {
	AuthMiddleware(c *gin.Context)
}

type authMiddleware struct {
	service *services.Services
}

type InitParam struct {
	Service *services.Services
}

func Init(params InitParam) Interface {
	return &authMiddleware{service: params.Service}
}

func (a *authMiddleware) AuthMiddleware(ctx *gin.Context) {
	authheader := ctx.GetHeader("Authorization")

	if !strings.Contains(authheader, "Bearer") {
		response := response.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil, nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	tokenString := ""
	arrayToken := strings.Split(authheader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, err := a.validateToken(tokenString)

	if err != nil {
		response := response.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil, err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		response := response.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil, errors.New("invalid token").Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userId := int(claim["user_id"].(float64))

	dateTime, err := time.Parse(time.RFC3339Nano, claim["time"].(string))

	if err != nil {
		response := response.APIResponse("Error Parse Date", http.StatusUnauthorized, "error", nil, err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if dateTime.Before(time.Now()) {
		response := response.APIResponse("Session End", http.StatusUnauthorized, "error", nil, errors.New("session end").Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	paging := paging.Paging[filter.UserFilter]{}
	paging.SetDefault()
	paging.Filter.Id = userId
	users, _, err := a.service.User.Get(ctx, paging)

	if err != nil {
		response := response.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil, err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	if len(users) == 0 {
		response := response.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil, errors.New("no user found").Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := users[0]

	ctx.Set(models.UserKey, user)
}

func (s *authMiddleware) validateToken(token string) (*jwt.Token, error) {
	encodeToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		secret, err := env.GetSecret()
		if err != nil {
			return nil, err
		}
		return []byte(secret), nil
	})

	if err != nil {
		return encodeToken, err
	}

	return encodeToken, nil
}
