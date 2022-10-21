package users

import (
	"fmt"
	"github.com/sum-project/bookstore/pkg/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	result := userDB[u.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("user %d not found", u.Id))
	}

	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	u.DateUpdated = result.DateUpdated

	return nil
}

func (u *User) Save() *errors.RestErr {
	current := userDB[u.Id]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestErr(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user %d alreadt exists", u.Id))
	}
	userDB[u.Id] = u

	return nil
}
