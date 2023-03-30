// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by mockery v2.20.0. DO NOT EDIT.

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package mocks

import (
	context "context"

	component "github.com/elastic/elastic-agent/pkg/component"

	mock "github.com/stretchr/testify/mock"

	runtime "github.com/elastic/elastic-agent/pkg/component/runtime"
)

// RuntimeManager is an autogenerated mock type for the RuntimeManager type
type RuntimeManager struct {
	mock.Mock
}

type RuntimeManager_Expecter struct {
	mock *mock.Mock
}

func (_m *RuntimeManager) EXPECT() *RuntimeManager_Expecter {
	return &RuntimeManager_Expecter{mock: &_m.Mock}
}

// Errors provides a mock function with given fields:
func (_m *RuntimeManager) Errors() <-chan error {
	ret := _m.Called()

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func() <-chan error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

// RuntimeManager_Errors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errors'
type RuntimeManager_Errors_Call struct {
	*mock.Call
}

// Errors is a helper method to define mock.On call
func (_e *RuntimeManager_Expecter) Errors() *RuntimeManager_Errors_Call {
	return &RuntimeManager_Errors_Call{Call: _e.mock.On("Errors")}
}

func (_c *RuntimeManager_Errors_Call) Run(run func()) *RuntimeManager_Errors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuntimeManager_Errors_Call) Return(_a0 <-chan error) *RuntimeManager_Errors_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeManager_Errors_Call) RunAndReturn(run func() <-chan error) *RuntimeManager_Errors_Call {
	_c.Call.Return(run)
	return _c
}

// PerformAction provides a mock function with given fields: ctx, comp, unit, name, params
func (_m *RuntimeManager) PerformAction(ctx context.Context, comp component.Component, unit component.Unit, name string, params map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(ctx, comp, unit, name, params)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, component.Component, component.Unit, string, map[string]interface{}) (map[string]interface{}, error)); ok {
		return rf(ctx, comp, unit, name, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, component.Component, component.Unit, string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(ctx, comp, unit, name, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, component.Component, component.Unit, string, map[string]interface{}) error); ok {
		r1 = rf(ctx, comp, unit, name, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuntimeManager_PerformAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PerformAction'
type RuntimeManager_PerformAction_Call struct {
	*mock.Call
}

// PerformAction is a helper method to define mock.On call
//   - ctx context.Context
//   - comp component.Component
//   - unit component.Unit
//   - name string
//   - params map[string]interface{}
func (_e *RuntimeManager_Expecter) PerformAction(ctx interface{}, comp interface{}, unit interface{}, name interface{}, params interface{}) *RuntimeManager_PerformAction_Call {
	return &RuntimeManager_PerformAction_Call{Call: _e.mock.On("PerformAction", ctx, comp, unit, name, params)}
}

func (_c *RuntimeManager_PerformAction_Call) Run(run func(ctx context.Context, comp component.Component, unit component.Unit, name string, params map[string]interface{})) *RuntimeManager_PerformAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(component.Component), args[2].(component.Unit), args[3].(string), args[4].(map[string]interface{}))
	})
	return _c
}

func (_c *RuntimeManager_PerformAction_Call) Return(_a0 map[string]interface{}, _a1 error) *RuntimeManager_PerformAction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RuntimeManager_PerformAction_Call) RunAndReturn(run func(context.Context, component.Component, component.Unit, string, map[string]interface{}) (map[string]interface{}, error)) *RuntimeManager_PerformAction_Call {
	_c.Call.Return(run)
	return _c
}

// PerformDiagnostics provides a mock function with given fields: _a0, _a1
func (_m *RuntimeManager) PerformDiagnostics(_a0 context.Context, _a1 ...runtime.ComponentUnitDiagnosticRequest) []runtime.ComponentUnitDiagnostic {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []runtime.ComponentUnitDiagnostic
	if rf, ok := ret.Get(0).(func(context.Context, ...runtime.ComponentUnitDiagnosticRequest) []runtime.ComponentUnitDiagnostic); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]runtime.ComponentUnitDiagnostic)
		}
	}

	return r0
}

// RuntimeManager_PerformDiagnostics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PerformDiagnostics'
type RuntimeManager_PerformDiagnostics_Call struct {
	*mock.Call
}

