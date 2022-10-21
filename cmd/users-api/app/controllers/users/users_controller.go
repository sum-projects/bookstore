package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/bookstore/cmd/users-api/app/domain/users"
	"github.com/sum-project/bookstore/cmd/users-api/app/services"
	"github.com/sum-project/bookstore/pkg/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestErr("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, getErr := services.GetUser(userId)
	if err != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
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
