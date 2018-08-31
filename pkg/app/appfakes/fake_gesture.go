// Code generated by counterfeiter. DO NOT EDIT.
package appfakes

import (
	"sync"

	"github.com/vitakras/libinput-gestures/pkg/app"
)

type FakeGesture struct {
	IsCompleteStub        func() bool
	isCompleteMutex       sync.RWMutex
	isCompleteArgsForCall []struct{}
	isCompleteReturns     struct {
		result1 bool
	}
	isCompleteReturnsOnCall map[int]struct {
		result1 bool
	}
	DirectionStub        func() string
	directionMutex       sync.RWMutex
	directionArgsForCall []struct{}
	directionReturns     struct {
		result1 string
	}
	directionReturnsOnCall map[int]struct {
		result1 string
	}
	FingersStub        func() int
	fingersMutex       sync.RWMutex
	fingersArgsForCall []struct{}
	fingersReturns     struct {
		result1 int
	}
	fingersReturnsOnCall map[int]struct {
		result1 int
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGesture) IsComplete() bool {
	fake.isCompleteMutex.Lock()
	ret, specificReturn := fake.isCompleteReturnsOnCall[len(fake.isCompleteArgsForCall)]
	fake.isCompleteArgsForCall = append(fake.isCompleteArgsForCall, struct{}{})
	fake.recordInvocation("IsComplete", []interface{}{})
	fake.isCompleteMutex.Unlock()
	if fake.IsCompleteStub != nil {
		return fake.IsCompleteStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isCompleteReturns.result1
}

func (fake *FakeGesture) IsCompleteCallCount() int {
	fake.isCompleteMutex.RLock()
	defer fake.isCompleteMutex.RUnlock()
	return len(fake.isCompleteArgsForCall)
}

func (fake *FakeGesture) IsCompleteReturns(result1 bool) {
	fake.IsCompleteStub = nil
	fake.isCompleteReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGesture) IsCompleteReturnsOnCall(i int, result1 bool) {
	fake.IsCompleteStub = nil
	if fake.isCompleteReturnsOnCall == nil {
		fake.isCompleteReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isCompleteReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGesture) Direction() string {
	fake.directionMutex.Lock()
	ret, specificReturn := fake.directionReturnsOnCall[len(fake.directionArgsForCall)]
	fake.directionArgsForCall = append(fake.directionArgsForCall, struct{}{})
	fake.recordInvocation("Direction", []interface{}{})
	fake.directionMutex.Unlock()
	if fake.DirectionStub != nil {
		return fake.DirectionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.directionReturns.result1
}

func (fake *FakeGesture) DirectionCallCount() int {
	fake.directionMutex.RLock()
	defer fake.directionMutex.RUnlock()
	return len(fake.directionArgsForCall)
}

func (fake *FakeGesture) DirectionReturns(result1 string) {
	fake.DirectionStub = nil
	fake.directionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGesture) DirectionReturnsOnCall(i int, result1 string) {
	fake.DirectionStub = nil
	if fake.directionReturnsOnCall == nil {
		fake.directionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.directionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGesture) Fingers() int {
	fake.fingersMutex.Lock()
	ret, specificReturn := fake.fingersReturnsOnCall[len(fake.fingersArgsForCall)]
	fake.fingersArgsForCall = append(fake.fingersArgsForCall, struct{}{})
	fake.recordInvocation("Fingers", []interface{}{})
	fake.fingersMutex.Unlock()
	if fake.FingersStub != nil {
		return fake.FingersStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.fingersReturns.result1
}

func (fake *FakeGesture) FingersCallCount() int {
	fake.fingersMutex.RLock()
	defer fake.fingersMutex.RUnlock()
	return len(fake.fingersArgsForCall)
}

func (fake *FakeGesture) FingersReturns(result1 int) {
	fake.FingersStub = nil
	fake.fingersReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeGesture) FingersReturnsOnCall(i int, result1 int) {
	fake.FingersStub = nil
	if fake.fingersReturnsOnCall == nil {
		fake.fingersReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.fingersReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeGesture) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.typeReturns.result1
}

func (fake *FakeGesture) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeGesture) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGesture) TypeReturnsOnCall(i int, result1 string) {
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGesture) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isCompleteMutex.RLock()
	defer fake.isCompleteMutex.RUnlock()
	fake.directionMutex.RLock()
	defer fake.directionMutex.RUnlock()
	fake.fingersMutex.RLock()
	defer fake.fingersMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGesture) recordInvocation(key string, args []interface{}) {
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

var _ app.Gesture = new(FakeGesture)
