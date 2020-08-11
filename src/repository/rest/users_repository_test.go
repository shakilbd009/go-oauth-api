package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `"email": "email@gmail.com", "password": "123"`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "123")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `"email": "email@gmail.com", "password": "123"`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials","status": "404","error":"not_found "}`,
	})
	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "123")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `"email": "email@gmail.com", "password": "123"`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials","status": 404,"error":"not_found "}`,
	})
	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "123")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `"email": "email@gmail.com", "password": "123"`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
			"id": "1",
			"date_created": "2020-07-24 14:27:22",
			"status": "active"
		}`,
	})
	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "123")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `"email": "email@gmail.com", "password": "123"`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
			"id": 1,
			"date_created": "2020-07-24 14:27:22",
			"status": "active"
		}`,
	})
	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
}
