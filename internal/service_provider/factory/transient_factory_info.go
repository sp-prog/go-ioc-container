package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
)

type TransientFactoryInfo struct {
	factoryFunc reflect.Value
	lifecycle   interfaces.Lifecycle
	objectType  reflect.Type
}

func (i *TransientFactoryInfo) ObjectType() reflect.Type {
	return i.objectType
}

func (i *TransientFactoryInfo) Lifecycle() interfaces.Lifecycle {
	return i.lifecycle
}

func (i *TransientFactoryInfo) FactoryFunc() reflect.Value {
	return i.factoryFunc
}

func (*TransientFactoryInfo) New(
	factoryFunc reflect.Value,
	lifecycle interfaces.Lifecycle,
	objectType reflect.Type,
) IScopeFactoryInfo {
	return &TransientFactoryInfo{
		factoryFunc: factoryFunc,
		lifecycle:   lifecycle,
		objectType:  objectType,
	}
}
