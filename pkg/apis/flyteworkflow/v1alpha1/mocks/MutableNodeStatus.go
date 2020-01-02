// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/lyft/flytestdlib/storage"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// MutableNodeStatus is an autogenerated mock type for the MutableNodeStatus type
type MutableNodeStatus struct {
	mock.Mock
}

// ClearDynamicNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearDynamicNodeStatus() {
	_m.Called()
}

// ClearLastAttemptStartedAt provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearLastAttemptStartedAt() {
	_m.Called()
}

// ClearTaskStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearTaskStatus() {
	_m.Called()
}

// ClearWorkflowStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearWorkflowStatus() {
	_m.Called()
}

type MutableNodeStatus_GetBranchStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetBranchStatus) Return(_a0 v1alpha1.MutableBranchNodeStatus) *MutableNodeStatus_GetBranchStatus {
	return &MutableNodeStatus_GetBranchStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetBranchStatus() *MutableNodeStatus_GetBranchStatus {
	c := _m.On("GetBranchStatus")
	return &MutableNodeStatus_GetBranchStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetBranchStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetBranchStatus {
	c := _m.On("GetBranchStatus", matchers...)
	return &MutableNodeStatus_GetBranchStatus{Call: c}
}

// GetBranchStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetDynamicNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetDynamicNodeStatus) Return(_a0 v1alpha1.MutableDynamicNodeStatus) *MutableNodeStatus_GetDynamicNodeStatus {
	return &MutableNodeStatus_GetDynamicNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetDynamicNodeStatus() *MutableNodeStatus_GetDynamicNodeStatus {
	c := _m.On("GetDynamicNodeStatus")
	return &MutableNodeStatus_GetDynamicNodeStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetDynamicNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetDynamicNodeStatus {
	c := _m.On("GetDynamicNodeStatus", matchers...)
	return &MutableNodeStatus_GetDynamicNodeStatus{Call: c}
}

// GetDynamicNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateBranchStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateBranchStatus) Return(_a0 v1alpha1.MutableBranchNodeStatus) *MutableNodeStatus_GetOrCreateBranchStatus {
	return &MutableNodeStatus_GetOrCreateBranchStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateBranchStatus() *MutableNodeStatus_GetOrCreateBranchStatus {
	c := _m.On("GetOrCreateBranchStatus")
	return &MutableNodeStatus_GetOrCreateBranchStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetOrCreateBranchStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateBranchStatus {
	c := _m.On("GetOrCreateBranchStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateBranchStatus{Call: c}
}

// GetOrCreateBranchStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateDynamicNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateDynamicNodeStatus) Return(_a0 v1alpha1.MutableDynamicNodeStatus) *MutableNodeStatus_GetOrCreateDynamicNodeStatus {
	return &MutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateDynamicNodeStatus() *MutableNodeStatus_GetOrCreateDynamicNodeStatus {
	c := _m.On("GetOrCreateDynamicNodeStatus")
	return &MutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetOrCreateDynamicNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateDynamicNodeStatus {
	c := _m.On("GetOrCreateDynamicNodeStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: c}
}

// GetOrCreateDynamicNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateTaskStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateTaskStatus) Return(_a0 v1alpha1.MutableTaskNodeStatus) *MutableNodeStatus_GetOrCreateTaskStatus {
	return &MutableNodeStatus_GetOrCreateTaskStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateTaskStatus() *MutableNodeStatus_GetOrCreateTaskStatus {
	c := _m.On("GetOrCreateTaskStatus")
	return &MutableNodeStatus_GetOrCreateTaskStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetOrCreateTaskStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateTaskStatus {
	c := _m.On("GetOrCreateTaskStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateTaskStatus{Call: c}
}

// GetOrCreateTaskStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateWorkflowStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateWorkflowStatus) Return(_a0 v1alpha1.MutableWorkflowNodeStatus) *MutableNodeStatus_GetOrCreateWorkflowStatus {
	return &MutableNodeStatus_GetOrCreateWorkflowStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateWorkflowStatus() *MutableNodeStatus_GetOrCreateWorkflowStatus {
	c := _m.On("GetOrCreateWorkflowStatus")
	return &MutableNodeStatus_GetOrCreateWorkflowStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetOrCreateWorkflowStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateWorkflowStatus {
	c := _m.On("GetOrCreateWorkflowStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateWorkflowStatus{Call: c}
}

// GetOrCreateWorkflowStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetTaskStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetTaskStatus) Return(_a0 v1alpha1.MutableTaskNodeStatus) *MutableNodeStatus_GetTaskStatus {
	return &MutableNodeStatus_GetTaskStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetTaskStatus() *MutableNodeStatus_GetTaskStatus {
	c := _m.On("GetTaskStatus")
	return &MutableNodeStatus_GetTaskStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetTaskStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetTaskStatus {
	c := _m.On("GetTaskStatus", matchers...)
	return &MutableNodeStatus_GetTaskStatus{Call: c}
}

// GetTaskStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetWorkflowStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetWorkflowStatus) Return(_a0 v1alpha1.MutableWorkflowNodeStatus) *MutableNodeStatus_GetWorkflowStatus {
	return &MutableNodeStatus_GetWorkflowStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetWorkflowStatus() *MutableNodeStatus_GetWorkflowStatus {
	c := _m.On("GetWorkflowStatus")
	return &MutableNodeStatus_GetWorkflowStatus{Call: c}
}

func (_m *MutableNodeStatus) OnGetWorkflowStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetWorkflowStatus {
	c := _m.On("GetWorkflowStatus", matchers...)
	return &MutableNodeStatus_GetWorkflowStatus{Call: c}
}

// GetWorkflowStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_IncrementAttempts struct {
	*mock.Call
}

func (_m MutableNodeStatus_IncrementAttempts) Return(_a0 uint32) *MutableNodeStatus_IncrementAttempts {
	return &MutableNodeStatus_IncrementAttempts{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnIncrementAttempts() *MutableNodeStatus_IncrementAttempts {
	c := _m.On("IncrementAttempts")
	return &MutableNodeStatus_IncrementAttempts{Call: c}
}

func (_m *MutableNodeStatus) OnIncrementAttemptsMatch(matchers ...interface{}) *MutableNodeStatus_IncrementAttempts {
	c := _m.On("IncrementAttempts", matchers...)
	return &MutableNodeStatus_IncrementAttempts{Call: c}
}

// IncrementAttempts provides a mock function with given fields:
func (_m *MutableNodeStatus) IncrementAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// ResetDirty provides a mock function with given fields:
func (_m *MutableNodeStatus) ResetDirty() {
	_m.Called()
}

// SetCached provides a mock function with given fields:
func (_m *MutableNodeStatus) SetCached() {
	_m.Called()
}

// SetDataDir provides a mock function with given fields: _a0
func (_m *MutableNodeStatus) SetDataDir(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// SetOutputDir provides a mock function with given fields: d
func (_m *MutableNodeStatus) SetOutputDir(d storage.DataReference) {
	_m.Called(d)
}

// SetParentNodeID provides a mock function with given fields: n
func (_m *MutableNodeStatus) SetParentNodeID(n *string) {
	_m.Called(n)
}

// SetParentTaskID provides a mock function with given fields: t
func (_m *MutableNodeStatus) SetParentTaskID(t *core.TaskExecutionIdentifier) {
	_m.Called(t)
}

// UpdatePhase provides a mock function with given fields: phase, occurredAt, reason
func (_m *MutableNodeStatus) UpdatePhase(phase v1alpha1.NodePhase, occurredAt v1.Time, reason string) {
	_m.Called(phase, occurredAt, reason)
}
