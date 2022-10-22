package users

import (
	"fmt"
	"github.com/sum-project/bookstore/cmd/users-api/app/datasources/mysql/user_db"
	"github.com/sum-project/bookstore/pkg/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, status, password) VALUES (?, ?, ?, ?, ?);"
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
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, "", "")
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	u.Id = userId

	return nil
}
