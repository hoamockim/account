package services

import (
	"github.com/tipee/account/db/repositories"
)

type Instance struct {
	UserRepo interface {
		repositories.UserInfoRepository
		repositories.UserAttributeRepository
	}
}

var (
	srv         *Instance
	userService interface {
		QueryUserService
		CommandUserService
	}
)

// initService
func InitService(userRepos interface {
	repositories.UserInfoRepository
	repositories.UserAttributeRepository
}) {
	srv = &Instance{UserRepo: userRepos}
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
