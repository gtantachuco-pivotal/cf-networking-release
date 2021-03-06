// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"service-discovery-controller/mbus"
	"sync"

	"github.com/nats-io/nats"
)

type NatsConnProvider struct {
	ConnectionStub        func(opts ...nats.Option) (mbus.NatsConn, error)
	connectionMutex       sync.RWMutex
	connectionArgsForCall []struct {
		opts []nats.Option
	}
	connectionReturns struct {
		result1 mbus.NatsConn
		result2 error
	}
	connectionReturnsOnCall map[int]struct {
		result1 mbus.NatsConn
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NatsConnProvider) Connection(opts ...nats.Option) (mbus.NatsConn, error) {
	fake.connectionMutex.Lock()
	ret, specificReturn := fake.connectionReturnsOnCall[len(fake.connectionArgsForCall)]
	fake.connectionArgsForCall = append(fake.connectionArgsForCall, struct {
		opts []nats.Option
	}{opts})
	fake.recordInvocation("Connection", []interface{}{opts})
	fake.connectionMutex.Unlock()
	if fake.ConnectionStub != nil {
		return fake.ConnectionStub(opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.connectionReturns.result1, fake.connectionReturns.result2
}

func (fake *NatsConnProvider) ConnectionCallCount() int {
	fake.connectionMutex.RLock()
	defer fake.connectionMutex.RUnlock()
	return len(fake.connectionArgsForCall)
}

func (fake *NatsConnProvider) ConnectionArgsForCall(i int) []nats.Option {
	fake.connectionMutex.RLock()
	defer fake.connectionMutex.RUnlock()
	return fake.connectionArgsForCall[i].opts
}

func (fake *NatsConnProvider) ConnectionReturns(result1 mbus.NatsConn, result2 error) {
	fake.ConnectionStub = nil
	fake.connectionReturns = struct {
		result1 mbus.NatsConn
		result2 error
	}{result1, result2}
}

func (fake *NatsConnProvider) ConnectionReturnsOnCall(i int, result1 mbus.NatsConn, result2 error) {
	fake.ConnectionStub = nil
	if fake.connectionReturnsOnCall == nil {
		fake.connectionReturnsOnCall = make(map[int]struct {
			result1 mbus.NatsConn
			result2 error
		})
	}
	fake.connectionReturnsOnCall[i] = struct {
		result1 mbus.NatsConn
		result2 error
	}{result1, result2}
}

func (fake *NatsConnProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.connectionMutex.RLock()
	defer fake.connectionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NatsConnProvider) recordInvocation(key string, args []interface{}) {
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

var _ mbus.NatsConnProvider = new(NatsConnProvider)
