package collection

import (
	"fmt"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
)

// Collection Коллекция зарегистрированных в IoC-контейнере типов
type Collection struct {
	factoryMap *collection.FactoryMap
}

// New NewServiceCollection Конструктор,
// который можно вызывать через пустую ссылку, т.е. вот так:
// (*collection.Collection)(nil).New()
func (*Collection) New() collection.ICollection {
	return &Collection{
		factoryMap: (*collection.FactoryMap)(nil).New(),
	}
}

func (sc *Collection) Services() *collection.FactoryMap {
	return sc.factoryMap
}

func (sc *Collection) AddTransient(factoryFunc interface{}) (err error) {
	return sc.addService(factory.Transient, factoryFunc)
}

func (sc *Collection) AddScoped(factoryFunc interface{}) (err error) {
	return sc.addService(factory.Scoped, factoryFunc)
}

func (sc *Collection) AddSingleton(factoryFunc interface{}) (err error) {
	return sc.addService(factory.Singleton, factoryFunc)
}

func (sc *Collection) addService(
	life factory.Lifecycle,
	factoryFunc interface{},
) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		sc.factoryMap.SetFactoryInfo(
			factoryFuncType.Out(0),
			(*collection.FactoryInfo)(nil).New(
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
