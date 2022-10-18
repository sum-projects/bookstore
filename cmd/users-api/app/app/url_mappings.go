package app

import (
	"github.com/sum-project/bookstore/cmd/users-api/app/controllers/ping"
	"github.com/sum-project/bookstore/cmd/users-api/app/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
