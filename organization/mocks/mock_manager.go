// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/pivotalservices/cf-mgmt/organization (interfaces: Manager)

package mock_organization

import (
	gomock "github.com/golang/mock/gomock"
	cloudcontroller "github.com/pivotalservices/cf-mgmt/cloudcontroller"
)

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) CreateOrgs(_param0 string) error {
	ret := _m.ctrl.Call(_m, "CreateOrgs", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) CreateOrgs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOrgs", arg0)
}

func (_m *MockManager) CreateQuotas(_param0 string) error {
	ret := _m.ctrl.Call(_m, "CreateQuotas", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) CreateQuotas(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateQuotas", arg0)
}

func (_m *MockManager) DeleteOrgs(_param0 string, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "DeleteOrgs", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) DeleteOrgs(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteOrgs", arg0, arg1)
}

func (_m *MockManager) FindOrg(_param0 string) (*cloudcontroller.Org, error) {
	ret := _m.ctrl.Call(_m, "FindOrg", _param0)
	ret0, _ := ret[0].(*cloudcontroller.Org)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) FindOrg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindOrg", arg0)
}

func (_m *MockManager) GetOrgGUID(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetOrgGUID", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) GetOrgGUID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrgGUID", arg0)
}

func (_m *MockManager) UpdateOrgUsers(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "UpdateOrgUsers", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) UpdateOrgUsers(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateOrgUsers", arg0, arg1)
}
