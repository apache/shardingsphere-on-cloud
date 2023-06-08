// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/kubernetes/chaosmesh/chaosmesh.go

// Package mock_chaosmesh is a generated GoMock package.
package mock_chaosmesh

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	chaosmesh "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh"
	gomock "github.com/golang/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
)

// MockChaos is a mock of Chaos interface.
type MockChaos struct {
	ctrl     *gomock.Controller
	recorder *MockChaosMockRecorder
}

// MockChaosMockRecorder is the mock recorder for MockChaos.
type MockChaosMockRecorder struct {
	mock *MockChaos
}

// NewMockChaos creates a new mock instance.
func NewMockChaos(ctrl *gomock.Controller) *MockChaos {
	mock := &MockChaos{ctrl: ctrl}
	mock.recorder = &MockChaosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChaos) EXPECT() *MockChaosMockRecorder {
	return m.recorder
}

// CreateNetworkChaos mocks base method.
func (m *MockChaos) CreateNetworkChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetworkChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNetworkChaos indicates an expected call of CreateNetworkChaos.
func (mr *MockChaosMockRecorder) CreateNetworkChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkChaos", reflect.TypeOf((*MockChaos)(nil).CreateNetworkChaos), arg0, arg1)
}

// CreatePodChaos mocks base method.
func (m *MockChaos) CreatePodChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePodChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePodChaos indicates an expected call of CreatePodChaos.
func (mr *MockChaosMockRecorder) CreatePodChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePodChaos", reflect.TypeOf((*MockChaos)(nil).CreatePodChaos), arg0, arg1)
}

// DeleteNetworkChaos mocks base method.
func (m *MockChaos) DeleteNetworkChaos(arg0 context.Context, arg1 chaosmesh.NetworkChaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetworkChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetworkChaos indicates an expected call of DeleteNetworkChaos.
func (mr *MockChaosMockRecorder) DeleteNetworkChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkChaos", reflect.TypeOf((*MockChaos)(nil).DeleteNetworkChaos), arg0, arg1)
}

// DeletePodChaos mocks base method.
func (m *MockChaos) DeletePodChaos(arg0 context.Context, arg1 chaosmesh.PodChaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePodChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePodChaos indicates an expected call of DeletePodChaos.
func (mr *MockChaosMockRecorder) DeletePodChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePodChaos", reflect.TypeOf((*MockChaos)(nil).DeletePodChaos), arg0, arg1)
}

// GetNetworkChaosByNamespacedName mocks base method.
func (m *MockChaos) GetNetworkChaosByNamespacedName(arg0 context.Context, arg1 types.NamespacedName) (chaosmesh.NetworkChaos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkChaosByNamespacedName", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.NetworkChaos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkChaosByNamespacedName indicates an expected call of GetNetworkChaosByNamespacedName.
func (mr *MockChaosMockRecorder) GetNetworkChaosByNamespacedName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkChaosByNamespacedName", reflect.TypeOf((*MockChaos)(nil).GetNetworkChaosByNamespacedName), arg0, arg1)
}

// GetPodChaosByNamespacedName mocks base method.
func (m *MockChaos) GetPodChaosByNamespacedName(arg0 context.Context, arg1 types.NamespacedName) (chaosmesh.PodChaos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodChaosByNamespacedName", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.PodChaos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodChaosByNamespacedName indicates an expected call of GetPodChaosByNamespacedName.
func (mr *MockChaosMockRecorder) GetPodChaosByNamespacedName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodChaosByNamespacedName", reflect.TypeOf((*MockChaos)(nil).GetPodChaosByNamespacedName), arg0, arg1)
}

// NewNetworkChaos mocks base method.
func (m *MockChaos) NewNetworkChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) chaosmesh.NetworkChaos {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNetworkChaos", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.NetworkChaos)
	return ret0
}

// NewNetworkChaos indicates an expected call of NewNetworkChaos.
func (mr *MockChaosMockRecorder) NewNetworkChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNetworkChaos", reflect.TypeOf((*MockChaos)(nil).NewNetworkChaos), arg0, arg1)
}

