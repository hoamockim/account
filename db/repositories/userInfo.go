package repositories

import (
	"errors"

	"github.com/tipee/account/db/models"
)

type UserFilterType int8

const (
	AccountType UserFilterType = 1 + iota
	EmailType
	PhoneType
	CodeType
	IdType
)

type UserFilter struct {
	InputType   UserFilterType // 1: {username & pass }, 2: email, 3: phoneNumber, 4: userCode
	UserName    string
	PassWord    string
	Email       string
	PhoneNumber string
	UserCode    string
	UserId      string
}

type UserInfoRepository interface {
	GetUserInfo(usf *UserFilter) (userInfo models.UserInfo, err error)
	SaveUserInfo(userInfo *models.UserInfo) error
}

func (repo *repo) GetUserInfo(usf *UserFilter) (userInfo models.UserInfo, err error) {
	var filter []FilterField

	switch usf.InputType {
	case IdType:
		filter = append(filter, FilterField{Column: "id", Value: usf.UserId})
	case AccountType: // user & pass
		filter = append(filter, FilterField{Column: "username", Value: usf.UserName})
		filter = append(filter, FilterField{Column: "password", Value: usf.PassWord})
	case EmailType:
		filter = append(filter, FilterField{Column: "email", Value: usf.Email})
	case PhoneType:
		filter = append(filter, FilterField{Column: "pri_mobi_phone", Value: usf.PhoneNumber})
	case CodeType:
		filter = append(filter, FilterField{Column: "code", Value: usf.UserCode})
	default:
		return userInfo, errors.New("input type is invalid")
	}

	//TODO: call filter
	return
}

func (repo *repo) SaveUserInfo(userInfo *models.UserInfo) error {
	if !userInfo.Validate() {
		return errors.New(models.DataInvalid)
	}
	return repo.save(userInfo)
}
