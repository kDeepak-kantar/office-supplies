package ulists

import (
	"fmt"
	"net/http"

	"github.com/Deepak/pkg/storage/userlist"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type ItemRequest struct {
	Quantity string `json:"qty"`
	Id       string `json:"id"`
}
type OrderRequest struct {
	UserID        string        `json:"userid"`
	Items         []ItemRequest `json:"items"`
	RequestedDate string        `json:"requestedDate"`
	DueDate       string        `json:"dueDate"`
	// LastStatusUpdate string        `json:"lastStatusUpdate"`
	Status string `json:"status"`
}

type Approved struct {
	ID         int    `json:"id"`
	LastUpdate string `json:"lastStatusUpdate"`
	Status     string `json:"status"`
}

func (r *repository) CreateUserList(c *gin.Context) {
	var orderitems OrderRequest

	err := c.ShouldBindJSON(&orderitems)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	items := []userlist.Item{}

	for _, itemx := range orderitems.Items {
		a := userlist.Item{
			Quantity: itemx.Quantity,
			ItemID:   itemx.Id,
		}
		items = append(items, a)

	}
	// requestedDate, err := time.Parse("02/01/2006", orderitems.RequestedDate)
	// if err != nil {
	// 	handleError(c, http.StatusInternalServerError, err)
	// 	return
	// }
	// dueDate, err := time.Parse("02/01/2006", orderitems.DueDate)
	// if err != nil {
	// 	handleError(c, http.StatusInternalServerError, err)
	// 	return
	// }
	uuids, err := uuid.FromString(orderitems.UserID)
	if err != nil {
		panic(err)
	}
	lst := r.Ulist.CreateUserList(&userlist.Order{
		UserID:        &uuids,
		Items:         items,
		RequestedDate: orderitems.RequestedDate,
		DueDate:       orderitems.DueDate,
		// LastStatusUpdate: orderitems.LastStatusUpdate,
		Status: orderitems.Status,
	})
	// lst := r.Ulist.CreateUserList(&twin)
	c.JSON(http.StatusOK, lst)
}
func (r *repository) GetAllUserLists(c *gin.Context) {
	allUsers, err := r.Ulist.GetAllUserLists()
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, allUsers)
}

func (r *repository) SendRemainderrest(c *gin.Context) {
	remainder, err := r.Ulist.SendRemainder()
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, remainder)
}

func (r *repository) GetAllApprovedUserLists(c *gin.Context) {
	allUsers, err := r.Ulist.GetAllApprovedUserLists()
	fmt.Println(allUsers)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, allUsers)
}
func (r *repository) GetAllNotApprovedUserLists(c *gin.Context) {
	allUsers, err := r.Ulist.GetAllNotApprovedUserLists()
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, allUsers)
}
func (r *repository) UpdateUserList(c *gin.Context) {
	var ListStat Approved
	err := c.ShouldBindJSON(&ListStat)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	user, err := r.Ulist.UpdateUserList(int(ListStat.ID), ListStat.Status)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	c.JSON(http.StatusOK, user)

}
