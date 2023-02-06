package auth

import (
	"errors"
	"net/http"

	"github.com/Deepak/pkg/domain/auth"
	"github.com/Deepak/pkg/http/web/usersession"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email     string `json:"email"`
	AuthToken string `json:"idtoken"`
}

type AdminRequest struct {
	UserID string
	Action string
}

var (
	ErrOperationNotAllowed = errors.New("operation not allowed")
	ErrInvalidAction       = errors.New("you cannot remove yourself")
)

func (r *repository) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}

	resp, err := r.Auth.LoginUser(auth.LoginRequest{
		Email: req.Email,
		Token: req.AuthToken,
	})
	if err != nil {
		handleError(c, http.StatusUnauthorized, err)
		return
	}

	err = r.UserSession.SetCookie(&usersession.SetCookieRequest{
		Context: c,
		UserID:  resp.User.ID.String(),
	})
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp.User)
}

func (r *repository) GetAllUsers(c *gin.Context) {
	//get the logged in user id and get his role, if not admin, send unauthorized error
	userId := c.Param("id")
	role, err := r.Auth.GetUserRole(userId)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if role != "admin" {
		handleError(c, http.StatusUnauthorized, ErrOperationNotAllowed)
		return
	}
	allUsers, err := r.Auth.GetAllUsers()
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, allUsers)
}

func (r *repository) Admin(c *gin.Context) {
	var req AdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if req.Action != "admin" {
		handleError(c, http.StatusInternalServerError, ErrOperationNotAllowed)
		return
	}

	userId := c.Param("id")
	role, err := r.Auth.GetUserRole(userId)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if role != "admin" {
		handleError(c, http.StatusUnauthorized, ErrOperationNotAllowed)
		return
	}

	user, err := r.Auth.AdminAccess(req.UserID)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
func (r *repository) RemoveUser(c *gin.Context) {
	var req AdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if req.Action != "Remove User" {
		handleError(c, http.StatusInternalServerError, ErrOperationNotAllowed)
		return
	}

	userId := c.Param("id")
	role, err := r.Auth.GetUserRole(userId)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if role != "admin" {
		handleError(c, http.StatusUnauthorized, ErrOperationNotAllowed)
		return
	}
	if userId == req.UserID {
		handleError(c, http.StatusUnauthorized, ErrInvalidAction)
		return
	}
	err = r.Auth.RemoveUser(req.UserID)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
