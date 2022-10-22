package users

import (
	"fmt"
	"github.com/sum-project/bookstore/cmd/users-api/app/datasources/mysql/user_db"
	"github.com/sum-project/bookstore/internal/mysql_utils"
	"github.com/sum-project/bookstore/pkg/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, status, password) VALUES (?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, status, password, date_created, date_updated FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err = result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.Password, &u.DateCreated, &u.DateUpdated); err != nil {
		return mysql_utils.ParseErr(err)
	}

	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password)
	if err != nil {
		return mysql_utils.ParseErr(err)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	u.Id = userId

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.Id)
	if err != nil {
		mysql_utils.ParseErr(err)
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		mysql_utils.ParseErr(err)
	}

	return nil
}
