package access_token

import (
	"github.com/sum-project/bookstore/cmd/oauth-api/app/domain/access_token"
	"github.com/sum-project/bookstore/cmd/oauth-api/app/repository/db"
	"github.com/sum-project/bookstore/pkg/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type service struct {
	repository db.DatabaseRepository
}

func NewService(repo db.DatabaseRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
