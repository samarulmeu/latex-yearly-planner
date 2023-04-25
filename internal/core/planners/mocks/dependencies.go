// Code generated by MockGen. DO NOT EDIT.
// Source: dependencies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	planners "github.com/kudrykv/latex-yearly-planner/internal/core/planners"
	entities "github.com/kudrykv/latex-yearly-planner/internal/core/planners/entities"
)

// MockPlannerBuilder is a mock of PlannerBuilder interface.
type MockPlannerBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockPlannerBuilderMockRecorder
}

// MockPlannerBuilderMockRecorder is the mock recorder for MockPlannerBuilder.
type MockPlannerBuilderMockRecorder struct {
	mock *MockPlannerBuilder
}

// NewMockPlannerBuilder creates a new mock instance.
func NewMockPlannerBuilder(ctrl *gomock.Controller) *MockPlannerBuilder {
	mock := &MockPlannerBuilder{ctrl: ctrl}
	mock.recorder = &MockPlannerBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlannerBuilder) EXPECT() *MockPlannerBuilderMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockPlannerBuilder) Generate(arg0 context.Context) (entities.NoteStructure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(entities.NoteStructure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockPlannerBuilderMockRecorder) Generate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockPlannerBuilder)(nil).Generate), arg0)
}

// MockNoteWriter is a mock of NoteWriter interface.
type MockNoteWriter struct {
	ctrl     *gomock.Controller
	recorder *MockNoteWriterMockRecorder
}

// MockNoteWriterMockRecorder is the mock recorder for MockNoteWriter.
type MockNoteWriterMockRecorder struct {
	mock *MockNoteWriter
}

// NewMockNoteWriter creates a new mock instance.
func NewMockNoteWriter(ctrl *gomock.Controller) *MockNoteWriter {
	mock := &MockNoteWriter{ctrl: ctrl}
	mock.recorder = &MockNoteWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoteWriter) EXPECT() *MockNoteWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockNoteWriter) Write(arg0 context.Context, arg1 entities.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockNoteWriterMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockNoteWriter)(nil).Write), arg0, arg1)
}

// MockCommander is a mock of Commander interface.
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander.
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance.
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockCommander) Run(arg0 context.Context, arg1 planners.CommandName, arg2 ...planners.StringArg) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockCommanderMockRecorder) Run(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCommander)(nil).Run), varargs...)
}