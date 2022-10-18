package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sum-project/bookstore/cmd/users-api/app/domain/users"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err = json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle error
		return
	}

	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}
