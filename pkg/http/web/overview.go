package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *repository) GetOverview(c *gin.Context) {
	cookie, err := r.UserSession.GetCookie(c)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	if cookie == nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	user, err := r.User.GetUserByStringId(cookie.UserID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	if user == nil {
		handleError(c, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	curYear, curWeek := time.Now().ISOWeek()

	// get current week group for user, if any
	// curDate, err := r.CoffeeDate.GetForWeekAndUser(uint(curWeek), uint(curYear), user.ID.String())
	// if err != nil {
	// 	handleError(c, http.StatusInternalServerError, err)
	// 	return
	// }

	// // get future schedule for current user (are they registered to for dates and on which weeks)
	// datesForUser, err := r.CoffeeDate.GetAllUserRegistryForFuture(user.ID.String(), uint(curWeek), uint(curYear))
	// if err != nil {
	// 	handleError(c, http.StatusInternalServerError, err)
	// 	return
	// }

	c.HTML(
		http.StatusOK,
		"overview",
		gin.H{
			"title":  fmt.Sprintf("Hello %s", user.Name),
			"week":   curWeek,
			"year":   curYear,
			"userID": cookie.UserID,
		},
	)
}