// PerformDiagnostics is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...runtime.ComponentUnitDiagnosticRequest
func (_e *RuntimeManager_Expecter) PerformDiagnostics(_a0 interface{}, _a1 ...interface{}) *RuntimeManager_PerformDiagnostics_Call {
	return &RuntimeManager_PerformDiagnostics_Call{Call: _e.mock.On("PerformDiagnostics",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *RuntimeManager_PerformDiagnostics_Call) Run(run func(_a0 context.Context, _a1 ...runtime.ComponentUnitDiagnosticRequest)) *RuntimeManager_PerformDiagnostics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]runtime.ComponentUnitDiagnosticRequest, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(runtime.ComponentUnitDiagnosticRequest)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *RuntimeManager_PerformDiagnostics_Call) Return(_a0 []runtime.ComponentUnitDiagnostic) *RuntimeManager_PerformDiagnostics_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeManager_PerformDiagnostics_Call) RunAndReturn(run func(context.Context, ...runtime.ComponentUnitDiagnosticRequest) []runtime.ComponentUnitDiagnostic) *RuntimeManager_PerformDiagnostics_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: _a0
func (_m *RuntimeManager) Run(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RuntimeManager_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type RuntimeManager_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *RuntimeManager_Expecter) Run(_a0 interface{}) *RuntimeManager_Run_Call {
	return &RuntimeManager_Run_Call{Call: _e.mock.On("Run", _a0)}
}

func (_c *RuntimeManager_Run_Call) Run(run func(_a0 context.Context)) *RuntimeManager_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RuntimeManager_Run_Call) Return(_a0 error) *RuntimeManager_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeManager_Run_Call) RunAndReturn(run func(context.Context) error) *RuntimeManager_Run_Call {
	_c.Call.Return(run)
	return _c
}

// State provides a mock function with given fields:
func (_m *RuntimeManager) State() []runtime.ComponentComponentState {
	ret := _m.Called()

	var r0 []runtime.ComponentComponentState
	if rf, ok := ret.Get(0).(func() []runtime.ComponentComponentState); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]runtime.ComponentComponentState)
		}
	}

	return r0
}

// RuntimeManager_State_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'State'
type RuntimeManager_State_Call struct {
	*mock.Call
}

// State is a helper method to define mock.On call
func (_e *RuntimeManager_Expecter) State() *RuntimeManager_State_Call {
	return &RuntimeManager_State_Call{Call: _e.mock.On("State")}
}

func (_c *RuntimeManager_State_Call) Run(run func()) *RuntimeManager_State_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuntimeManager_State_Call) Return(_a0 []runtime.ComponentComponentState) *RuntimeManager_State_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeManager_State_Call) RunAndReturn(run func() []runtime.ComponentComponentState) *RuntimeManager_State_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeAll provides a mock function with given fields: _a0
func (_m *RuntimeManager) SubscribeAll(_a0 context.Context) *runtime.SubscriptionAll {
	ret := _m.Called(_a0)

	var r0 *runtime.SubscriptionAll
	if rf, ok := ret.Get(0).(func(context.Context) *runtime.SubscriptionAll); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.SubscriptionAll)
		}
	}

	return r0
}

// RuntimeManager_SubscribeAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeAll'
type RuntimeManager_SubscribeAll_Call struct {
	*mock.Call
}

// SubscribeAll is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *RuntimeManager_Expecter) SubscribeAll(_a0 interface{}) *RuntimeManager_SubscribeAll_Call {
	return &RuntimeManager_SubscribeAll_Call{Call: _e.mock.On("SubscribeAll", _a0)}
}

func (_c *RuntimeManager_SubscribeAll_Call) Run(run func(_a0 context.Context)) *RuntimeManager_SubscribeAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RuntimeManager_SubscribeAll_Call) Return(_a0 *runtime.SubscriptionAll) *RuntimeManager_SubscribeAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeManager_SubscribeAll_Call) RunAndReturn(run func(context.Context) *runtime.SubscriptionAll) *RuntimeManager_SubscribeAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *RuntimeManager) Update(_a0 []component.Component) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]component.Component) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RuntimeManager_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type RuntimeManager_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 []component.Component
func (_e *RuntimeManager_Expecter) Update(_a0 interface{}) *RuntimeManager_Update_Call {
	return &RuntimeManager_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *RuntimeManager_Update_Call) Run(run func(_a0 []component.Component)) *RuntimeManager_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]component.Component))
	})
	return _c
}

func (_c *RuntimeManager_Update_Call) Return(_a0 error) *RuntimeManager_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeManager_Update_Call) RunAndReturn(run func([]component.Component) error) *RuntimeManager_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRuntimeManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewRuntimeManager creates a new instance of RuntimeManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRuntimeManager(t mockConstructorTestingTNewRuntimeManager) *RuntimeManager {
	mock := &RuntimeManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
