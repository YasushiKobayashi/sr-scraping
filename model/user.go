package model

import (
	"errors"

	"github.com/YasushiKobayashi/countrobu/app_error"
)

type (
	User struct {
		AccountId string
		Password  string
	}
)

func SetUser(accountId string, password string) (*User, error) {
	u := &User{}
	if accountId == "" {
		err := errors.New("accountId is required")
		return u, app_error.NewBadRequestErr(err)
	}

	if password == "" {
		err := errors.New("password is required")
		return u, app_error.NewBadRequestErr(err)
	}

	u.AccountId = accountId
	u.Password = password
	return u, nil
}
