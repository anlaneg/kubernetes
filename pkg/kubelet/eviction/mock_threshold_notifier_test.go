/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package eviction is a generated GoMock package.
package eviction

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/kubelet/pkg/apis/stats/v1alpha1"
	reflect "reflect"
	time "time"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockManager) Start(diskInfoProvider DiskInfoProvider, podFunc ActivePodsFunc, podCleanedUpFunc PodCleanedUpFunc, monitoringInterval time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", diskInfoProvider, podFunc, podCleanedUpFunc, monitoringInterval)
}

// Start indicates an expected call of Start
func (mr *MockManagerMockRecorder) Start(diskInfoProvider, podFunc, podCleanedUpFunc, monitoringInterval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start), diskInfoProvider, podFunc, podCleanedUpFunc, monitoringInterval)
}

// IsUnderMemoryPressure mocks base method
func (m *MockManager) IsUnderMemoryPressure() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnderMemoryPressure")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnderMemoryPressure indicates an expected call of IsUnderMemoryPressure
func (mr *MockManagerMockRecorder) IsUnderMemoryPressure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnderMemoryPressure", reflect.TypeOf((*MockManager)(nil).IsUnderMemoryPressure))
}

// IsUnderDiskPressure mocks base method
func (m *MockManager) IsUnderDiskPressure() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnderDiskPressure")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnderDiskPressure indicates an expected call of IsUnderDiskPressure
func (mr *MockManagerMockRecorder) IsUnderDiskPressure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnderDiskPressure", reflect.TypeOf((*MockManager)(nil).IsUnderDiskPressure))
}

// IsUnderPIDPressure mocks base method
func (m *MockManager) IsUnderPIDPressure() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnderPIDPressure")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnderPIDPressure indicates an expected call of IsUnderPIDPressure
func (mr *MockManagerMockRecorder) IsUnderPIDPressure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnderPIDPressure", reflect.TypeOf((*MockManager)(nil).IsUnderPIDPressure))
}

// MockDiskInfoProvider is a mock of DiskInfoProvider interface
type MockDiskInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDiskInfoProviderMockRecorder
}

// MockDiskInfoProviderMockRecorder is the mock recorder for MockDiskInfoProvider
type MockDiskInfoProviderMockRecorder struct {
	mock *MockDiskInfoProvider
}

// NewMockDiskInfoProvider creates a new mock instance
func NewMockDiskInfoProvider(ctrl *gomock.Controller) *MockDiskInfoProvider {
	mock := &MockDiskInfoProvider{ctrl: ctrl}
	mock.recorder = &MockDiskInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiskInfoProvider) EXPECT() *MockDiskInfoProviderMockRecorder {
	return m.recorder
}

// HasDedicatedImageFs mocks base method
func (m *MockDiskInfoProvider) HasDedicatedImageFs() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasDedicatedImageFs")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasDedicatedImageFs indicates an expected call of HasDedicatedImageFs
func (mr *MockDiskInfoProviderMockRecorder) HasDedicatedImageFs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasDedicatedImageFs", reflect.TypeOf((*MockDiskInfoProvider)(nil).HasDedicatedImageFs))
}

// MockImageGC is a mock of ImageGC interface
type MockImageGC struct {
	ctrl     *gomock.Controller
	recorder *MockImageGCMockRecorder
}

// MockImageGCMockRecorder is the mock recorder for MockImageGC
type MockImageGCMockRecorder struct {
	mock *MockImageGC
}

// NewMockImageGC creates a new mock instance
func NewMockImageGC(ctrl *gomock.Controller) *MockImageGC {
	mock := &MockImageGC{ctrl: ctrl}
	mock.recorder = &MockImageGCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageGC) EXPECT() *MockImageGCMockRecorder {
	return m.recorder
}

// DeleteUnusedImages mocks base method
func (m *MockImageGC) DeleteUnusedImages() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUnusedImages")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnusedImages indicates an expected call of DeleteUnusedImages
func (mr *MockImageGCMockRecorder) DeleteUnusedImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnusedImages", reflect.TypeOf((*MockImageGC)(nil).DeleteUnusedImages))
}

// MockContainerGC is a mock of ContainerGC interface
type MockContainerGC struct {
	ctrl     *gomock.Controller
	recorder *MockContainerGCMockRecorder
}

// MockContainerGCMockRecorder is the mock recorder for MockContainerGC
type MockContainerGCMockRecorder struct {
	mock *MockContainerGC
}

