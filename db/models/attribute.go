package models

const (
	tblUserAttribute = "user_attribute"
)

type UserAttribute struct {
	BaseModel
	UserId int
	Key    string
	Value  string
}

func (usat *UserAttribute) Validate() bool {
	return true
}

func (usat *UserAttribute) GetTableName() string {
	return tblUserAttribute
}

func (usat *UserAttribute) IsCached() bool {
	return true
}
