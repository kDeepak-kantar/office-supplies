package ulist

import (
	"testing"

	usermock "github.com/Deepak/pkg/storage/db/user/mocks"
	"github.com/Deepak/pkg/storage/userlist"
	ulistmock "github.com/Deepak/pkg/storage/userlist/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserList(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)
	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	uslist, _ := Init(input)
	usrlist := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "",
	}
	myuliststorage.On("CreateUserList", mock.Anything).Return(nil)
	err := uslist.CreateUserList(usrlist)
	assert.Nil(t, err)
}

func TestGetAllUserLists(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)

	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	usrslist, _ := Init(input)
	alllist := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "",
	}
	orderlist := []*userlist.Order{}
	orderlist = append(orderlist, alllist)

	myuliststorage.On("GetAllist", mock.Anything).Return(orderlist, nil)
	orderlist, err := usrslist.GetAllUserLists()
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
}

func TestGetAllApprovedUserLists(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)

	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	usrslist, _ := Init(input)
	alllist := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "approved",
	}
	orderlist := []*userlist.Order{}
	orderlist = append(orderlist, alllist)

	myuliststorage.On("GetAllApproved", mock.Anything).Return(orderlist, nil)
	orderlist, err := usrslist.GetAllApprovedUserLists()
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
}
func TestGetAllNotApprovedUserLists(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)

	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	usrslist, _ := Init(input)
	alllist := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "pending",
	}
	orderlist := []*userlist.Order{}
	orderlist = append(orderlist, alllist)

	myuliststorage.On("GetAllNotApproved", mock.Anything).Return(orderlist, nil)
	orderlist, err := usrslist.GetAllNotApprovedUserLists()
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
}

func TestGetUserLists(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)
	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	uslist, _ := Init(input)
	usrid := "id"
	usrlist := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "",
	}
	orderlist := []*userlist.Order{}
	orderlist = append(orderlist, usrlist)
	myuliststorage.On("GetUserList", "id").Return(orderlist, nil)
	orderlist, err := uslist.GetUserLists(usrid)
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
}

func TestUpdateUserListstat(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)
	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	uslist, _ := Init(input)
	usrin := 123
	usrstatus := "approved"
	orderinput := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "pending",
	}
	orderoutput := &userlist.Order{
		UserID:        nil,
		Items:         nil,
		EmpName:       "emp1",
		EmpEmail:      "emp@blackwoodseven.com",
		RequestedDate: "",
		DueDate:       "",
		Status:        "approved",
	}
	myuliststorage.On("GetOrderByStringiId", usrin).Return(orderinput, nil)
	myuliststorage.On("UpdateList", orderoutput).Return(nil)
	orderlist, err := uslist.UpdateUserListstat(usrin, usrstatus)
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
}
