// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3aggregator/aggregator/handler/common (interfaces: Queue)

package common

import (
	"github.com/golang/mock/gomock"
)

// Mock of Queue interface
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *_MockQueueRecorder
}

// Recorder for MockQueue (not exported)
type _MockQueueRecorder struct {
	mock *MockQueue
}

func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &_MockQueueRecorder{mock}
	return mock
}

func (_m *MockQueue) EXPECT() *_MockQueueRecorder {
	return _m.recorder
}

func (_m *MockQueue) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockQueueRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockQueue) Enqueue(_param0 *RefCountedBuffer) error {
	ret := _m.ctrl.Call(_m, "Enqueue", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueueRecorder) Enqueue(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Enqueue", arg0)
}
