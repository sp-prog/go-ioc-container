package type_factory

import "reflect"

type FactoryInfo map[reflect.Type]BindingMap

func NewFactoryInfo() FactoryInfo {
	return make(FactoryInfo)
}
