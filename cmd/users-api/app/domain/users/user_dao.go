package users

import (
	"fmt"
	"github.com/sum-project/bookstore/cmd/users-api/app/datasources/mysql/user_db"
	"github.com/sum-project/bookstore/pkg/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, status, password) VALUES (?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, date_updated FROM users WHERE id=?;"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err = result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.DateUpdated); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when thrying to get user %d: %s", u.Id, err.Error()))
	}

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
