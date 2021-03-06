// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import internal "github.com/kyma-project/kyma/components/application-broker/internal"

import mock "github.com/stretchr/testify/mock"

// AppGetter is an autogenerated mock type for the AppGetter type
type AppGetter struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *AppGetter) Get(_a0 internal.ApplicationName) (*internal.Application, error) {
	ret := _m.Called(_a0)

	var r0 *internal.Application
	if rf, ok := ret.Get(0).(func(internal.ApplicationName) *internal.Application); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.ApplicationName) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
