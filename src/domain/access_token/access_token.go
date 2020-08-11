package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/shakilbd009/go-oauth-api/src/utils/crypto_utils"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	//Used for password grant_type
	UserName string `json:"username"`
	Password string `json:"password"`
	//used for client_credentials
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() rest_errors.RestErr {

	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		rest_errors.NewBadRequestError("invalid grant_type parameter")
	}
	//TODO: validate rests
	return nil
}

func (at *AccessToken) Validate() rest_errors.RestErr {

	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		rest_errors.NewBadRequestError("invalid access token")
	}
	if at.ClientID <= 0 {
		rest_errors.NewBadRequestError("invalid client ID")
	}
	if at.Expires <= 0 {
		rest_errors.NewBadRequestError("invalid expiration time")
	}
	if at.UserID <= 0 {
		rest_errors.NewBadRequestError("invalid user ID")
	}
	return nil
}

func NewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println(expirationTime)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
