package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
)

type TransientFactoryInfo struct {
	factoryFunc reflect.Value
	lifecycle   factory.Lifecycle
	objectType  reflect.Type
}

func (i *TransientFactoryInfo) ObjectType() reflect.Type {
	return i.objectType
}

func (i *TransientFactoryInfo) Lifecycle() factory.Lifecycle {
	return i.lifecycle
}

func (i *TransientFactoryInfo) FactoryFunc() reflect.Value {
	return i.factoryFunc
}

func (*TransientFactoryInfo) New(
	factoryFunc reflect.Value,
) *TransientFactoryInfo {
	objectType := factoryFunc.Type().Out(0)

	return &TransientFactoryInfo{
		factoryFunc: factoryFunc,
		lifecycle:   factory.Transient,
		objectType:  objectType,
	}
}

func (*TransientFactoryInfo) NewWithLifecycle(
	factoryFunc reflect.Value,
	lifecycle factory.Lifecycle,
) *TransientFactoryInfo {
	objectType := factoryFunc.Type().Out(0)

	return &TransientFactoryInfo{
		factoryFunc: factoryFunc,
		lifecycle:   lifecycle,
		objectType:  objectType,
	}
}
