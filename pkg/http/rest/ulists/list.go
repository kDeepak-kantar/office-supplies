package ulists

import (
	"fmt"
	"net/http"

	"github.com/Deepak/pkg/storage/userlist"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type ItemRequest struct {
	Id       string `json:"id"`
	Quantity string `json:"qty"`
}
type OrderRequest struct {
	UserID        string        `json:"userid"`
	Items         []ItemRequest `json:"items"`
	EmpName       string        `json:"employeeName"`
	EmpEmail      string        `json:"email"`
	RequestedDate string        `json:"requestedDate"`
	DueDate       string        `json:"dueDate"`
	Status        string        `json:"status"`
}

type OrderUpdate struct {
	Id       int           `json:"id"`
	UserID   string        `json:"userid"`
	Items    []ItemRequest `json:"items"`
	EmpName  string        `json:"employeeName"`
	EmpEmail string        `json:"email"`
}

type Approved struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type UserOrder struct {
	UserId string `json:"userid"`
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
	uuids, err := uuid.FromString(orderitems.UserID)
	if err != nil {
		panic(err)
	}
	lst := r.Ulist.CreateUserList(&userlist.Order{
		UserID:        &uuids,
		Items:         items,
		EmpName:       orderitems.EmpName,
		EmpEmail:      orderitems.EmpEmail,
		RequestedDate: orderitems.RequestedDate,
		DueDate:       orderitems.DueDate,
		Status:        orderitems.Status,
	})
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

func (r *repository) GetUserList(c *gin.Context) {
	var orderitem UserOrder

	err := c.ShouldBindJSON(&orderitem)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	Uslist, err := r.Ulist.GetUserLists(orderitem.UserId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Uslist)
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
func (r *repository) UpdateUserListstat(c *gin.Context) {
	var ListStat Approved
	err := c.ShouldBindJSON(&ListStat)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	user, err := r.Ulist.UpdateUserListstat(ListStat.ID, ListStat.Status)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

func (r *repository) UpdateUserList(c *gin.Context) {
	var ListStat OrderUpdate
	err := c.ShouldBindJSON(&ListStat)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	items := []userlist.Item{}

	for _, itemx := range ListStat.Items {
		as := userlist.Item{
			Quantity: itemx.Quantity,
			ItemID:   itemx.Id,
		}
		items = append(items, as)

	}
	uuids, err := uuid.FromString(ListStat.UserID)
	if err != nil {
		panic(err)
	}
	lst, err := r.Ulist.UpdateUserList(&userlist.OrderUpdate{
		Id:       ListStat.Id,
		UserID:   uuids.String(),
		Items:    items,
		EmpName:  ListStat.EmpName,
		EmpEmail: ListStat.EmpEmail,
	})
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
	}

	c.JSON(http.StatusOK, lst)
}
