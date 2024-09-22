// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NotfiService is an autogenerated mock type for the NotfiService type
type NotfiService struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: message
func (_m *NotfiService) SendEmail(message string) (bool, error) {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for SendEmail")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(message)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNotfiService creates a new instance of NotfiService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotfiService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotfiService {
	mock := &NotfiService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
