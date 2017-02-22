// Automatically generated by MockGen. DO NOT EDIT!
// Source: vendor/github.com/cloudfoundry/libbuildpack/manifest.go

package main_test

import (
	. "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *_MockManifestRecorder
}

// Recorder for MockManifest (not exported)
type _MockManifestRecorder struct {
	mock *MockManifest
}

func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &_MockManifestRecorder{mock}
	return mock
}

func (_m *MockManifest) EXPECT() *_MockManifestRecorder {
	return _m.recorder
}

func (_m *MockManifest) DefaultVersion(depName string) (Dependency, error) {
	ret := _m.ctrl.Call(_m, "DefaultVersion", depName)
	ret0, _ := ret[0].(Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManifestRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DefaultVersion", arg0)
}

func (_m *MockManifest) FetchDependency(dep Dependency, outputFile string) error {
	ret := _m.ctrl.Call(_m, "FetchDependency", dep, outputFile)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManifestRecorder) FetchDependency(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchDependency", arg0, arg1)
}

func (_m *MockManifest) InstallDependency(dep Dependency, outputDir string) error {
	ret := _m.ctrl.Call(_m, "InstallDependency", dep, outputDir)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManifestRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstallDependency", arg0, arg1)
}

func (_m *MockManifest) Version() (string, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManifestRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

func (_m *MockManifest) Language() string {
	ret := _m.ctrl.Call(_m, "Language")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockManifestRecorder) Language() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Language")
}

func (_m *MockManifest) CheckStackSupport() error {
	ret := _m.ctrl.Call(_m, "CheckStackSupport")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManifestRecorder) CheckStackSupport() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckStackSupport")
}

func (_m *MockManifest) RootDir() string {
	ret := _m.ctrl.Call(_m, "RootDir")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockManifestRecorder) RootDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RootDir")
}

func (_m *MockManifest) CheckBuildpackVersion(cacheDir string) {
	_m.ctrl.Call(_m, "CheckBuildpackVersion", cacheDir)
}

func (_mr *_MockManifestRecorder) CheckBuildpackVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckBuildpackVersion", arg0)
}

func (_m *MockManifest) StoreBuildpackMetadata(cacheDir string) {
	_m.ctrl.Call(_m, "StoreBuildpackMetadata", cacheDir)
}

func (_mr *_MockManifestRecorder) StoreBuildpackMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StoreBuildpackMetadata", arg0)
}
