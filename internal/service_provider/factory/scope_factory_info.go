package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
	"sync"
)

type ScopeFactoryInfo struct {
	TransientFactoryInfo
	once *sync.Once
}

func (*ScopeFactoryInfo) New(
	factoryFunc reflect.Value,
	lifecycle interfaces.Lifecycle,
	objectType reflect.Type,
) IScopeFactoryInfo {

	var results []reflect.Value
	once := sync.Once{}
	wrapper := reflect.MakeFunc(factoryFunc.Type(),
		func(args []reflect.Value) []reflect.Value {
			once.Do(func() {
				results = factoryFunc.Call(args)
			})
			return results
		})

	transientFactoryInfo := TransientFactoryInfo{
		factoryFunc: wrapper,
		lifecycle:   lifecycle,
		objectType:  objectType,
	}

	return &ScopeFactoryInfo{
		TransientFactoryInfo: transientFactoryInfo,
		once:                 &once,
	}
}
