// Package health is responsible for delering the health probe needed by our infrastructure to determine if the app is live or not.
package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Repository provides access for retrieving service health information.
type Repository interface {
	GetHealth(*gin.Context)
}

// Adapter is a structure that contains the necessary dependencies for
// service health.
type repository struct{}

// New creates an adapter given the necessary dependencies.
func Init() Repository {
	return &repository{}
}

// GetHealth retrieves the service health information given a request context.
func (a *repository) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
