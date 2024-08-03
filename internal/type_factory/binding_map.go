package type_factory

import (
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle   Lifecycle
}

func (b *BindingMap) Lifecycle() Lifecycle {
	return b.lifecycle
}

func (b *BindingMap) FactoryFunc() reflect.Value {
	return b.factoryFunc
}

func NewBindingMap(
	factoryFunc reflect.Value,
	lifecycle Lifecycle,
) BindingMap {
	return BindingMap{
		factoryFunc: factoryFunc,
		lifecycle:   lifecycle,
	}
}
