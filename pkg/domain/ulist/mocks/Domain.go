// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	userlist "github.com/Deepak/pkg/storage/db/userlist"
)

// Domain is an autogenerated mock type for the Domain type
type Domain struct {
	mock.Mock
}

// CreateUserList provides a mock function with given fields: c
func (_m *Domain) CreateUserList(c *userlist.Order) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*userlist.Order) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllApprovedUserLists provides a mock function with given fields:
func (_m *Domain) GetAllApprovedUserLists() ([]*userlist.Order, error) {
	ret := _m.Called()

	var r0 []*userlist.Order
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*userlist.Order, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*userlist.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userlist.Order)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllNotApprovedUserLists provides a mock function with given fields:
func (_m *Domain) GetAllNotApprovedUserLists() ([]*userlist.Order, error) {
	ret := _m.Called()

	var r0 []*userlist.Order
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*userlist.Order, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*userlist.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userlist.Order)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUserLists provides a mock function with given fields:
func (_m *Domain) GetAllUserLists() ([]*userlist.Order, error) {
	ret := _m.Called()

	var r0 []*userlist.Order
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*userlist.Order, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*userlist.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userlist.Order)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserListStatus provides a mock function with given fields: id
func (_m *Domain) GetUserListStatus(id int) (string, error) {
	ret := _m.Called(id)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserLists provides a mock function with given fields: id
func (_m *Domain) GetUserLists(id string) ([]*userlist.Order, error) {
	ret := _m.Called(id)

	var r0 []*userlist.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*userlist.Order, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) []*userlist.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userlist.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRemainder provides a mock function with given fields:
func (_m *Domain) SendRemainder() (map[string]interface{}, error) {
	ret := _m.Called()

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserList provides a mock function with given fields: c
func (_m *Domain) UpdateUserList(c *userlist.OrderUpdate) (*userlist.Order, error) {
	ret := _m.Called(c)

	var r0 *userlist.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(*userlist.OrderUpdate) (*userlist.Order, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*userlist.OrderUpdate) *userlist.Order); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userlist.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(*userlist.OrderUpdate) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserListstat provides a mock function with given fields: id, status
func (_m *Domain) UpdateUserListstat(id int, status string) (*userlist.Order, error) {
	ret := _m.Called(id, status)

	var r0 *userlist.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*userlist.Order, error)); ok {
		return rf(id, status)
	}
	if rf, ok := ret.Get(0).(func(int, string) *userlist.Order); ok {
		r0 = rf(id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userlist.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDomain interface {
	mock.TestingT
	Cleanup(func())
}

// NewDomain creates a new instance of Domain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDomain(t mockConstructorTestingTNewDomain) *Domain {
	mock := &Domain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
