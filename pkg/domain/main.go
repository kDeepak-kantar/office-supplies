package domain

import (
	"errors"

	"github.com/Deepak/pkg/domain/auth"
	"github.com/Deepak/pkg/domain/ulist"
	"github.com/Deepak/pkg/storage/db/user"
	"github.com/Deepak/pkg/storage/db/userlist"
)

type Domains struct {
	Auth  auth.Domain
	Ulist ulist.Domain
}

type Input struct {
	User     user.Repository
	UserList userlist.Repository
}

var (
	ErrInvalidInputUser = errors.New("invalid input. User repository missing")
	ErrInvalidInputCD   = errors.New("invalid input. Coffee date repository missing")
)

func validateInput(input Input) error {
	if input.User == nil {
		return ErrInvalidInputUser
	}
	if input.UserList == nil {
		return ErrInvalidInputCD
	}

	return nil
}

func Util() error {
	return nil
}
