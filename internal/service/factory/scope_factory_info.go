package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
	"sync"
)

type ScopeFactoryInfo struct {
	*TransientFactoryInfo
	once *sync.Once
}

func (*ScopeFactoryInfo) New(
	factoryFunc reflect.Value,
	lifecycle factory.Lifecycle,
) IScopeFactoryInfo {
	//хм. Возможно в этом нет ничего страшного,
	//т.к. результат живет и умирает вместе с объектом этой структуры
	var results []reflect.Value

	once := sync.Once{}
	wrapper := reflect.MakeFunc(
		factoryFunc.Type(),
		func(args []reflect.Value) []reflect.Value {
			once.Do(func() {
				results = factoryFunc.Call(args)
			})
			return results
		})

	//Разница устывается при создании объектов
	//по полю области жизни
	transientFactoryInfo := (*TransientFactoryInfo)(nil).NewWithLifecycle(
		wrapper,
		lifecycle,
	)

	return &ScopeFactoryInfo{
		TransientFactoryInfo: transientFactoryInfo,
		once:                 &once,
	}
}
