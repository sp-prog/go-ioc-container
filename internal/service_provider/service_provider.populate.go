package service_provider

import (
	"errors"
	"reflect"
)

func (sp *ServiceProvider) Populate(target interface{}) error {
	return sp.PopulateWithExtras(
		target,
		(*ServiceMap)(nil).New(),
	)
}

func (sp *ServiceProvider) PopulateWithExtras(
	target interface{},
	extras ServiceMap,
) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().Kind() == reflect.Struct {
		targetValue = targetValue.Elem()
		for i := 0; i < targetValue.Type().NumField(); i++ {
			fieldVal := targetValue.Field(i)
			if fieldVal.CanSet() {
				if extra, ok := extras[fieldVal.Type()]; ok {
					fieldVal.Set(extra)
				} else {
					err = sp.resolveServiceFromValue(fieldVal.Addr())
				}
			}

		}
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
