package app

import "github.com/sum-project/bookstore/cmd/users-api/app/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
