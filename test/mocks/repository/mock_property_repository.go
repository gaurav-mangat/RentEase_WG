// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\gmangat\RentEase\internal\domain\interfaces\property_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	entities "rentease/internal/domain/entities"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockPropertyRepo is a mock of PropertyRepo interface.
type MockPropertyRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPropertyRepoMockRecorder
}

// MockPropertyRepoMockRecorder is the mock recorder for MockPropertyRepo.
type MockPropertyRepoMockRecorder struct {
	mock *MockPropertyRepo
}

// NewMockPropertyRepo creates a new mock instance.
func NewMockPropertyRepo(ctrl *gomock.Controller) *MockPropertyRepo {
	mock := &MockPropertyRepo{ctrl: ctrl}
	mock.recorder = &MockPropertyRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPropertyRepo) EXPECT() *MockPropertyRepoMockRecorder {
	return m.recorder
}

// DeleteAllListedPropertiesOfaUser mocks base method.
func (m *MockPropertyRepo) DeleteAllListedPropertiesOfaUser(username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllListedPropertiesOfaUser", username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllListedPropertiesOfaUser indicates an expected call of DeleteAllListedPropertiesOfaUser.
func (mr *MockPropertyRepoMockRecorder) DeleteAllListedPropertiesOfaUser(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllListedPropertiesOfaUser", reflect.TypeOf((*MockPropertyRepo)(nil).DeleteAllListedPropertiesOfaUser), username)
}

// DeleteListedProperty mocks base method.
func (m *MockPropertyRepo) DeleteListedProperty(propertyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListedProperty", propertyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteListedProperty indicates an expected call of DeleteListedProperty.
func (mr *MockPropertyRepoMockRecorder) DeleteListedProperty(propertyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListedProperty", reflect.TypeOf((*MockPropertyRepo)(nil).DeleteListedProperty), propertyID)
}

// FindByID mocks base method.
func (m *MockPropertyRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*entities.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*entities.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPropertyRepoMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPropertyRepo)(nil).FindByID), ctx, id)
}

// FindPendingProperties mocks base method.
func (m *MockPropertyRepo) FindPendingProperties() ([]entities.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPendingProperties")
	ret0, _ := ret[0].([]entities.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPendingProperties indicates an expected call of FindPendingProperties.
func (mr *MockPropertyRepoMockRecorder) FindPendingProperties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPendingProperties", reflect.TypeOf((*MockPropertyRepo)(nil).FindPendingProperties))
}

// GetAllListedProperties mocks base method.
func (m *MockPropertyRepo) GetAllListedProperties(activerUseronly bool) ([]entities.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllListedProperties", activerUseronly)
	ret0, _ := ret[0].([]entities.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllListedProperties indicates an expected call of GetAllListedProperties.
func (mr *MockPropertyRepoMockRecorder) GetAllListedProperties(activerUseronly interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllListedProperties", reflect.TypeOf((*MockPropertyRepo)(nil).GetAllListedProperties), activerUseronly)
}

// SaveProperty mocks base method.
func (m *MockPropertyRepo) SaveProperty(property entities.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProperty", property)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveProperty indicates an expected call of SaveProperty.
func (mr *MockPropertyRepoMockRecorder) SaveProperty(property interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProperty", reflect.TypeOf((*MockPropertyRepo)(nil).SaveProperty), property)
}

// UpdateApprovalStatus mocks base method.
func (m *MockPropertyRepo) UpdateApprovalStatus(propertyID primitive.ObjectID, approved bool, adminUsername string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApprovalStatus", propertyID, approved, adminUsername)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApprovalStatus indicates an expected call of UpdateApprovalStatus.
func (mr *MockPropertyRepoMockRecorder) UpdateApprovalStatus(propertyID, approved, adminUsername interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApprovalStatus", reflect.TypeOf((*MockPropertyRepo)(nil).UpdateApprovalStatus), propertyID, approved, adminUsername)
}

// UpdateListedProperty mocks base method.
func (m *MockPropertyRepo) UpdateListedProperty(property entities.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateListedProperty", property)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateListedProperty indicates an expected call of UpdateListedProperty.
func (mr *MockPropertyRepoMockRecorder) UpdateListedProperty(property interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateListedProperty", reflect.TypeOf((*MockPropertyRepo)(nil).UpdateListedProperty), property)
}
