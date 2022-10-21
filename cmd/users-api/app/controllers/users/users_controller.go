package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/bookstore/cmd/users-api/app/domain/users"
	"github.com/sum-project/bookstore/cmd/users-api/app/services"
	"github.com/sum-project/bookstore/pkg/errors"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
