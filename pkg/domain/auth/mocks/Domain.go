// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	auth "github.com/Deepak/pkg/domain/auth"
	mock "github.com/stretchr/testify/mock"

	user "github.com/Deepak/pkg/storage/db/user"
)

// Domain is an autogenerated mock type for the Domain type
type Domain struct {
	mock.Mock
}

// AdminAccess provides a mock function with given fields: userId
func (_m *Domain) AdminAccess(userId string) (*user.User, error) {
	ret := _m.Called(userId)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields:
func (_m *Domain) GetAllUsers() ([]*user.User, error) {
	ret := _m.Called()

	var r0 []*user.User
	if rf, ok := ret.Get(0).(func() []*user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRole provides a mock function with given fields: userId
func (_m *Domain) GetUserRole(userId string) (string, error) {
	ret := _m.Called(userId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: req
func (_m *Domain) LoginUser(req auth.LoginRequest) (*auth.LoginRespose, error) {
	ret := _m.Called(req)

	var r0 *auth.LoginRespose
	if rf, ok := ret.Get(0).(func(auth.LoginRequest) *auth.LoginRespose); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.LoginRespose)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(auth.LoginRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUser provides a mock function with given fields: userId
func (_m *Domain) RemoveUser(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scheduler provides a mock function with given fields:
func (_m *Domain) Scheduler() {
	_m.Called()
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
