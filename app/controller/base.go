package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tipee/account/app/dto"
	"github.com/tipee/account/app/service"
	"github.com/tipee/account/pkg/errors"
)

type Meta struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

const (
	PROFILE = "profile"
	AUTH    = "auth"
)

var authService service.AuthService
var userService interface {
	service.CommandUserService
	service.QueryUserService
}

func init() {
	authService = service.GetAuthService()
	userService = service.GetUserService()
}

func success(ctx *gin.Context, data interface{}) {
	res := &dto.ResData{Data: data}
	ctx.JSON(http.StatusOK, res)
}

func error(ctx *gin.Context, httpCode int, serviceName, errorCode, errorMessage string) {
	appErr := errors.New(serviceName, errorCode, errorMessage)
	ctx.JSON(httpCode, appErr)
}
