package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shakilbd009/go-oauth-api/src/http"
	"github.com/shakilbd009/go-oauth-api/src/repository/db"
	"github.com/shakilbd009/go-oauth-api/src/repository/rest"
	"github.com/shakilbd009/go-oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApp() {

	atService := access_token.NewService(
		rest.NewRestUsersRepository(),
		db.NewDbRepository(),
	)
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8082")
}