// NewPodChaos mocks base method.
func (m *MockChaos) NewPodChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) chaosmesh.PodChaos {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPodChaos", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.PodChaos)
	return ret0
}

// NewPodChaos indicates an expected call of NewPodChaos.
func (mr *MockChaosMockRecorder) NewPodChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPodChaos", reflect.TypeOf((*MockChaos)(nil).NewPodChaos), arg0, arg1)
}

// UpdateNetworkChaos mocks base method.
func (m *MockChaos) UpdateNetworkChaos(arg0 context.Context, arg1 chaosmesh.NetworkChaos, arg2 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNetworkChaos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNetworkChaos indicates an expected call of UpdateNetworkChaos.
func (mr *MockChaosMockRecorder) UpdateNetworkChaos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNetworkChaos", reflect.TypeOf((*MockChaos)(nil).UpdateNetworkChaos), arg0, arg1, arg2)
}

// UpdatePodChaos mocks base method.
func (m *MockChaos) UpdatePodChaos(arg0 context.Context, arg1 chaosmesh.PodChaos, arg2 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePodChaos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePodChaos indicates an expected call of UpdatePodChaos.
func (mr *MockChaosMockRecorder) UpdatePodChaos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePodChaos", reflect.TypeOf((*MockChaos)(nil).UpdatePodChaos), arg0, arg1, arg2)
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// NewNetworkChaos mocks base method.
func (m *MockBuilder) NewNetworkChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) chaosmesh.NetworkChaos {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNetworkChaos", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.NetworkChaos)
	return ret0
}

// NewNetworkChaos indicates an expected call of NewNetworkChaos.
func (mr *MockBuilderMockRecorder) NewNetworkChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNetworkChaos", reflect.TypeOf((*MockBuilder)(nil).NewNetworkChaos), arg0, arg1)
}

// NewPodChaos mocks base method.
func (m *MockBuilder) NewPodChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) chaosmesh.PodChaos {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPodChaos", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.PodChaos)
	return ret0
}

// NewPodChaos indicates an expected call of NewPodChaos.
func (mr *MockBuilderMockRecorder) NewPodChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPodChaos", reflect.TypeOf((*MockBuilder)(nil).NewPodChaos), arg0, arg1)
}

// MockGetter is a mock of Getter interface.
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetterMockRecorder
}

// MockGetterMockRecorder is the mock recorder for MockGetter.
type MockGetterMockRecorder struct {
	mock *MockGetter
}

// NewMockGetter creates a new mock instance.
func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &MockGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetter) EXPECT() *MockGetterMockRecorder {
	return m.recorder
}

// GetNetworkChaosByNamespacedName mocks base method.
func (m *MockGetter) GetNetworkChaosByNamespacedName(arg0 context.Context, arg1 types.NamespacedName) (chaosmesh.NetworkChaos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkChaosByNamespacedName", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.NetworkChaos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkChaosByNamespacedName indicates an expected call of GetNetworkChaosByNamespacedName.
func (mr *MockGetterMockRecorder) GetNetworkChaosByNamespacedName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkChaosByNamespacedName", reflect.TypeOf((*MockGetter)(nil).GetNetworkChaosByNamespacedName), arg0, arg1)
}

// GetPodChaosByNamespacedName mocks base method.
func (m *MockGetter) GetPodChaosByNamespacedName(arg0 context.Context, arg1 types.NamespacedName) (chaosmesh.PodChaos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodChaosByNamespacedName", arg0, arg1)
	ret0, _ := ret[0].(chaosmesh.PodChaos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodChaosByNamespacedName indicates an expected call of GetPodChaosByNamespacedName.
func (mr *MockGetterMockRecorder) GetPodChaosByNamespacedName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodChaosByNamespacedName", reflect.TypeOf((*MockGetter)(nil).GetPodChaosByNamespacedName), arg0, arg1)
}

// MockSetter is a mock of Setter interface.
type MockSetter struct {
	ctrl     *gomock.Controller
	recorder *MockSetterMockRecorder
}

// MockSetterMockRecorder is the mock recorder for MockSetter.
type MockSetterMockRecorder struct {
	mock *MockSetter
}

// NewMockSetter creates a new mock instance.
func NewMockSetter(ctrl *gomock.Controller) *MockSetter {
	mock := &MockSetter{ctrl: ctrl}
	mock.recorder = &MockSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSetter) EXPECT() *MockSetterMockRecorder {
	return m.recorder
}

// CreateNetworkChaos mocks base method.
func (m *MockSetter) CreateNetworkChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetworkChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNetworkChaos indicates an expected call of CreateNetworkChaos.
func (mr *MockSetterMockRecorder) CreateNetworkChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkChaos", reflect.TypeOf((*MockSetter)(nil).CreateNetworkChaos), arg0, arg1)
}

