package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, errorCode int, err error) {
	respStruct := struct {
		Code    int     `json:"code"`
		Message *string `json:"msg"`
	}{
		Code: errorCode,
	}

	if err != nil {
		errMsg := err.Error()
		respStruct.Message = &errMsg
	}

	switch err {
	default:
		respStruct.Code = http.StatusInternalServerError
	}

	c.JSON(errorCode, respStruct)
}
