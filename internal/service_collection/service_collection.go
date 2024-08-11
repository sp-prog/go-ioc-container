package service_collection

import (
	"fmt"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
)

// ServiceCollection Коллекция зарегистрированных в IoC-контейнере типов
type ServiceCollection struct {
	factoryMap *interfaces.FactoryMap
}

// New NewServiceCollection Конструктор,
// который можно вызывать через пустую ссылку, т.е. вот так:
// (*service_collection.ServiceCollection)(nil).New()
func (*ServiceCollection) New() interfaces.IServiceCollection {
	return &ServiceCollection{
		factoryMap: (*interfaces.FactoryMap)(nil).New(),
	}
}

func (sc *ServiceCollection) Services() *interfaces.FactoryMap {
	return sc.factoryMap
}

func (sc *ServiceCollection) AddTransient(factoryFunc interface{}) (err error) {
	return sc.addService(interfaces.Transient, factoryFunc)
}

func (sc *ServiceCollection) AddScoped(factoryFunc interface{}) (err error) {
	return sc.addService(interfaces.Scoped, factoryFunc)
}

func (sc *ServiceCollection) AddSingleton(factoryFunc interface{}) (err error) {
	return sc.addService(interfaces.Singleton, factoryFunc)
}

func (sc *ServiceCollection) addService(
	life interfaces.Lifecycle,
	factoryFunc interface{},
) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		sc.factoryMap.SetFactoryInfo(
			factoryFuncType.Out(0),
			(*interfaces.FactoryInfo)(nil).New(
				reflect.ValueOf(factoryFunc),
				life,
				factoryFuncType.Out(0),
			),
		)
	} else {
		err = fmt.Errorf("Type cannot be used as service: %v", factoryFuncType)

	}
	return
}
