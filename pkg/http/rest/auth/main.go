package auth

import (
	"github.com/Deepak/pkg/domain/auth"
	"github.com/Deepak/pkg/http/web/usersession"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Login(c *gin.Context)
	GetAllUsers(c *gin.Context)
	Admin(c *gin.Context)
	RemoveUser(c *gin.Context)
}

type repository struct {
	Input
}

type Input struct {
	Auth        auth.Domain
	UserSession usersession.Repository
}

func Init(input Input) Repository {
	return &repository{
		input,
	}
}
