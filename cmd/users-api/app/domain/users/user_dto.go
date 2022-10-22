package users

import (
	"github.com/sum-project/bookstore/pkg/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Password    string `json:"-"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return errors.NewBadRequestErr("invalid email address")
	}
	return nil
}
