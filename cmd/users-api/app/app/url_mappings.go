package app

import (
	"github.com/sum-project/bookstore/cmd/users-api/app/controllers/ping"
	"github.com/sum-project/bookstore/cmd/users-api/app/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
