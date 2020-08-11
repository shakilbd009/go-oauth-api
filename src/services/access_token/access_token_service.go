package access_token

import (
	"strings"

	"github.com/shakilbd009/go-oauth-api/src/domain/access_token"
	"github.com/shakilbd009/go-oauth-api/src/repository/db"
	"github.com/shakilbd009/go-oauth-api/src/repository/rest"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr)
	UpdateExpirationTime(*access_token.AccessToken) rest_errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(userRepo rest.RestUsersRepository, repo db.DbRepository) Service {
	return &service{
		restUsersRepo: userRepo,
		dbRepo:        repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, rest_errors.RestErr) {

	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {

	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := s.restUsersRepo.LoginUser(request.UserName, request.Password)
	if err != nil {
		return nil, err
	}
	//Generate new access token
	at := access_token.NewAccessToken(user.Id)
	at.Generate()

	//Save the new access token in cassandra
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at *access_token.AccessToken) rest_errors.RestErr {

	if err := at.Validate(); err != nil {
		return err
	}
	return s.UpdateExpirationTime(at)
}
