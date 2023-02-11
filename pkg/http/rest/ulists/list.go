package ulists

import (
	"fmt"
	"net/http"

	"github.com/Deepak/pkg/storage/db/userlist"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

const (
	ResponseMessage string = "Order Placed Successfully"
	ErrorResponse   string = "Order not placed"
)

type ItemRequest struct {
	ItemID   string `json:"itemid"`
	Quantity string `json:"quantity"`
}
type OrderRequest struct {
	UserID        string        `json:"userid"`
	Items         []ItemRequest `json:"items"`
	EmpName       string        `json:"employeename"`
	EmpEmail      string        `json:"emailid"`
	RequestedDate string        `json:"requesteddate"`
	DueDate       string        `json:"duedate"`
	Status        string        `json:"orderstatus"`
}

type OrderUpdate struct {
	OrderID  int           `json:"orderid"`
	UserID   string        `json:"userid"`
	Items    []ItemRequest `json:"items"`
	EmpName  string        `json:"employeename"`
	EmpEmail string        `json:"emailid"`
}

type Approved struct {
	OrderID int    `json:"orderid"`
	Status  string `json:"orderstatus"`
}

type UserOrder struct {
	UserID string `json:"userid"`
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
			ItemID:   itemx.ItemID,
		}
		items = append(items, a)

	}
	uuids, err := uuid.FromString(orderitems.UserID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
	}
	if lst := r.Ulist.CreateUserList(&userlist.Order{
		UserID:        &uuids,
		Items:         items,
		EmpName:       orderitems.EmpName,
		EmpEmail:      orderitems.EmpEmail,
		RequestedDate: orderitems.RequestedDate,
		DueDate:       orderitems.DueDate,
		Status:        orderitems.Status,
	}); lst != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse)
	}
	c.JSON(http.StatusOK, ResponseMessage)
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
	Uslist, err := r.Ulist.GetUserLists(orderitem.UserID)
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
	user, err := r.Ulist.UpdateUserListstat(ListStat.OrderID, ListStat.Status)
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
			ItemID:   itemx.ItemID,
		}
		items = append(items, as)

	}
	uuids, err := uuid.FromString(ListStat.UserID)
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return

	}
	lst, err := r.Ulist.UpdateUserList(&userlist.OrderUpdate{
		Id:       ListStat.OrderID,
		UserID:   uuids.String(),
		Items:    items,
		EmpName:  ListStat.EmpName,
		EmpEmail: ListStat.EmpEmail,
	})
	if err != nil {
		handleError(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, lst)
}
