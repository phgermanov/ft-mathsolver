// Code generated by counterfeiter. DO NOT EDIT.
package modelfakes

import (
	"sync"

	"github.com/phgermanov/ft-mathsolver/internal/model"
)

type FakeErrorRecorder struct {
	GetErrorsStub        func() []model.Error
	getErrorsMutex       sync.RWMutex
	getErrorsArgsForCall []struct {
	}
	getErrorsReturns struct {
		result1 []model.Error
	}
	getErrorsReturnsOnCall map[int]struct {
		result1 []model.Error
	}
	RecordErrorStub        func(model.Expression, string, error)
	recordErrorMutex       sync.RWMutex
	recordErrorArgsForCall []struct {
		arg1 model.Expression
		arg2 string
		arg3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeErrorRecorder) GetErrors() []model.Error {
	fake.getErrorsMutex.Lock()
	ret, specificReturn := fake.getErrorsReturnsOnCall[len(fake.getErrorsArgsForCall)]
	fake.getErrorsArgsForCall = append(fake.getErrorsArgsForCall, struct {
	}{})
	stub := fake.GetErrorsStub
	fakeReturns := fake.getErrorsReturns
	fake.recordInvocation("GetErrors", []interface{}{})
	fake.getErrorsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeErrorRecorder) GetErrorsCallCount() int {
	fake.getErrorsMutex.RLock()
	defer fake.getErrorsMutex.RUnlock()
	return len(fake.getErrorsArgsForCall)
}

func (fake *FakeErrorRecorder) GetErrorsCalls(stub func() []model.Error) {
	fake.getErrorsMutex.Lock()
	defer fake.getErrorsMutex.Unlock()
	fake.GetErrorsStub = stub
}

func (fake *FakeErrorRecorder) GetErrorsReturns(result1 []model.Error) {
	fake.getErrorsMutex.Lock()
	defer fake.getErrorsMutex.Unlock()
	fake.GetErrorsStub = nil
	fake.getErrorsReturns = struct {
		result1 []model.Error
	}{result1}
}

func (fake *FakeErrorRecorder) GetErrorsReturnsOnCall(i int, result1 []model.Error) {
	fake.getErrorsMutex.Lock()
	defer fake.getErrorsMutex.Unlock()
	fake.GetErrorsStub = nil
	if fake.getErrorsReturnsOnCall == nil {
		fake.getErrorsReturnsOnCall = make(map[int]struct {
			result1 []model.Error
		})
	}
	fake.getErrorsReturnsOnCall[i] = struct {
		result1 []model.Error
	}{result1}
}

func (fake *FakeErrorRecorder) RecordError(arg1 model.Expression, arg2 string, arg3 error) {
	fake.recordErrorMutex.Lock()
	fake.recordErrorArgsForCall = append(fake.recordErrorArgsForCall, struct {
		arg1 model.Expression
		arg2 string
		arg3 error
	}{arg1, arg2, arg3})
	stub := fake.RecordErrorStub
	fake.recordInvocation("RecordError", []interface{}{arg1, arg2, arg3})
	fake.recordErrorMutex.Unlock()
	if stub != nil {
		fake.RecordErrorStub(arg1, arg2, arg3)
	}
}

func (fake *FakeErrorRecorder) RecordErrorCallCount() int {
	fake.recordErrorMutex.RLock()
	defer fake.recordErrorMutex.RUnlock()
	return len(fake.recordErrorArgsForCall)
}

func (fake *FakeErrorRecorder) RecordErrorCalls(stub func(model.Expression, string, error)) {
	fake.recordErrorMutex.Lock()
	defer fake.recordErrorMutex.Unlock()
	fake.RecordErrorStub = stub
}

func (fake *FakeErrorRecorder) RecordErrorArgsForCall(i int) (model.Expression, string, error) {
	fake.recordErrorMutex.RLock()
	defer fake.recordErrorMutex.RUnlock()
	argsForCall := fake.recordErrorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeErrorRecorder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getErrorsMutex.RLock()
	defer fake.getErrorsMutex.RUnlock()
	fake.recordErrorMutex.RLock()
	defer fake.recordErrorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeErrorRecorder) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ model.ErrorRecorder = new(FakeErrorRecorder)