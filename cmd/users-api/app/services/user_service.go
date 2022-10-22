package services

import (
	"github.com/sum-project/bookstore/cmd/users-api/app/domain/users"
	"github.com/sum-project/bookstore/pkg/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err = user.Validate(); err != nil {
		return nil, err
	}

	if user.FirstName != "" || !isPartial {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" || !isPartial {
		current.LastName = user.LastName
	}
	if user.Email != "" || !isPartial {
		current.Email = user.Email
	}
	if user.Status != "" || !isPartial {
		current.Status = user.Status
	}
	if user.Password != "" || !isPartial {
		current.Password = user.Password
	}

	if err = current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
