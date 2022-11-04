package access_token

import (
	"github.com/sum-project/bookstore/pkg/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestErr("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestErr("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestErr("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestErr("invalid expiration time")
	}

	return nil
}

func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}
