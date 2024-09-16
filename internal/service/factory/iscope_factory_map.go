package factory

import (
	"reflect"
)

type IScopeFactoryMap interface {
	New() IScopeFactoryMap
	SetFactoryInfo(factoryInfo IScopeFactoryInfo)
	GetFactoryInfo(target interface{}) (IScopeFactoryInfo, bool)
	GetFactoryInfoReflectType(reflectType reflect.Type) (IScopeFactoryInfo, bool)
	Copy() IScopeFactoryMap
}
