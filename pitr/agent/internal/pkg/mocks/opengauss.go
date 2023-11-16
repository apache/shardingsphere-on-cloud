// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/opengauss.go

// Package mock_pkg is a generated GoMock package.
package mock_pkg

import (
	model "github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIOpenGauss is a mock of IOpenGauss interface
type MockIOpenGauss struct {
	ctrl     *gomock.Controller
	recorder *MockIOpenGaussMockRecorder
}

// MockIOpenGaussMockRecorder is the mock recorder for MockIOpenGauss
type MockIOpenGaussMockRecorder struct {
	mock *MockIOpenGauss
}

// NewMockIOpenGauss creates a new mock instance
func NewMockIOpenGauss(ctrl *gomock.Controller) *MockIOpenGauss {
	mock := &MockIOpenGauss{ctrl: ctrl}
	mock.recorder = &MockIOpenGaussMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIOpenGauss) EXPECT() *MockIOpenGaussMockRecorder {
	return m.recorder
}

// AsyncBackup mocks base method
func (m *MockIOpenGauss) AsyncBackup(backupPath, instanceName, backupMode string, threadsNum uint8, dbPort uint16) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncBackup", backupPath, instanceName, backupMode, threadsNum, dbPort)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AsyncBackup indicates an expected call of AsyncBackup
func (mr *MockIOpenGaussMockRecorder) AsyncBackup(backupPath, instanceName, backupMode, threadsNum, dbPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncBackup", reflect.TypeOf((*MockIOpenGauss)(nil).AsyncBackup), backupPath, instanceName, backupMode, threadsNum, dbPort)
}

// ShowBackup mocks base method
func (m *MockIOpenGauss) ShowBackup(backupPath, instanceName, backupID string) (*model.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowBackup", backupPath, instanceName, backupID)
	ret0, _ := ret[0].(*model.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowBackup indicates an expected call of ShowBackup
func (mr *MockIOpenGaussMockRecorder) ShowBackup(backupPath, instanceName, backupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowBackup", reflect.TypeOf((*MockIOpenGauss)(nil).ShowBackup), backupPath, instanceName, backupID)
}

// Init mocks base method
func (m *MockIOpenGauss) Init(backupPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", backupPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockIOpenGaussMockRecorder) Init(backupPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockIOpenGauss)(nil).Init), backupPath)
}

// AddInstance mocks base method
func (m *MockIOpenGauss) AddInstance(backupPath, instance string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInstance", backupPath, instance)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInstance indicates an expected call of AddInstance
func (mr *MockIOpenGaussMockRecorder) AddInstance(backupPath, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInstance", reflect.TypeOf((*MockIOpenGauss)(nil).AddInstance), backupPath, instance)
}

// DelInstance mocks base method
func (m *MockIOpenGauss) DelInstance(backupPath, instance string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelInstance", backupPath, instance)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelInstance indicates an expected call of DelInstance
func (mr *MockIOpenGaussMockRecorder) DelInstance(backupPath, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelInstance", reflect.TypeOf((*MockIOpenGauss)(nil).DelInstance), backupPath, instance)
}

// DelBackup mocks base method
func (m *MockIOpenGauss) DelBackup(backupPath, instance, backupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelBackup", backupPath, instance, backupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelBackup indicates an expected call of DelBackup
func (mr *MockIOpenGaussMockRecorder) DelBackup(backupPath, instance, backupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelBackup", reflect.TypeOf((*MockIOpenGauss)(nil).DelBackup), backupPath, instance, backupID)
}

// Start mocks base method
func (m *MockIOpenGauss) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockIOpenGaussMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockIOpenGauss)(nil).Start))
}

// Stop mocks base method
func (m *MockIOpenGauss) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockIOpenGaussMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIOpenGauss)(nil).Stop))
}

// Status mocks base method
func (m *MockIOpenGauss) Status() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockIOpenGaussMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockIOpenGauss)(nil).Status))
}

// Restore mocks base method
func (m *MockIOpenGauss) Restore(backupPath, instance, backupID string, threadsNum uint8) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", backupPath, instance, backupID, threadsNum)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockIOpenGaussMockRecorder) Restore(backupPath, instance, backupID, threadsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockIOpenGauss)(nil).Restore), backupPath, instance, backupID, threadsNum)
}

// ShowBackupList mocks base method
func (m *MockIOpenGauss) ShowBackupList(backupPath, instanceName string) ([]*model.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowBackupList", backupPath, instanceName)
	ret0, _ := ret[0].([]*model.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowBackupList indicates an expected call of ShowBackupList
func (mr *MockIOpenGaussMockRecorder) ShowBackupList(backupPath, instanceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowBackupList", reflect.TypeOf((*MockIOpenGauss)(nil).ShowBackupList), backupPath, instanceName)
}

// Auth mocks base method
func (m *MockIOpenGauss) Auth(user, password, dbName string, dbPort uint16) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", user, password, dbName, dbPort)
	ret0, _ := ret[0].(error)
	return ret0
}

// Auth indicates an expected call of Auth
func (mr *MockIOpenGaussMockRecorder) Auth(user, password, dbName, dbPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockIOpenGauss)(nil).Auth), user, password, dbName, dbPort)
}

// CheckSchema mocks base method
func (m *MockIOpenGauss) CheckSchema(user, password, dbName string, dbPort uint16, schema string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSchema", user, password, dbName, dbPort, schema)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSchema indicates an expected call of CheckSchema
func (mr *MockIOpenGaussMockRecorder) CheckSchema(user, password, dbName, dbPort, schema interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSchema", reflect.TypeOf((*MockIOpenGauss)(nil).CheckSchema), user, password, dbName, dbPort, schema)
}

// MvTempToPgData mocks base method
func (m *MockIOpenGauss) MvTempToPgData() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MvTempToPgData")
	ret0, _ := ret[0].(error)
	return ret0
}

// MvTempToPgData indicates an expected call of MvTempToPgData
func (mr *MockIOpenGaussMockRecorder) MvTempToPgData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MvTempToPgData", reflect.TypeOf((*MockIOpenGauss)(nil).MvTempToPgData))
}

// MvPgDataToTemp mocks base method
func (m *MockIOpenGauss) MvPgDataToTemp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MvPgDataToTemp")
	ret0, _ := ret[0].(error)
	return ret0
}

// MvPgDataToTemp indicates an expected call of MvPgDataToTemp
func (mr *MockIOpenGaussMockRecorder) MvPgDataToTemp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MvPgDataToTemp", reflect.TypeOf((*MockIOpenGauss)(nil).MvPgDataToTemp))
}

// CleanPgDataTemp mocks base method
func (m *MockIOpenGauss) CleanPgDataTemp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanPgDataTemp")
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanPgDataTemp indicates an expected call of CleanPgDataTemp
func (mr *MockIOpenGaussMockRecorder) CleanPgDataTemp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanPgDataTemp", reflect.TypeOf((*MockIOpenGauss)(nil).CleanPgDataTemp))
}
