// Code generated by mockery v2.43.2. DO NOT EDIT.

package testing

import (
	context "context"

	sensor "github.com/joshuar/go-hass-agent/internal/hass/sensor"
	mock "github.com/stretchr/testify/mock"
)

// MockWorker is an autogenerated mock type for the Worker type
type MockWorker struct {
	mock.Mock
}

type MockWorker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorker) EXPECT() *MockWorker_Expecter {
	return &MockWorker_Expecter{mock: &_m.Mock}
}

// ID provides a mock function with given fields:
func (_m *MockWorker) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWorker_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type MockWorker_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *MockWorker_Expecter) ID() *MockWorker_ID_Call {
	return &MockWorker_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *MockWorker_ID_Call) Run(run func()) *MockWorker_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorker_ID_Call) Return(_a0 string) *MockWorker_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorker_ID_Call) RunAndReturn(run func() string) *MockWorker_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Sensors provides a mock function with given fields: ctx
func (_m *MockWorker) Sensors(ctx context.Context) ([]sensor.Details, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Sensors")
	}

	var r0 []sensor.Details
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]sensor.Details, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []sensor.Details); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sensor.Details)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorker_Sensors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sensors'
type MockWorker_Sensors_Call struct {
	*mock.Call
}

// Sensors is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockWorker_Expecter) Sensors(ctx interface{}) *MockWorker_Sensors_Call {
	return &MockWorker_Sensors_Call{Call: _e.mock.On("Sensors", ctx)}
}

func (_c *MockWorker_Sensors_Call) Run(run func(ctx context.Context)) *MockWorker_Sensors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockWorker_Sensors_Call) Return(_a0 []sensor.Details, _a1 error) *MockWorker_Sensors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorker_Sensors_Call) RunAndReturn(run func(context.Context) ([]sensor.Details, error)) *MockWorker_Sensors_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockWorker) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorker_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockWorker_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockWorker_Expecter) Stop() *MockWorker_Stop_Call {
	return &MockWorker_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockWorker_Stop_Call) Run(run func()) *MockWorker_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorker_Stop_Call) Return(_a0 error) *MockWorker_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorker_Stop_Call) RunAndReturn(run func() error) *MockWorker_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Updates provides a mock function with given fields: ctx
func (_m *MockWorker) Updates(ctx context.Context) (<-chan sensor.Details, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Updates")
	}

	var r0 <-chan sensor.Details
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan sensor.Details, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan sensor.Details); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan sensor.Details)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorker_Updates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Updates'
type MockWorker_Updates_Call struct {
	*mock.Call
}

// Updates is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockWorker_Expecter) Updates(ctx interface{}) *MockWorker_Updates_Call {
	return &MockWorker_Updates_Call{Call: _e.mock.On("Updates", ctx)}
}

func (_c *MockWorker_Updates_Call) Run(run func(ctx context.Context)) *MockWorker_Updates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockWorker_Updates_Call) Return(_a0 <-chan sensor.Details, _a1 error) *MockWorker_Updates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorker_Updates_Call) RunAndReturn(run func(context.Context) (<-chan sensor.Details, error)) *MockWorker_Updates_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorker creates a new instance of MockWorker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorker {
	mock := &MockWorker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
