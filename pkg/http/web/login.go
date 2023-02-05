package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *repository) GetLogin(c *gin.Context) {
	fmt.Println("login to app")
	c.HTML(
		http.StatusOK,
		"login",
		gin.H{},
	)
}
