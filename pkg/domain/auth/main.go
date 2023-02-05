package auth

import (
	"errors"

	"github.com/Deepak/pkg/storage/db/user"
)

type Domain interface {
	LoginUser(req LoginRequest) (*LoginRespose, error)
	GetUserRole(userId string) (string, error)
	GetAllUsers() ([]*user.User, error)
	AdminAccess(userId string) (*user.User, error)
	RemoveUser(userId string) error
	Scheduler()
}

type domain struct {
	Input
}

type Input struct {
	User user.Repository
}

var (
	ErrInvalidInputUser = errors.New("invalid input. User repository missing")
)

func validateInput(input Input) error {
	if input.User == nil {
		return ErrInvalidInputUser
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
