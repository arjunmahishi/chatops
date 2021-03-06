// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	exoctx "bitbucket.org/exotel/exotel_code/commonix/lib/exoctx/go"
	"github.com/arjunmahishi/Chatops/payload"
	mock "github.com/stretchr/testify/mock"
)

// Commander is an autogenerated mock type for the Commander type
type Commander struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *Commander) Execute(_a0 exoctx.Ctx, _a1 payload.Handler) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(exoctx.Ctx, payload.Handler) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(exoctx.Ctx, payload.Handler) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCatagory provides a mock function with given fields:
func (_m *Commander) GetCatagory() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *Commander) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MatchCommand provides a mock function with given fields: _a0
func (_m *Commander) MatchCommand(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
