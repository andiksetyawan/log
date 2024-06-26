// Code generated by mockery v2.42.1. DO NOT EDIT.

package logmock

import (
	slog "github.com/andiksetyawan/log/slog"
	mock "github.com/stretchr/testify/mock"
)

// OptFunc is an autogenerated mock type for the OptFunc type
type OptFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *OptFunc) Execute(_a0 *slog.Log) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*slog.Log) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOptFunc creates a new instance of OptFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOptFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *OptFunc {
	mock := &OptFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
