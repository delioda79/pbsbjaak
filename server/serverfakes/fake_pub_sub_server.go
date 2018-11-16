// Code generated by counterfeiter. DO NOT EDIT.
package serverfakes

import (
	"sync"

	"github.com/delioda79/pbsbjaak/server"
)

type FakePubSubServer struct {
	SubscribeStub        func(topic, subscriber string) bool
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		topic      string
		subscriber string
	}
	subscribeReturns struct {
		result1 bool
	}
	subscribeReturnsOnCall map[int]struct {
		result1 bool
	}
	PublishStub        func(topic string, message []byte) int
	publishMutex       sync.RWMutex
	publishArgsForCall []struct {
		topic   string
		message []byte
	}
	publishReturns struct {
		result1 int
	}
	publishReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePubSubServer) Subscribe(topic string, subscriber string) bool {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		topic      string
		subscriber string
	}{topic, subscriber})
	fake.recordInvocation("Subscribe", []interface{}{topic, subscriber})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		return fake.SubscribeStub(topic, subscriber)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.subscribeReturns.result1
}

func (fake *FakePubSubServer) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakePubSubServer) SubscribeArgsForCall(i int) (string, string) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return fake.subscribeArgsForCall[i].topic, fake.subscribeArgsForCall[i].subscriber
}

func (fake *FakePubSubServer) SubscribeReturns(result1 bool) {
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePubSubServer) SubscribeReturnsOnCall(i int, result1 bool) {
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakePubSubServer) Publish(topic string, message []byte) int {
	var messageCopy []byte
	if message != nil {
		messageCopy = make([]byte, len(message))
		copy(messageCopy, message)
	}
	fake.publishMutex.Lock()
	ret, specificReturn := fake.publishReturnsOnCall[len(fake.publishArgsForCall)]
	fake.publishArgsForCall = append(fake.publishArgsForCall, struct {
		topic   string
		message []byte
	}{topic, messageCopy})
	fake.recordInvocation("Publish", []interface{}{topic, messageCopy})
	fake.publishMutex.Unlock()
	if fake.PublishStub != nil {
		return fake.PublishStub(topic, message)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.publishReturns.result1
}

func (fake *FakePubSubServer) PublishCallCount() int {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return len(fake.publishArgsForCall)
}

func (fake *FakePubSubServer) PublishArgsForCall(i int) (string, []byte) {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return fake.publishArgsForCall[i].topic, fake.publishArgsForCall[i].message
}

func (fake *FakePubSubServer) PublishReturns(result1 int) {
	fake.PublishStub = nil
	fake.publishReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakePubSubServer) PublishReturnsOnCall(i int, result1 int) {
	fake.PublishStub = nil
	if fake.publishReturnsOnCall == nil {
		fake.publishReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.publishReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakePubSubServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePubSubServer) recordInvocation(key string, args []interface{}) {
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

var _ server.PubSubServer = new(FakePubSubServer)
