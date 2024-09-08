package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"reflect"
)

type IScopeFactoryMap interface {
	New() IScopeFactoryMap
	SetFactoryInfo(factoryInfo *collection.FactoryInfo)
	GetFactoryInfo(target interface{}) (IScopeFactoryInfo, bool)
	GetFactoryInfoReflectType(reflectType reflect.Type) (IScopeFactoryInfo, bool)
}
