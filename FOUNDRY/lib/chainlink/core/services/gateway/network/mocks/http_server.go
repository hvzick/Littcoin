// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	network "github.com/smartcontractkit/chainlink/v2/core/services/gateway/network"
	mock "github.com/stretchr/testify/mock"
)

// HttpServer is an autogenerated mock type for the HttpServer type
type HttpServer struct {
	mock.Mock
}

type HttpServer_Expecter struct {
	mock *mock.Mock
}

func (_m *HttpServer) EXPECT() *HttpServer_Expecter {
	return &HttpServer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *HttpServer) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HttpServer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type HttpServer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *HttpServer_Expecter) Close() *HttpServer_Close_Call {
	return &HttpServer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *HttpServer_Close_Call) Run(run func()) *HttpServer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HttpServer_Close_Call) Return(_a0 error) *HttpServer_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HttpServer_Close_Call) RunAndReturn(run func() error) *HttpServer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetPort provides a mock function with no fields
func (_m *HttpServer) GetPort() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPort")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// HttpServer_GetPort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPort'
type HttpServer_GetPort_Call struct {
	*mock.Call
}

// GetPort is a helper method to define mock.On call
func (_e *HttpServer_Expecter) GetPort() *HttpServer_GetPort_Call {
	return &HttpServer_GetPort_Call{Call: _e.mock.On("GetPort")}
}

func (_c *HttpServer_GetPort_Call) Run(run func()) *HttpServer_GetPort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HttpServer_GetPort_Call) Return(_a0 int) *HttpServer_GetPort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HttpServer_GetPort_Call) RunAndReturn(run func() int) *HttpServer_GetPort_Call {
	_c.Call.Return(run)
	return _c
}

// SetHTTPRequestHandler provides a mock function with given fields: handler
func (_m *HttpServer) SetHTTPRequestHandler(handler network.HTTPRequestHandler) {
	_m.Called(handler)
}

// HttpServer_SetHTTPRequestHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHTTPRequestHandler'
type HttpServer_SetHTTPRequestHandler_Call struct {
	*mock.Call
}

// SetHTTPRequestHandler is a helper method to define mock.On call
//   - handler network.HTTPRequestHandler
func (_e *HttpServer_Expecter) SetHTTPRequestHandler(handler interface{}) *HttpServer_SetHTTPRequestHandler_Call {
	return &HttpServer_SetHTTPRequestHandler_Call{Call: _e.mock.On("SetHTTPRequestHandler", handler)}
}

func (_c *HttpServer_SetHTTPRequestHandler_Call) Run(run func(handler network.HTTPRequestHandler)) *HttpServer_SetHTTPRequestHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(network.HTTPRequestHandler))
	})
	return _c
}

func (_c *HttpServer_SetHTTPRequestHandler_Call) Return() *HttpServer_SetHTTPRequestHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *HttpServer_SetHTTPRequestHandler_Call) RunAndReturn(run func(network.HTTPRequestHandler)) *HttpServer_SetHTTPRequestHandler_Call {
	_c.Run(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *HttpServer) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HttpServer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type HttpServer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *HttpServer_Expecter) Start(_a0 interface{}) *HttpServer_Start_Call {
	return &HttpServer_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *HttpServer_Start_Call) Run(run func(_a0 context.Context)) *HttpServer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *HttpServer_Start_Call) Return(_a0 error) *HttpServer_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HttpServer_Start_Call) RunAndReturn(run func(context.Context) error) *HttpServer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewHttpServer creates a new instance of HttpServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHttpServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *HttpServer {
	mock := &HttpServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
