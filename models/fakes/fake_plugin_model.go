// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"github.com/cloudfoundry-incubator/cli-plugin-repo/models"
)

type FakePluginModel struct {
	PopulateModelStub        func(interface{})
	populateModelMutex       sync.RWMutex
	populateModelArgsForCall []struct {
		arg1 interface{}
	}
	PluginsModelStub        func() []models.Plugin
	pluginsModelMutex       sync.RWMutex
	pluginsModelArgsForCall []struct{}
	pluginsModelReturns struct {
		result1 []models.Plugin
	}
}

func (fake *FakePluginModel) PopulateModel(arg1 interface{}) {
	fake.populateModelMutex.Lock()
	defer fake.populateModelMutex.Unlock()
	fake.populateModelArgsForCall = append(fake.populateModelArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	if fake.PopulateModelStub != nil {
		fake.PopulateModelStub(arg1)
	}
}

func (fake *FakePluginModel) PopulateModelCallCount() int {
	fake.populateModelMutex.RLock()
	defer fake.populateModelMutex.RUnlock()
	return len(fake.populateModelArgsForCall)
}

func (fake *FakePluginModel) PopulateModelArgsForCall(i int) interface{} {
	fake.populateModelMutex.RLock()
	defer fake.populateModelMutex.RUnlock()
	return fake.populateModelArgsForCall[i].arg1
}

func (fake *FakePluginModel) PluginsModel() []models.Plugin {
	fake.pluginsModelMutex.Lock()
	defer fake.pluginsModelMutex.Unlock()
	fake.pluginsModelArgsForCall = append(fake.pluginsModelArgsForCall, struct{}{})
	if fake.PluginsModelStub != nil {
		return fake.PluginsModelStub()
	} else {
		return fake.pluginsModelReturns.result1
	}
}

func (fake *FakePluginModel) PluginsModelCallCount() int {
	fake.pluginsModelMutex.RLock()
	defer fake.pluginsModelMutex.RUnlock()
	return len(fake.pluginsModelArgsForCall)
}

func (fake *FakePluginModel) PluginsModelReturns(result1 []models.Plugin) {
	fake.PluginsModelStub = nil
	fake.pluginsModelReturns = struct {
		result1 []models.Plugin
	}{result1}
}

var _ models.PluginModel = new(FakePluginModel)
