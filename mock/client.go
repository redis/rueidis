package mock

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

var _ rueidis.Client = (*Client)(nil)
var _ rueidis.DedicatedClient = (*DedicatedClient)(nil)

// Client is a mock of Client interface.
type Client struct {
	ctrl     *gomock.Controller
	recorder *ClientMockRecorder
}

// ClientMockRecorder is the mock recorder for Client.
type ClientMockRecorder struct {
	mock *Client
}

// NewClient creates a new mock instance.
func NewClient(ctrl *gomock.Controller) *Client {
	mock := &Client{ctrl: ctrl}
	mock.recorder = &ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Client) EXPECT() *ClientMockRecorder {
	return m.recorder
}

// B mocks base method.
func (m *Client) B() cmds.Builder {
	return cmds.NewBuilder(cmds.InitSlot)
}

// Close mocks base method.
func (m *Client) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *ClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Client)(nil).Close))
}

// Dedicate mocks base method.
func (m *Client) Dedicate() (rueidis.DedicatedClient, func()) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dedicate")
	ret0, _ := ret[0].(rueidis.DedicatedClient)
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// Dedicate indicates an expected call of Dedicate.
func (mr *ClientMockRecorder) Dedicate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dedicate", reflect.TypeOf((*Client)(nil).Dedicate))
}

// Dedicated mocks base method.
func (m *Client) Dedicated(arg0 func(rueidis.DedicatedClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dedicated", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dedicated indicates an expected call of Dedicated.
func (mr *ClientMockRecorder) Dedicated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dedicated", reflect.TypeOf((*Client)(nil).Dedicated), arg0)
}

// Do mocks base method.
func (m *Client) Do(arg0 context.Context, arg1 cmds.Completed) rueidis.RedisResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(rueidis.RedisResult)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *ClientMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*Client)(nil).Do), arg0, arg1)
}

// DoCache mocks base method.
func (m *Client) DoCache(arg0 context.Context, arg1 cmds.Cacheable, arg2 time.Duration) rueidis.RedisResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoCache", arg0, arg1, arg2)
	ret0, _ := ret[0].(rueidis.RedisResult)
	return ret0
}

// DoCache indicates an expected call of DoCache.
func (mr *ClientMockRecorder) DoCache(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoCache", reflect.TypeOf((*Client)(nil).DoCache), arg0, arg1, arg2)
}

// DoMulti mocks base method.
func (m *Client) DoMulti(arg0 context.Context, arg1 ...cmds.Completed) []rueidis.RedisResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMulti", varargs...)
	ret0, _ := ret[0].([]rueidis.RedisResult)
	return ret0
}

// DoMulti indicates an expected call of DoMulti.
func (mr *ClientMockRecorder) DoMulti(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMulti", reflect.TypeOf((*Client)(nil).DoMulti), varargs...)
}

// DoMultiCache mocks base method.
func (m *Client) DoMultiCache(arg0 context.Context, arg1 ...rueidis.CacheableTTL) []rueidis.RedisResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMultiCache", varargs...)
	ret0, _ := ret[0].([]rueidis.RedisResult)
	return ret0
}

// DoMultiCache indicates an expected call of DoMultiCache.
func (mr *ClientMockRecorder) DoMultiCache(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMultiCache", reflect.TypeOf((*Client)(nil).DoMultiCache), varargs...)
}

// Nodes mocks base method.
func (m *Client) Nodes() map[string]rueidis.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(map[string]rueidis.Client)
	return ret0
}

// Nodes indicates an expected call of Nodes.
func (mr *ClientMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*Client)(nil).Nodes))
}

// Receive mocks base method.
func (m *Client) Receive(arg0 context.Context, arg1 cmds.Completed, arg2 func(rueidis.PubSubMessage)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *ClientMockRecorder) Receive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*Client)(nil).Receive), arg0, arg1, arg2)
}

// DedicatedClient is a mock of DedicatedClient interface.
type DedicatedClient struct {
	ctrl     *gomock.Controller
	recorder *DedicatedClientMockRecorder
}

// DedicatedClientMockRecorder is the mock recorder for DedicatedClient.
type DedicatedClientMockRecorder struct {
	mock *DedicatedClient
}

// NewDedicatedClient creates a new mock instance.
func NewDedicatedClient(ctrl *gomock.Controller) *DedicatedClient {
	mock := &DedicatedClient{ctrl: ctrl}
	mock.recorder = &DedicatedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *DedicatedClient) EXPECT() *DedicatedClientMockRecorder {
	return m.recorder
}

// B mocks base method.
func (m *DedicatedClient) B() cmds.Builder {
	return cmds.NewBuilder(cmds.InitSlot)
}

// Close mocks base method.
func (m *DedicatedClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *DedicatedClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*DedicatedClient)(nil).Close))
}

// Do mocks base method.
func (m *DedicatedClient) Do(arg0 context.Context, arg1 cmds.Completed) rueidis.RedisResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(rueidis.RedisResult)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *DedicatedClientMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*DedicatedClient)(nil).Do), arg0, arg1)
}

// DoMulti mocks base method.
func (m *DedicatedClient) DoMulti(arg0 context.Context, arg1 ...cmds.Completed) []rueidis.RedisResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMulti", varargs...)
	ret0, _ := ret[0].([]rueidis.RedisResult)
	return ret0
}

// DoMulti indicates an expected call of DoMulti.
func (mr *DedicatedClientMockRecorder) DoMulti(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMulti", reflect.TypeOf((*DedicatedClient)(nil).DoMulti), varargs...)
}

// Receive mocks base method.
func (m *DedicatedClient) Receive(arg0 context.Context, arg1 cmds.Completed, arg2 func(rueidis.PubSubMessage)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *DedicatedClientMockRecorder) Receive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*DedicatedClient)(nil).Receive), arg0, arg1, arg2)
}

// SetPubSubHooks mocks base method.
func (m *DedicatedClient) SetPubSubHooks(arg0 rueidis.PubSubHooks) <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPubSubHooks", arg0)
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// SetPubSubHooks indicates an expected call of SetPubSubHooks.
func (mr *DedicatedClientMockRecorder) SetPubSubHooks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPubSubHooks", reflect.TypeOf((*DedicatedClient)(nil).SetPubSubHooks), arg0)
}
