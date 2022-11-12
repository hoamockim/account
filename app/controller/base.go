package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tipee/account/app/services"
	"github.com/tipee/account/pkg/errors"
)

type Meta struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

const (
	PROFILE = "profile"
	AUTH    = "auth"
)

var authService services.AuthService
var userService interface {
	services.CommandUserService
	services.QueryUserService
}

func init() {
	authService = services.GetAuthService()
	userService = services.GetUserService()
}

func success(ctx *gin.Context, data interface{}) {
	res := &BaseResponse{
		Meta: Meta{
			Code:    "200",
			Message: "success",
		},
	}
	ctx.JSON(http.StatusOK, res)
}

func error(ctx *gin.Context, httpCode int, serviceName, errorCode, errorMessage string) {
	appErr := errors.New(serviceName, errorCode, errorMessage)
	ctx.JSON(httpCode, appErr)
}
