package provider

import (
	"reflect"
)

type ServiceMap map[reflect.Type]reflect.Value

func (*ServiceMap) New() ServiceMap {
	return make(ServiceMap)
}
