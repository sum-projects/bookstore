package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/bookstore/cmd/oauth-api/app/clients/cassandra"
	"github.com/sum-project/bookstore/cmd/oauth-api/app/http"
	"github.com/sum-project/bookstore/cmd/oauth-api/app/repository/db"
	"github.com/sum-project/bookstore/cmd/oauth-api/app/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
