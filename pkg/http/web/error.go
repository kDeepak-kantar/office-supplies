package web

import (
	"fmt"
	"net/http"

	"github.com/Deepak/pkg/logger"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, code uint, err error) {
	var message string

	switch err := err.(type) {
	default:
		logger.LogContextError(err, c)
		message = fmt.Sprintf("Encountered an Unexpected Error (%s)", err.Error())
	}

	logger.Log("Handle Error", logger.SeverityError, message, err, nil)
	c.HTML(
		http.StatusOK,
		"error",
		gin.H{
			"basicErrorMsg": message,
			"errorMsg":      err.Error(),
		})
}
