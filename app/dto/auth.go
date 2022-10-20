package dto

import "github.com/dgrijalva/jwt-go"

type SignInType int8

const (
	UserPass SignInType = iota
	Email
	Facebook
	PhoneNumber
	Refresh
)

const (
	emailPattern = "(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)"
	phonePattern = "^[0-9]{10,11}$"
)

type JwtTokenBody struct {
	SignInType   SignInType `json:"sign_in_type"`
	UserName     string     `json:"user_name"`
	PassWord     string     `json:"pass_word"`
	Email        string     `json:"email"`
	Facebook     string     `json:"facebook"`
	Phone        string     `json:"phone"`
	RefreshToken string     `json:"refresh_token"`
}

type AuthClaims struct {
	*jwt.StandardClaims
	UserCode string `json:"user_code"`
	Email    string `json:"email"`
}

type JwtTokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (bodyReq *JwtTokenBody) Validate() bool {
	return true
}
