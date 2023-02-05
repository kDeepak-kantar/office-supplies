package ulists

import (
	"errors"

	"github.com/Deepak/pkg/domain/ulist"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUserList(c *gin.Context)
	GetAllUserLists(c *gin.Context)
	UpdateUserList(c *gin.Context)
	GetAllApprovedUserLists(c *gin.Context)
	GetAllNotApprovedUserLists(c *gin.Context)
	SendRemainderrest(c *gin.Context)
}

type repository struct {
	Input
}

type Input struct {
	Ulist ulist.Domain
}

var (
	ErrOperationNotAllowed = errors.New("status is null")
	ErrInvalidInputlist    = errors.New("invalid input. ulist domain missing")
)

func validateInput(input Input) error {

	if input.Ulist == nil {
		return ErrInvalidInputlist
	}

	return nil
}

func Init(input Input) (Repository, error) {
	if err := validateInput(input); err != nil {
		return nil, err
	}
	return &repository{
		input,
	}, nil
}
