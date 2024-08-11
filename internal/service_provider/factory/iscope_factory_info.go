package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
)

type IScopeFactoryInfo interface {
	ObjectType() reflect.Type
	Lifecycle() interfaces.Lifecycle
	FactoryFunc() reflect.Value
}
