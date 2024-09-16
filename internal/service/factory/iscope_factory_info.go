package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
)

type IScopeFactoryInfo interface {
	ObjectType() reflect.Type
	Lifecycle() factory.Lifecycle
	FactoryFunc() reflect.Value
	Copy() IScopeFactoryInfo
}
