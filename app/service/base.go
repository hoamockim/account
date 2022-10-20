package service

import (
	"github.com/tipee/account/db/repositories"
)

type Instance struct {
	UserRepo interface {
		repositories.UserInfoRepository
		repositories.UserAttributeRepository
	}
}

var userService interface {
	QueryUserService
	CommandUserService
}

var srv *Instance

func init() {
	initService(repositories.New())
	//initJwtParse()
}

// initService
func initService(userInterface interface {
	repositories.UserInfoRepository
	repositories.UserAttributeRepository
}) {
	srv = &Instance{UserRepo: userInterface}
}

// GetUserService get user service
func GetUserService() interface {
	QueryUserService
	CommandUserService
} {
	return userService
}

// GetAuthService get authentication service
func GetAuthService() AuthService {
	return srv
}
