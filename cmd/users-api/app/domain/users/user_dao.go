package users

import (
	"github.com/sum-project/bookstore/cmd/users-api/app/datasources/mysql/user_db"
	"github.com/sum-project/bookstore/cmd/users-api/app/logger"
	"github.com/sum-project/bookstore/pkg/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, status, password) VALUES (?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, status, password, date_created, date_updated FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, status, password, date_created, date_updated FROM users WHERE status=?;"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err = result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.Password, &u.DateCreated, &u.DateUpdated); err != nil {
		logger.Error("error when trying to get user statement", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password)
	if err != nil {
		logger.Error("error when trying to prepare save user", err)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}
	u.Id = userId

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.Id)
	if err != nil {
		logger.Error("error when trying to prepare update user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := user_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.Password, &u.DateCreated, &u.DateUpdated); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	return results, nil
}
