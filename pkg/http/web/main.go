package web

import (
	"github.com/Deepak/pkg/http/web/usersession"
	"github.com/Deepak/pkg/storage/db/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Repository interface {
	GetLogin(c *gin.Context)
	GetOverview(c *gin.Context)
	NewMultiTemplate(templateMap map[string]map[string][]string) render.HTMLRender
	GetViews() map[string][]string
}

type repository struct {
	Input
}

type Input struct {
	User        user.Repository
	UserSession usersession.Repository
}

func Init(input Input) Repository {
	return &repository{
		input,
	}
}
