// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package pathlib

import mock "github.com/stretchr/testify/mock"

// mockNamer is an autogenerated mock type for the namer type
type mockNamer struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *mockNamer) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
