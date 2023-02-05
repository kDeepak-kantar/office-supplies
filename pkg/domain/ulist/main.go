package ulist

import (
	"errors"

	"github.com/Deepak/pkg/storage/db/user"
	"github.com/Deepak/pkg/storage/userlist"
)

type Domain interface {
	CreateUserList(c *userlist.Order) error
	GetAllUserLists() ([]*userlist.Order, error)
	GetUserListStatus(id int) (string, error)
	UpdateUserList(id int, status string) (*userlist.Order, error)
	GetAllApprovedUserLists() ([]*userlist.Order, error)
	GetAllNotApprovedUserLists() ([]*userlist.Order, error)
	SendRemainder() ([]string, error)
}

type domain struct {
	Input
}

type Input struct {
	User     user.Repository
	UserList userlist.Repository
}

var (
	ErrInvalidInputUser = errors.New("invalid input. User repository missing")
	ErrInvalidInputCD   = errors.New("invalid input. userlist repository missing")
	ErrNotFound         = errors.New("not found")
	ErrInvalidInputlist = errors.New("invalid userlist repository missing")
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

func Init(input Input) (Domain, error) {
	if err := validateInput(input); err != nil {
		return nil, err
	}

	return &domain{
		input,
	}, nil
}
