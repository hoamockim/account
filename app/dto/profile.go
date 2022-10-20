package dto

import (
	"github.com/tipee/account/pkg/util"
)

type UserInfoReq struct {
	FullName         string `json:"fullName"`
	Address          string `json:"address,omitempty"`
	PhoneNumber      string `json:"phoneNumber"`
	ExtraPhoneNumber string `json:"extraPhoneNumber"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Gender           int8   `json:"gender,omitempty"`
	Birthday         string `json:"birthday,omitempty"`
}

type UserInfoRes struct {
	Code     string `json:"code"`
	FullName string `json:"fullName"`
	Active   bool   `json:"active"`
}

func (req *UserInfoReq) ValidateBeforeCreating() (matched bool) {
	return util.IsEmail(req.Email) && util.IsPassword(req.Password)
}
