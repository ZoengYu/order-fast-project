// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ZoengYu/order-fast-project/db/sqlc (interfaces: DBService)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockDBService is a mock of DBService interface.
type MockDBService struct {
	ctrl     *gomock.Controller
	recorder *MockDBServiceMockRecorder
}

// MockDBServiceMockRecorder is the mock recorder for MockDBService.
type MockDBServiceMockRecorder struct {
	mock *MockDBService
}

// NewMockDBService creates a new mock instance.
func NewMockDBService(ctrl *gomock.Controller) *MockDBService {
	mock := &MockDBService{ctrl: ctrl}
	mock.recorder = &MockDBServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBService) EXPECT() *MockDBServiceMockRecorder {
	return m.recorder
}

// CreateMenuFood mocks base method.
func (m *MockDBService) CreateMenuFood(arg0 context.Context, arg1 db.CreateMenuFoodParams) (db.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenuFood", arg0, arg1)
	ret0, _ := ret[0].(db.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMenuFood indicates an expected call of CreateMenuFood.
func (mr *MockDBServiceMockRecorder) CreateMenuFood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuFood", reflect.TypeOf((*MockDBService)(nil).CreateMenuFood), arg0, arg1)
}

// CreateMenuFoodTag mocks base method.
func (m *MockDBService) CreateMenuFoodTag(arg0 context.Context, arg1 db.CreateMenuFoodTagParams) (db.FoodTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenuFoodTag", arg0, arg1)
	ret0, _ := ret[0].(db.FoodTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMenuFoodTag indicates an expected call of CreateMenuFoodTag.
func (mr *MockDBServiceMockRecorder) CreateMenuFoodTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuFoodTag", reflect.TypeOf((*MockDBService)(nil).CreateMenuFoodTag), arg0, arg1)
}

// CreateStore mocks base method.
func (m *MockDBService) CreateStore(arg0 context.Context, arg1 db.CreateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStore indicates an expected call of CreateStore.
func (mr *MockDBServiceMockRecorder) CreateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStore", reflect.TypeOf((*MockDBService)(nil).CreateStore), arg0, arg1)
}

// CreateStoreMenu mocks base method.
func (m *MockDBService) CreateStoreMenu(arg0 context.Context, arg1 db.CreateStoreMenuParams) (db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStoreMenu", arg0, arg1)
	ret0, _ := ret[0].(db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStoreMenu indicates an expected call of CreateStoreMenu.
func (mr *MockDBServiceMockRecorder) CreateStoreMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStoreMenu", reflect.TypeOf((*MockDBService)(nil).CreateStoreMenu), arg0, arg1)
}

// CreateTable mocks base method.
func (m *MockDBService) CreateTable(arg0 context.Context, arg1 db.CreateTableParams) (db.Table, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", arg0, arg1)
	ret0, _ := ret[0].(db.Table)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockDBServiceMockRecorder) CreateTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockDBService)(nil).CreateTable), arg0, arg1)
}

// DeleteMenu mocks base method.
func (m *MockDBService) DeleteMenu(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMenu", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMenu indicates an expected call of DeleteMenu.
func (mr *MockDBServiceMockRecorder) DeleteMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMenu", reflect.TypeOf((*MockDBService)(nil).DeleteMenu), arg0, arg1)
}

// DeleteStore mocks base method.
func (m *MockDBService) DeleteStore(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStore indicates an expected call of DeleteStore.
func (mr *MockDBServiceMockRecorder) DeleteStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStore", reflect.TypeOf((*MockDBService)(nil).DeleteStore), arg0, arg1)
}

// DeleteStoreTable mocks base method.
func (m *MockDBService) DeleteStoreTable(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStoreTable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStoreTable indicates an expected call of DeleteStoreTable.
func (mr *MockDBServiceMockRecorder) DeleteStoreTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStoreTable", reflect.TypeOf((*MockDBService)(nil).DeleteStoreTable), arg0, arg1)
}

// DeleteStoreTableByName mocks base method.
func (m *MockDBService) DeleteStoreTableByName(arg0 context.Context, arg1 db.DeleteStoreTableByNameParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStoreTableByName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStoreTableByName indicates an expected call of DeleteStoreTableByName.
func (mr *MockDBServiceMockRecorder) DeleteStoreTableByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStoreTableByName", reflect.TypeOf((*MockDBService)(nil).DeleteStoreTableByName), arg0, arg1)
}

// GetMenuFood mocks base method.
func (m *MockDBService) GetMenuFood(arg0 context.Context, arg1 db.GetMenuFoodParams) (db.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuFood", arg0, arg1)
	ret0, _ := ret[0].(db.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuFood indicates an expected call of GetMenuFood.
func (mr *MockDBServiceMockRecorder) GetMenuFood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuFood", reflect.TypeOf((*MockDBService)(nil).GetMenuFood), arg0, arg1)
}

// GetMenuFoodTag mocks base method.
func (m *MockDBService) GetMenuFoodTag(arg0 context.Context, arg1 db.GetMenuFoodTagParams) (db.FoodTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuFoodTag", arg0, arg1)
	ret0, _ := ret[0].(db.FoodTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuFoodTag indicates an expected call of GetMenuFoodTag.
func (mr *MockDBServiceMockRecorder) GetMenuFoodTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuFoodTag", reflect.TypeOf((*MockDBService)(nil).GetMenuFoodTag), arg0, arg1)
}

// GetStore mocks base method.
func (m *MockDBService) GetStore(arg0 context.Context, arg1 int64) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore.
func (mr *MockDBServiceMockRecorder) GetStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockDBService)(nil).GetStore), arg0, arg1)
}

