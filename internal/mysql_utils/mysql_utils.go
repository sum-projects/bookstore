package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/sum-project/bookstore/pkg/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseErr(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundErr("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestErr("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
