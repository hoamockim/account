package models

import "strings"

const (
	tblUserInfo = "user_profile"
)

//UserInfo
type UserInfo struct {
	BaseModel
	Code             string
	FullName         string
	Gender           int8
	Dob              uint
	BirthDay         string
	PriEmail         string
	ExtraEmail       string
	PriMobilePhone   string
	ExtraMobilePhone string
	PriAddress       string
	ExtraAddress1    string
	ExtraAddress2    string
}

//Validate
func (user *UserInfo) Validate() bool {
	return !(strings.TrimSpace(user.FullName) == "" ||
		user.Dob < 0 ||
		strings.TrimSpace(user.PriEmail) == "")
}

func (user *UserInfo) GetTableName() string {
	return tblUserInfo
}
