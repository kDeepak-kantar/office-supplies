package usersession

import (
	"github.com/gin-gonic/gin"
)

type Repository interface {
	SetCookie(req *SetCookieRequest) error
	GetCookie(c *gin.Context) (*Cookie, error)
	DeleteCookie(c *gin.Context)
}

type service struct {
}

func Init() Repository {
	return &service{}
}