// NewMockContainerGC creates a new mock instance
func NewMockContainerGC(ctrl *gomock.Controller) *MockContainerGC {
	mock := &MockContainerGC{ctrl: ctrl}
	mock.recorder = &MockContainerGCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainerGC) EXPECT() *MockContainerGCMockRecorder {
	return m.recorder
}

// DeleteAllUnusedContainers mocks base method
func (m *MockContainerGC) DeleteAllUnusedContainers() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllUnusedContainers")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllUnusedContainers indicates an expected call of DeleteAllUnusedContainers
func (mr *MockContainerGCMockRecorder) DeleteAllUnusedContainers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllUnusedContainers", reflect.TypeOf((*MockContainerGC)(nil).DeleteAllUnusedContainers))
}

// MockCgroupNotifier is a mock of CgroupNotifier interface
type MockCgroupNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockCgroupNotifierMockRecorder
}

// MockCgroupNotifierMockRecorder is the mock recorder for MockCgroupNotifier
type MockCgroupNotifierMockRecorder struct {
	mock *MockCgroupNotifier
}

// NewMockCgroupNotifier creates a new mock instance
func NewMockCgroupNotifier(ctrl *gomock.Controller) *MockCgroupNotifier {
	mock := &MockCgroupNotifier{ctrl: ctrl}
	mock.recorder = &MockCgroupNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCgroupNotifier) EXPECT() *MockCgroupNotifierMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockCgroupNotifier) Start(eventCh chan<- struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", eventCh)
}

// Start indicates an expected call of Start
func (mr *MockCgroupNotifierMockRecorder) Start(eventCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCgroupNotifier)(nil).Start), eventCh)
}

// Stop mocks base method
func (m *MockCgroupNotifier) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockCgroupNotifierMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCgroupNotifier)(nil).Stop))
}

// MockNotifierFactory is a mock of NotifierFactory interface
type MockNotifierFactory struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierFactoryMockRecorder
}

// MockNotifierFactoryMockRecorder is the mock recorder for MockNotifierFactory
type MockNotifierFactoryMockRecorder struct {
	mock *MockNotifierFactory
}

// NewMockNotifierFactory creates a new mock instance
func NewMockNotifierFactory(ctrl *gomock.Controller) *MockNotifierFactory {
	mock := &MockNotifierFactory{ctrl: ctrl}
	mock.recorder = &MockNotifierFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotifierFactory) EXPECT() *MockNotifierFactoryMockRecorder {
	return m.recorder
}

// NewCgroupNotifier mocks base method
func (m *MockNotifierFactory) NewCgroupNotifier(path, attribute string, threshold int64) (CgroupNotifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCgroupNotifier", path, attribute, threshold)
	ret0, _ := ret[0].(CgroupNotifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCgroupNotifier indicates an expected call of NewCgroupNotifier
func (mr *MockNotifierFactoryMockRecorder) NewCgroupNotifier(path, attribute, threshold interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCgroupNotifier", reflect.TypeOf((*MockNotifierFactory)(nil).NewCgroupNotifier), path, attribute, threshold)
}

// MockThresholdNotifier is a mock of ThresholdNotifier interface
type MockThresholdNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockThresholdNotifierMockRecorder
}

// MockThresholdNotifierMockRecorder is the mock recorder for MockThresholdNotifier
type MockThresholdNotifierMockRecorder struct {
	mock *MockThresholdNotifier
}

// NewMockThresholdNotifier creates a new mock instance
func NewMockThresholdNotifier(ctrl *gomock.Controller) *MockThresholdNotifier {
	mock := &MockThresholdNotifier{ctrl: ctrl}
	mock.recorder = &MockThresholdNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThresholdNotifier) EXPECT() *MockThresholdNotifierMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockThresholdNotifier) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockThresholdNotifierMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockThresholdNotifier)(nil).Start))
}

// UpdateThreshold mocks base method
func (m *MockThresholdNotifier) UpdateThreshold(summary *v1alpha1.Summary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThreshold", summary)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateThreshold indicates an expected call of UpdateThreshold
func (mr *MockThresholdNotifierMockRecorder) UpdateThreshold(summary interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThreshold", reflect.TypeOf((*MockThresholdNotifier)(nil).UpdateThreshold), summary)
}

// Description mocks base method
func (m *MockThresholdNotifier) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description
func (mr *MockThresholdNotifierMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockThresholdNotifier)(nil).Description))
}
