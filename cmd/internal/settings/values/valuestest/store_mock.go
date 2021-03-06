// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package valuestest is a generated GoMock package.
package valuestest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetString mocks base method
func (m *MockStore) GetString(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString
func (mr *MockStoreMockRecorder) GetString(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockStore)(nil).GetString), key)
}

// GetStringSlice mocks base method
func (m *MockStore) GetStringSlice(key string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringSlice", key)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStringSlice indicates an expected call of GetStringSlice
func (mr *MockStoreMockRecorder) GetStringSlice(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringSlice", reflect.TypeOf((*MockStore)(nil).GetStringSlice), key)
}

// GetInt mocks base method
func (m *MockStore) GetInt(key string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", key)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt
func (mr *MockStoreMockRecorder) GetInt(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockStore)(nil).GetInt), key)
}

// GetBool mocks base method
func (m *MockStore) GetBool(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBool", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool
func (mr *MockStoreMockRecorder) GetBool(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockStore)(nil).GetBool), key)
}

// IsSet mocks base method
func (m *MockStore) IsSet(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSet", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSet indicates an expected call of IsSet
func (mr *MockStoreMockRecorder) IsSet(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSet", reflect.TypeOf((*MockStore)(nil).IsSet), key)
}

// Set mocks base method
func (m *MockStore) Set(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set
func (mr *MockStoreMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStore)(nil).Set), key, value)
}
