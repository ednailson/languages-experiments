// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package gomock is a generated GoMock package.
package gomock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStore) Close(at time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", at)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStoreMockRecorder) Close(at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close), at)
}

// Get mocks base method.
func (m *MockStore) Get(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), name)
}

// Open mocks base method.
func (m *MockStore) Open(at time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", at)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockStoreMockRecorder) Open(at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockStore)(nil).Open), at)
}

// Sell mocks base method.
func (m *MockStore) Sell(id string, quantity int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sell", id, quantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sell indicates an expected call of Sell.
func (mr *MockStoreMockRecorder) Sell(id, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sell", reflect.TypeOf((*MockStore)(nil).Sell), id, quantity)
}
