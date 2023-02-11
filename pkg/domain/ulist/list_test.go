package ulist

import (
	"testing"

	usermock "github.com/Deepak/pkg/storage/db/user/mocks"
	"github.com/Deepak/pkg/storage/db/userlist"
	ulistmock "github.com/Deepak/pkg/storage/db/userlist/mocks"
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
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
		Status:        "pending",
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
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
		Status:        "pending",
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
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
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
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
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
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
		Status:        "pending",
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
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
		Status:        "pending",
	}
	orderoutput := &userlist.Order{
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
		Status:        "approved",
	}
	myuliststorage.On("GetOrderByStringiId", usrin).Return(orderinput, nil)
	myuliststorage.On("UpdateList", orderoutput).Return(nil)
	orderlist, err := uslist.UpdateUserListstat(usrin, usrstatus)
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
}

func TestUpdateUserList(t *testing.T) {
	myuliststorage := ulistmock.NewRepository(t)
	myuserstorage := usermock.NewRepository(t)
	input := Input{
		User:     myuserstorage,
		UserList: myuliststorage,
	}
	uslist, _ := Init(input)
	usrin := 123
	orderinput := &userlist.OrderUpdate{
		Id:     123,
		UserID: "b72037f5-ef12-466b-a9cc-6a8b57dc8c02",
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
		},
		EmpName:  "Deepak",
		EmpEmail: "Deepak@blackwoodseven.com",
	}
	orderoutput := &userlist.Order{
		UserID: nil,
		Items: []userlist.Item{
			{
				ItemID:   "3",
				Quantity: "2",
			},
			{
				ItemID:   "5",
				Quantity: "4",
			},
			{
				ItemID:   "25",
				Quantity: "6",
			},
		},
		EmpName:       "Deepak",
		EmpEmail:      "Deepak@blackwoodseven.com",
		RequestedDate: "15/02/2023",
		DueDate:       "15/02/2023",
		Status:        "approved",
	}
	myuliststorage.On("GetOrderByStringiId", usrin).Return(orderoutput, nil)
	myuliststorage.On("UpdateList", orderoutput).Return(nil)
	orderlist, err := uslist.UpdateUserList(orderinput)
	assert.Nil(t, err)
	assert.NotNil(t, orderlist)
	myuliststorage.AssertExpectations(t)
}
