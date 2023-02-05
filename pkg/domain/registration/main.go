package registration

// import (
// 	"errors"

// 	"github.com/Deepak/pkg/storage/db/user"
// 	"github.com/Deepak/pkg/storage/userlist"
// )

// type Domain interface {
// 	Register(userID string) error
// }

// type domain struct {
// 	Input
// }

// type Input struct {
// 	User     user.Repository
// 	UserList userlist.Repository
// }

// var (
// 	ErrInvalidInputUser = errors.New("invalid input. User repository missing")
// 	ErrInvalidInputCD   = errors.New("invalid input. Users List repository missing")
// 	ErrNotFound         = errors.New("not found")
// )

// func validateInput(input Input) error {
// 	if input.User == nil {
// 		return ErrInvalidInputUser
// 	}

// 	if input.UserList == nil {
// 		return ErrInvalidInputCD
// 	}

// 	return nil
// }

// func Init(input Input) (Domain, error) {
// 	if err := validateInput(input); err != nil {
// 		return nil, err
// 	}

// 	return &domain{
// 		input,
// 	}, nil
// }
