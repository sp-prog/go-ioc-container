package collection

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
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

func (fi *FactoryInfo) Copy() factory2.IScopeFactoryInfo {
	return fi.New(
		fi.factoryFunc,
		fi.lifecycle,
		fi.objectType,
	)
}
