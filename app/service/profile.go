package service

import (
	"errors"
	"github.com/tipee/account/app/dto"
	"github.com/tipee/account/db/models"
	"github.com/tipee/account/db/repositories"
	"github.com/tipee/account/pkg/util"
)

type QueryUserService interface {
	GetUserInfo(userCode string) (userInfo *dto.UserInfoRes, err error)
}

type CommandUserService interface {
	Register(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error)
	ActiveUser(req dto.UserInfoReq) (string, error)
	Verify(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error)
	ChangePassword(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error)
}

// Register create new user for system
func (srv *Instance) Register(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error) {
	if isValid := req.ValidateBeforeCreating(); !isValid {
		err = errors.New("email or password is not correct")
		return
	}
	//1: check account exist with the current information
	//if existed: return error with message user is existed
	var entity models.UserInfo
	entity, _ = srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.EmailType,
		UserCode:  req.Email,
	})

	entity = models.UserInfo{
		Code:           util.UUID8(),
		PriEmail:       req.Email,
		PriMobilePhone: req.PhoneNumber,
		FullName:       req.FullName,
	}

	// 2: create new account if it's not existed
	// create verify link and send via email
	// tracking register
	if err = srv.UserRepo.SaveUserInfo(&entity); err != nil {
		return
	}
	userInfo = &dto.UserInfoRes{
		Code:     entity.Code,
		FullName: entity.FullName,
	}
	return
}

func (srv *Instance) ActiveUser(req dto.UserInfoReq) (string, error) {
	return "", nil
}

// GetUserInfo get user's information
func (srv *Instance) GetUserInfo(userCode string) (*dto.UserInfoRes, error) {
	var res *dto.UserInfoRes
	entity, _ := srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.CodeType,
		UserCode:  userCode,
	})

	res = &dto.UserInfoRes{}
	res.Code = entity.Code
	res.FullName = entity.FullName
	return res, nil
}

// VerifyByEmail verify an account by email
func (srv *Instance) VerifyByEmail(email, password string) (userInfo *dto.UserInfoRes, err error) {
	entity, _ := srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.EmailType,
		Email:     email,
		PassWord:  password,
	})
	if err != nil {
		return
	}
	userInfo = &dto.UserInfoRes{}
	userInfo.Code = entity.Code
	userInfo.FullName = entity.FullName
	return
}

func (srv *Instance) ChangePassword(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error) {
	return
}