// CreatePodChaos mocks base method.
func (m *MockSetter) CreatePodChaos(arg0 context.Context, arg1 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePodChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePodChaos indicates an expected call of CreatePodChaos.
func (mr *MockSetterMockRecorder) CreatePodChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePodChaos", reflect.TypeOf((*MockSetter)(nil).CreatePodChaos), arg0, arg1)
}

// DeleteNetworkChaos mocks base method.
func (m *MockSetter) DeleteNetworkChaos(arg0 context.Context, arg1 chaosmesh.NetworkChaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetworkChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetworkChaos indicates an expected call of DeleteNetworkChaos.
func (mr *MockSetterMockRecorder) DeleteNetworkChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkChaos", reflect.TypeOf((*MockSetter)(nil).DeleteNetworkChaos), arg0, arg1)
}

// DeletePodChaos mocks base method.
func (m *MockSetter) DeletePodChaos(arg0 context.Context, arg1 chaosmesh.PodChaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePodChaos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePodChaos indicates an expected call of DeletePodChaos.
func (mr *MockSetterMockRecorder) DeletePodChaos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePodChaos", reflect.TypeOf((*MockSetter)(nil).DeletePodChaos), arg0, arg1)
}

// UpdateNetworkChaos mocks base method.
func (m *MockSetter) UpdateNetworkChaos(arg0 context.Context, arg1 chaosmesh.NetworkChaos, arg2 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNetworkChaos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNetworkChaos indicates an expected call of UpdateNetworkChaos.
func (mr *MockSetterMockRecorder) UpdateNetworkChaos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNetworkChaos", reflect.TypeOf((*MockSetter)(nil).UpdateNetworkChaos), arg0, arg1, arg2)
}

// UpdatePodChaos mocks base method.
func (m *MockSetter) UpdatePodChaos(arg0 context.Context, arg1 chaosmesh.PodChaos, arg2 *v1alpha1.Chaos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePodChaos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePodChaos indicates an expected call of UpdatePodChaos.
func (mr *MockSetterMockRecorder) UpdatePodChaos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePodChaos", reflect.TypeOf((*MockSetter)(nil).UpdatePodChaos), arg0, arg1, arg2)
}

// MockPodChaos is a mock of PodChaos interface.
type MockPodChaos struct {
	ctrl     *gomock.Controller
	recorder *MockPodChaosMockRecorder
}

// MockPodChaosMockRecorder is the mock recorder for MockPodChaos.
type MockPodChaosMockRecorder struct {
	mock *MockPodChaos
}

// NewMockPodChaos creates a new mock instance.
func NewMockPodChaos(ctrl *gomock.Controller) *MockPodChaos {
	mock := &MockPodChaos{ctrl: ctrl}
	mock.recorder = &MockPodChaosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodChaos) EXPECT() *MockPodChaosMockRecorder {
	return m.recorder
}

// MockNetworkChaos is a mock of NetworkChaos interface.
type MockNetworkChaos struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkChaosMockRecorder
}

// MockNetworkChaosMockRecorder is the mock recorder for MockNetworkChaos.
type MockNetworkChaosMockRecorder struct {
	mock *MockNetworkChaos
}

// NewMockNetworkChaos creates a new mock instance.
func NewMockNetworkChaos(ctrl *gomock.Controller) *MockNetworkChaos {
	mock := &MockNetworkChaos{ctrl: ctrl}
	mock.recorder = &MockNetworkChaosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkChaos) EXPECT() *MockNetworkChaosMockRecorder {
	return m.recorder
}
