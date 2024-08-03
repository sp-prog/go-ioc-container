package service_provider

import (
	"fmt"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
)

// ServiceCollection Коллекция зарегистрированных в IoC-контейнере типов
type ServiceCollection struct {
	services type_factory.FactoryInfo
}

// NewServiceCollection Конструктор
func NewServiceCollection() interfaces.IServiceCollection {
	return &ServiceCollection{
		services: type_factory.NewFactoryInfo(),
	}
}

func (sc *ServiceCollection) Services() type_factory.FactoryInfo {
	return sc.services
}

func (sc *ServiceCollection) AddTransient(factoryFunc interface{}) (err error) {
	return sc.addService(type_factory.Transient, factoryFunc)
}

func (sc *ServiceCollection) AddScoped(factoryFunc interface{}) (err error) {
	return sc.addService(type_factory.Scoped, factoryFunc)
}

func (sc *ServiceCollection) AddSingleton(factoryFunc interface{}) (err error) {
	factoryFuncVal := reflect.ValueOf(factoryFunc)
	if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
		err = sc.addService(type_factory.Singleton, factoryFunc)
	}
	return
}

func (sc *ServiceCollection) addService(
	life type_factory.Lifecycle,
	factoryFunc interface{},
) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		sc.services[factoryFuncType.Out(0)] = type_factory.NewBindingMap(
			reflect.ValueOf(factoryFunc),
			life,
		)
	} else {
		err = fmt.Errorf("Type cannot be used as service: %v", factoryFuncType)

	}
	return
}
