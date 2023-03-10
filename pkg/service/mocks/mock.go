// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	notes_app "github.com/Liopun/notes-app"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user notes_app.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockNotesList is a mock of NotesList interface.
type MockNotesList struct {
	ctrl     *gomock.Controller
	recorder *MockNotesListMockRecorder
}

// MockNotesListMockRecorder is the mock recorder for MockNotesList.
type MockNotesListMockRecorder struct {
	mock *MockNotesList
}

// NewMockNotesList creates a new mock instance.
func NewMockNotesList(ctrl *gomock.Controller) *MockNotesList {
	mock := &MockNotesList{ctrl: ctrl}
	mock.recorder = &MockNotesListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotesList) EXPECT() *MockNotesListMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNotesList) Create(userId int, list notes_app.NotesList) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, list)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockNotesListMockRecorder) Create(userId, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNotesList)(nil).Create), userId, list)
}

// Delete mocks base method.
func (m *MockNotesList) Delete(userId, listId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, listId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNotesListMockRecorder) Delete(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNotesList)(nil).Delete), userId, listId)
}

// GetAll mocks base method.
func (m *MockNotesList) GetAll(userId int) ([]notes_app.NotesList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]notes_app.NotesList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNotesListMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNotesList)(nil).GetAll), userId)
}

// GetById mocks base method.
func (m *MockNotesList) GetById(userId, listId int) (notes_app.NotesList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, listId)
	ret0, _ := ret[0].(notes_app.NotesList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockNotesListMockRecorder) GetById(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockNotesList)(nil).GetById), userId, listId)
}

// Update mocks base method.
func (m *MockNotesList) Update(userId, listId int, inp notes_app.UpdateListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, listId, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNotesListMockRecorder) Update(userId, listId, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNotesList)(nil).Update), userId, listId, inp)
}

// MockNotesItem is a mock of NotesItem interface.
type MockNotesItem struct {
	ctrl     *gomock.Controller
	recorder *MockNotesItemMockRecorder
}

// MockNotesItemMockRecorder is the mock recorder for MockNotesItem.
type MockNotesItemMockRecorder struct {
	mock *MockNotesItem
}

// NewMockNotesItem creates a new mock instance.
func NewMockNotesItem(ctrl *gomock.Controller) *MockNotesItem {
	mock := &MockNotesItem{ctrl: ctrl}
	mock.recorder = &MockNotesItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotesItem) EXPECT() *MockNotesItemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNotesItem) Create(userId, listId int, item notes_app.NotesItem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, listId, item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockNotesItemMockRecorder) Create(userId, listId, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNotesItem)(nil).Create), userId, listId, item)
}

// Delete mocks base method.
func (m *MockNotesItem) Delete(userId, itemId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, itemId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNotesItemMockRecorder) Delete(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNotesItem)(nil).Delete), userId, itemId)
}

// GetAll mocks base method.
func (m *MockNotesItem) GetAll(userId, listId int) ([]notes_app.NotesItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId, listId)
	ret0, _ := ret[0].([]notes_app.NotesItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNotesItemMockRecorder) GetAll(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNotesItem)(nil).GetAll), userId, listId)
}

// GetById mocks base method.
func (m *MockNotesItem) GetById(userId, itemId int) (notes_app.NotesItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, itemId)
	ret0, _ := ret[0].(notes_app.NotesItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockNotesItemMockRecorder) GetById(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockNotesItem)(nil).GetById), userId, itemId)
}

// Update mocks base method.
func (m *MockNotesItem) Update(userId, itemId int, inp notes_app.UpdateItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, itemId, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNotesItemMockRecorder) Update(userId, itemId, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNotesItem)(nil).Update), userId, itemId, inp)
}
