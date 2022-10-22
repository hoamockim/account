package service

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tipee/account/app/dto"
	"github.com/tipee/account/pkg/configs"
	"github.com/tipee/account/pkg/util"
)

type AuthService interface {
	SignIn(reqBody *dto.JwtTokenBody) (*dto.JwtTokenData, error)
	RefreshToken(reqBody dto.JwtTokenBody) (dto.JwtTokenData, error)
	Verify(data string, claims jwt.Claims) error
}

const (
	RS256 = "RS256"
)

type jwtParse struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	parser     *jwt.Parser
}

var jwp *jwtParse

func InitJwtParse() {
	if jwp != nil {
		return
	}
	priBytes, pubBytes := readKeys(configs.GetJwtKey())
	priKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priBytes)
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	jwp = &jwtParse{
		privateKey: priKey,
		publicKey:  pubKey,
		parser: &jwt.Parser{
			ValidMethods:         []string{RS256},
			SkipClaimsValidation: false,
		},
	}
}

func readKeys(priKeyPath, pubKeyPath string) (priBytes, pubBytes []byte) {
	priBytes, _ = ioutil.ReadFile(priKeyPath)
	pubBytes, _ = ioutil.ReadFile(pubKeyPath)

	if len(priBytes) == 0 || len(pubBytes) == 0 {
		panic("jwt is invalid")
	}
	return
}

func (srv *Instance) Verify(data string, claims jwt.Claims) (err error) {
	if _, err = jwt.ParseWithClaims(data, claims, func(token *jwt.Token) (interface{}, error) {
		return jwp.publicKey, nil
	}); err != nil {
		return
	}

	if authClaim, ok := claims.(*dto.AuthClaims); ok {
		if time.Unix(authClaim.ExpiresAt, 0).Sub(time.Now().UTC()) > 30*time.Second {
			return errors.New("token is expired")
		}
	}
	return
}

func (srv *Instance) SignIn(reqBody *dto.JwtTokenBody) (tokenData *dto.JwtTokenData, err error) {
	var userInfo *dto.UserInfoRes
	if userInfo, err = srv.VerifyByEmail(reqBody.Email, reqBody.PassWord); err == nil {
		return
	}
	claims := dto.AuthClaims{
		Email:    reqBody.Email,
		UserCode: userInfo.Code,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: util.ConvertTimestampToMilliSecond(time.Now().Add(24 * time.Hour).Unix()),
			IssuedAt:  int64(util.MakeCurrentTimestampMilliSecond()),
		},
	}

	tkn := jwt.NewWithClaims(jwt.GetSigningMethod(RS256), claims)
	var token string
	if token, err = tkn.SignedString(jwp.privateKey); err != nil {
		return
	}

	//set value for response data
	tokenData = &dto.JwtTokenData{
		AccessToken: token,
	}
	return
}

func (srv *Instance) RefreshToken(reqBody dto.JwtTokenBody) (resData dto.JwtTokenData, err error) {
	return
}
