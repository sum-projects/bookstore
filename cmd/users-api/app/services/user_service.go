package services

import (
	"github.com/sum-project/bookstore/cmd/users-api/app/domain/users"
	"github.com/sum-project/bookstore/internal/crypto_utils"
	"github.com/sum-project/bookstore/pkg/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User, bool) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
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

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