// GetStoreByName mocks base method.
func (m *MockDBService) GetStoreByName(arg0 context.Context, arg1 string) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreByName", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreByName indicates an expected call of GetStoreByName.
func (mr *MockDBServiceMockRecorder) GetStoreByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreByName", reflect.TypeOf((*MockDBService)(nil).GetStoreByName), arg0, arg1)
}

// GetStoreMenu mocks base method.
func (m *MockDBService) GetStoreMenu(arg0 context.Context, arg1 db.GetStoreMenuParams) (db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreMenu", arg0, arg1)
	ret0, _ := ret[0].(db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreMenu indicates an expected call of GetStoreMenu.
func (mr *MockDBServiceMockRecorder) GetStoreMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreMenu", reflect.TypeOf((*MockDBService)(nil).GetStoreMenu), arg0, arg1)
}

// GetStoreTable mocks base method.
func (m *MockDBService) GetStoreTable(arg0 context.Context, arg1 db.GetStoreTableParams) (db.Table, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreTable", arg0, arg1)
	ret0, _ := ret[0].(db.Table)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreTable indicates an expected call of GetStoreTable.
func (mr *MockDBServiceMockRecorder) GetStoreTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreTable", reflect.TypeOf((*MockDBService)(nil).GetStoreTable), arg0, arg1)
}

// ListMenuFoodTag mocks base method.
func (m *MockDBService) ListMenuFoodTag(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenuFoodTag", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenuFoodTag indicates an expected call of ListMenuFoodTag.
func (mr *MockDBServiceMockRecorder) ListMenuFoodTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenuFoodTag", reflect.TypeOf((*MockDBService)(nil).ListMenuFoodTag), arg0, arg1)
}

// ListStoreMenu mocks base method.
func (m *MockDBService) ListStoreMenu(arg0 context.Context, arg1 db.ListStoreMenuParams) ([]db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoreMenu", arg0, arg1)
	ret0, _ := ret[0].([]db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoreMenu indicates an expected call of ListStoreMenu.
func (mr *MockDBServiceMockRecorder) ListStoreMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoreMenu", reflect.TypeOf((*MockDBService)(nil).ListStoreMenu), arg0, arg1)
}

// ListStoreTables mocks base method.
func (m *MockDBService) ListStoreTables(arg0 context.Context, arg1 db.ListStoreTablesParams) ([]db.Table, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoreTables", arg0, arg1)
	ret0, _ := ret[0].([]db.Table)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoreTables indicates an expected call of ListStoreTables.
func (mr *MockDBServiceMockRecorder) ListStoreTables(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoreTables", reflect.TypeOf((*MockDBService)(nil).ListStoreTables), arg0, arg1)
}

// RemoveMenuFoodTag mocks base method.
func (m *MockDBService) RemoveMenuFoodTag(arg0 context.Context, arg1 db.RemoveMenuFoodTagParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMenuFoodTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMenuFoodTag indicates an expected call of RemoveMenuFoodTag.
func (mr *MockDBServiceMockRecorder) RemoveMenuFoodTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMenuFoodTag", reflect.TypeOf((*MockDBService)(nil).RemoveMenuFoodTag), arg0, arg1)
}

// UpdateStore mocks base method.
func (m *MockDBService) UpdateStore(arg0 context.Context, arg1 db.UpdateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStore indicates an expected call of UpdateStore.
func (mr *MockDBServiceMockRecorder) UpdateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStore", reflect.TypeOf((*MockDBService)(nil).UpdateStore), arg0, arg1)
}

// UpdateStoreMenu mocks base method.
func (m *MockDBService) UpdateStoreMenu(arg0 context.Context, arg1 db.UpdateStoreMenuParams) (db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStoreMenu", arg0, arg1)
	ret0, _ := ret[0].(db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStoreMenu indicates an expected call of UpdateStoreMenu.
func (mr *MockDBServiceMockRecorder) UpdateStoreMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStoreMenu", reflect.TypeOf((*MockDBService)(nil).UpdateStoreMenu), arg0, arg1)
}

// UpdateStoreTable mocks base method.
func (m *MockDBService) UpdateStoreTable(arg0 context.Context, arg1 db.UpdateStoreTableParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStoreTable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStoreTable indicates an expected call of UpdateStoreTable.
func (mr *MockDBServiceMockRecorder) UpdateStoreTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStoreTable", reflect.TypeOf((*MockDBService)(nil).UpdateStoreTable), arg0, arg1)
}
