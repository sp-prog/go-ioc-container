package type_factory

import (
	"reflect"
)

type ServiceMap map[reflect.Type]reflect.Value

func NewServiceMap() ServiceMap {
	return make(ServiceMap)
}
