package collection

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
)

type FactoryInfo struct {
	factoryFunc reflect.Value
	lifecycle   factory.Lifecycle
	objectType  reflect.Type
}

func (*FactoryInfo) New(
	factoryFunc reflect.Value,
	lifecycle factory.Lifecycle,
	objectType reflect.Type,
) *FactoryInfo {
	return &FactoryInfo{
		factoryFunc: factoryFunc,
		lifecycle:   lifecycle,
		objectType:  objectType,
	}
}

func (fi *FactoryInfo) ObjectType() reflect.Type {
	return fi.objectType
}

func (fi *FactoryInfo) Lifecycle() factory.Lifecycle {
	return fi.lifecycle
}

func (fi *FactoryInfo) FactoryFunc() reflect.Value {
	return fi.factoryFunc
}
