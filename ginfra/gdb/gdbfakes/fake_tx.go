// Code generated by counterfeiter. DO NOT EDIT.
package gdbfakes

import (
	"context"
	"sync"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type FakeTx struct {
	DoTransactionStub        func(context.Context, *gdb.TxOption, func(c context.Context) (commit bool, err error)) error
	doTransactionMutex       sync.RWMutex
	doTransactionArgsForCall []struct {
		arg1 context.Context
		arg2 *gdb.TxOption
		arg3 func(c context.Context) (commit bool, err error)
	}
	doTransactionReturns struct {
		result1 error
	}
	doTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTx) DoTransaction(arg1 context.Context, arg2 *gdb.TxOption, arg3 func(c context.Context) (commit bool, err error)) error {
	fake.doTransactionMutex.Lock()
	ret, specificReturn := fake.doTransactionReturnsOnCall[len(fake.doTransactionArgsForCall)]
	fake.doTransactionArgsForCall = append(fake.doTransactionArgsForCall, struct {
		arg1 context.Context
		arg2 *gdb.TxOption
		arg3 func(c context.Context) (commit bool, err error)
	}{arg1, arg2, arg3})
	stub := fake.DoTransactionStub
	fakeReturns := fake.doTransactionReturns
	fake.recordInvocation("DoTransaction", []interface{}{arg1, arg2, arg3})
	fake.doTransactionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTx) DoTransactionCallCount() int {
	fake.doTransactionMutex.RLock()
	defer fake.doTransactionMutex.RUnlock()
	return len(fake.doTransactionArgsForCall)
}

func (fake *FakeTx) DoTransactionCalls(stub func(context.Context, *gdb.TxOption, func(c context.Context) (commit bool, err error)) error) {
	fake.doTransactionMutex.Lock()
	defer fake.doTransactionMutex.Unlock()
	fake.DoTransactionStub = stub
}

func (fake *FakeTx) DoTransactionArgsForCall(i int) (context.Context, *gdb.TxOption, func(c context.Context) (commit bool, err error)) {
	fake.doTransactionMutex.RLock()
	defer fake.doTransactionMutex.RUnlock()
	argsForCall := fake.doTransactionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTx) DoTransactionReturns(result1 error) {
	fake.doTransactionMutex.Lock()
	defer fake.doTransactionMutex.Unlock()
	fake.DoTransactionStub = nil
	fake.doTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) DoTransactionReturnsOnCall(i int, result1 error) {
	fake.doTransactionMutex.Lock()
	defer fake.doTransactionMutex.Unlock()
	fake.DoTransactionStub = nil
	if fake.doTransactionReturnsOnCall == nil {
		fake.doTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doTransactionMutex.RLock()
	defer fake.doTransactionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTx) recordInvocation(key string, args []interface{}) {
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

var _ gdb.Tx = new(FakeTx)
