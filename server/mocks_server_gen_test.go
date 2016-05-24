// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/smancke/guble/server (interfaces: Router,WSConnection,Startable,Stopable,Endpoint)

package server

import (
	gomock "github.com/golang/mock/gomock"
	guble "github.com/smancke/guble/guble"
	
	auth "github.com/smancke/guble/server/auth"
	store "github.com/smancke/guble/store"
	http "net/http"
)

// Mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *_MockRouterRecorder
}

// Recorder for MockRouter (not exported)
type _MockRouterRecorder struct {
	mock *MockRouter
}

func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &_MockRouterRecorder{mock}
	return mock
}

func (_m *MockRouter) EXPECT() *_MockRouterRecorder {
	return _m.recorder
}

func (_m *MockRouter) AccessManager() (auth.AccessManager, error) {
	ret := _m.ctrl.Call(_m, "AccessManager")
	ret0, _ := ret[0].(auth.AccessManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) AccessManager() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AccessManager")
}

func (_m *MockRouter) HandleMessage(_param0 *guble.Message) error {
	ret := _m.ctrl.Call(_m, "HandleMessage", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouterRecorder) HandleMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleMessage", arg0)
}

func (_m *MockRouter) KVStore() (store.KVStore, error) {
	ret := _m.ctrl.Call(_m, "KVStore")
	ret0, _ := ret[0].(store.KVStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) KVStore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KVStore")
}

func (_m *MockRouter) MessageStore() (store.MessageStore, error) {
	ret := _m.ctrl.Call(_m, "MessageStore")
	ret0, _ := ret[0].(store.MessageStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) MessageStore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MessageStore")
}

func (_m *MockRouter) Subscribe(_param0 *Route) (*Route, error) {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0)
	ret0, _ := ret[0].(*Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", arg0)
}

func (_m *MockRouter) Unsubscribe(_param0 *Route) {
	_m.ctrl.Call(_m, "Unsubscribe", _param0)
}

func (_mr *_MockRouterRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unsubscribe", arg0)
}

// Mock of WSConnection interface
type MockWSConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockWSConnectionRecorder
}

// Recorder for MockWSConnection (not exported)
type _MockWSConnectionRecorder struct {
	mock *MockWSConnection
}

func NewMockWSConnection(ctrl *gomock.Controller) *MockWSConnection {
	mock := &MockWSConnection{ctrl: ctrl}
	mock.recorder = &_MockWSConnectionRecorder{mock}
	return mock
}

func (_m *MockWSConnection) EXPECT() *_MockWSConnectionRecorder {
	return _m.recorder
}

func (_m *MockWSConnection) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockWSConnectionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockWSConnection) Receive(_param0 *[]byte) error {
	ret := _m.ctrl.Call(_m, "Receive", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWSConnectionRecorder) Receive(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Receive", arg0)
}

func (_m *MockWSConnection) Send(_param0 []byte) error {
	ret := _m.ctrl.Call(_m, "Send", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWSConnectionRecorder) Send(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0)
}

// Mock of Startable interface
type MockStartable struct {
	ctrl     *gomock.Controller
	recorder *_MockStartableRecorder
}

// Recorder for MockStartable (not exported)
type _MockStartableRecorder struct {
	mock *MockStartable
}

func NewMockStartable(ctrl *gomock.Controller) *MockStartable {
	mock := &MockStartable{ctrl: ctrl}
	mock.recorder = &_MockStartableRecorder{mock}
	return mock
}

func (_m *MockStartable) EXPECT() *_MockStartableRecorder {
	return _m.recorder
}

func (_m *MockStartable) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStartableRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

// Mock of Stopable interface
type MockStopable struct {
	ctrl     *gomock.Controller
	recorder *_MockStopableRecorder
}

// Recorder for MockStopable (not exported)
type _MockStopableRecorder struct {
	mock *MockStopable
}

func NewMockStopable(ctrl *gomock.Controller) *MockStopable {
	mock := &MockStopable{ctrl: ctrl}
	mock.recorder = &_MockStopableRecorder{mock}
	return mock
}

func (_m *MockStopable) EXPECT() *_MockStopableRecorder {
	return _m.recorder
}

func (_m *MockStopable) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStopableRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

// Mock of Endpoint interface
type MockEndpoint struct {
	ctrl     *gomock.Controller
	recorder *_MockEndpointRecorder
}

// Recorder for MockEndpoint (not exported)
type _MockEndpointRecorder struct {
	mock *MockEndpoint
}

func NewMockEndpoint(ctrl *gomock.Controller) *MockEndpoint {
	mock := &MockEndpoint{ctrl: ctrl}
	mock.recorder = &_MockEndpointRecorder{mock}
	return mock
}

func (_m *MockEndpoint) EXPECT() *_MockEndpointRecorder {
	return _m.recorder
}

func (_m *MockEndpoint) GetPrefix() string {
	ret := _m.ctrl.Call(_m, "GetPrefix")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockEndpointRecorder) GetPrefix() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPrefix")
}

func (_m *MockEndpoint) ServeHTTP(_param0 http.ResponseWriter, _param1 *http.Request) {
	_m.ctrl.Call(_m, "ServeHTTP", _param0, _param1)
}

func (_mr *_MockEndpointRecorder) ServeHTTP(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServeHTTP", arg0, arg1)
}
