package mock

import (
	"context"
	"reflect"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
	"go.uber.org/mock/gomock"
)

var _ rueidis.Client = (*Client)(nil)
var _ rueidis.DedicatedClient = (*DedicatedClient)(nil)

// ClientOption is an optional function parameter for NewClient
type ClientOption func(c any)

// WithSlotCheck enables the command builder of Client to check if the command built across multiple slots and then panic
func WithSlotCheck() ClientOption {
	return func(c any) {
		if cc, ok := c.(*Client); ok {
			cc.slot = cmds.InitSlot
		}
		if cc, ok := c.(*DedicatedClient); ok {
			cc.slot = cmds.InitSlot
		}
	}
}

// Client mocks the Client interface.
type Client struct {
	ctrl     *gomock.Controller
	recorder *ClientMockRecorder
	slot     uint16
}

// ClientMockRecorder is the mock recorder for Client.
type ClientMockRecorder struct {
	mock *Client
}

// NewClient creates a new mock instance.
func NewClient(ctrl *gomock.Controller, options ...ClientOption) *Client {
	mock := &Client{ctrl: ctrl, slot: cmds.NoSlot}
	mock.recorder = &ClientMockRecorder{mock}
	for _, opt := range options {
		opt(mock)
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate the expected use.
func (m *Client) EXPECT() *ClientMockRecorder {
	return m.recorder
}

// B mocks base method.
func (m *Client) B() rueidis.Builder {
	return cmds.NewBuilder(m.slot)
}

// Mode mocks base method.
func (m *Client) Mode() rueidis.ClientMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(rueidis.ClientMode)
	return ret0
}

// Mode indicates an expected call of Mode.
func (mr *ClientMockRecorder) Mode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*Client)(nil).Mode))
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
func (mr *ClientMockRecorder) Dedicated(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dedicated", reflect.TypeOf((*Client)(nil).Dedicated), arg0)
}

// Do mocks the base method.
func (m *Client) Do(arg0 context.Context, arg1 rueidis.Completed) rueidis.RedisResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(rueidis.RedisResult)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *ClientMockRecorder) Do(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*Client)(nil).Do), arg0, arg1)
}

// DoStream mocks base method.
func (m *Client) DoStream(arg0 context.Context, arg1 rueidis.Completed) rueidis.RedisResultStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoStream", arg0, arg1)
	ret0, _ := ret[0].(rueidis.RedisResultStream)
	return ret0
}

// DoStream indicates an expected call of DoStream.
func (mr *ClientMockRecorder) DoStream(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoStream", reflect.TypeOf((*Client)(nil).DoStream), arg0, arg1)
}

// DoCache mocks base method.
func (m *Client) DoCache(arg0 context.Context, arg1 rueidis.Cacheable, arg2 time.Duration) rueidis.RedisResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoCache", arg0, arg1, arg2)
	ret0, _ := ret[0].(rueidis.RedisResult)
	return ret0
}

// DoCache indicates an expected call of DoCache.
func (mr *ClientMockRecorder) DoCache(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoCache", reflect.TypeOf((*Client)(nil).DoCache), arg0, arg1, arg2)
}

// DoMulti mocks base method.
func (m *Client) DoMulti(arg0 context.Context, arg1 ...rueidis.Completed) []rueidis.RedisResult {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMulti", varargs...)
	ret0, _ := ret[0].([]rueidis.RedisResult)
	return ret0
}

// DoMulti indicates an expected call of DoMulti.
func (mr *ClientMockRecorder) DoMulti(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMulti", reflect.TypeOf((*Client)(nil).DoMulti), varargs...)
}

// DoMultiStream mocks base method.
func (m *Client) DoMultiStream(arg0 context.Context, arg1 ...rueidis.Completed) rueidis.MultiRedisResultStream {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMultiStream", varargs...)
	ret0, _ := ret[0].(rueidis.MultiRedisResultStream)
	return ret0
}

// DoMultiStream indicates an expected call of DoMultiStream.
func (mr *ClientMockRecorder) DoMultiStream(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMultiStream", reflect.TypeOf((*Client)(nil).DoMultiStream), varargs...)
}

// DoMultiCache mocks the base method.
func (m *Client) DoMultiCache(arg0 context.Context, arg1 ...rueidis.CacheableTTL) []rueidis.RedisResult {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMultiCache", varargs...)
	ret0, _ := ret[0].([]rueidis.RedisResult)
	return ret0
}

// DoMultiCache indicates an expected call of DoMultiCache.
func (mr *ClientMockRecorder) DoMultiCache(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMultiCache", reflect.TypeOf((*Client)(nil).DoMultiCache), varargs...)
}

// Nodes mocks the base method.
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
func (m *Client) Receive(arg0 context.Context, arg1 rueidis.Completed, arg2 func(rueidis.PubSubMessage)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *ClientMockRecorder) Receive(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*Client)(nil).Receive), arg0, arg1, arg2)
}

// DedicatedClient mocks the DedicatedClient interface.
type DedicatedClient struct {
	ctrl     *gomock.Controller
	recorder *DedicatedClientMockRecorder
	slot     uint16
}

// DedicatedClientMockRecorder is the mock recorder for DedicatedClient.
type DedicatedClientMockRecorder struct {
	mock *DedicatedClient
}

// NewDedicatedClient creates a new mock instance.
func NewDedicatedClient(ctrl *gomock.Controller, options ...ClientOption) *DedicatedClient {
	mock := &DedicatedClient{ctrl: ctrl, slot: cmds.NoSlot}
	mock.recorder = &DedicatedClientMockRecorder{mock}
	for _, opt := range options {
		opt(mock)
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate the expected use.
func (m *DedicatedClient) EXPECT() *DedicatedClientMockRecorder {
	return m.recorder
}

// B mocks base method.
func (m *DedicatedClient) B() rueidis.Builder {
	return cmds.NewBuilder(m.slot)
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

// Do mocks the base method.
func (m *DedicatedClient) Do(arg0 context.Context, arg1 rueidis.Completed) rueidis.RedisResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(rueidis.RedisResult)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *DedicatedClientMockRecorder) Do(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*DedicatedClient)(nil).Do), arg0, arg1)
}

// DoMulti mocks base method.
func (m *DedicatedClient) DoMulti(arg0 context.Context, arg1 ...rueidis.Completed) []rueidis.RedisResult {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoMulti", varargs...)
	ret0, _ := ret[0].([]rueidis.RedisResult)
	return ret0
}

// DoMulti indicates an expected call of DoMulti.
func (mr *DedicatedClientMockRecorder) DoMulti(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoMulti", reflect.TypeOf((*DedicatedClient)(nil).DoMulti), varargs...)
}

// Receive mocks base method.
func (m *DedicatedClient) Receive(arg0 context.Context, arg1 rueidis.Completed, arg2 func(rueidis.PubSubMessage)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *DedicatedClientMockRecorder) Receive(arg0, arg1, arg2 any) *gomock.Call {
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
func (mr *DedicatedClientMockRecorder) SetPubSubHooks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPubSubHooks", reflect.TypeOf((*DedicatedClient)(nil).SetPubSubHooks), arg0)
}
