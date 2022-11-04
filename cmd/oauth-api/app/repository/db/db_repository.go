package db

import (
	"github.com/sum-project/bookstore/cmd/oauth-api/app/domain/access_token"
	"github.com/sum-project/bookstore/pkg/errors"
)

func NewRepository() DatabaseRepository {
	return &dbRepository{}
}

type DatabaseRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implement yet")
}
