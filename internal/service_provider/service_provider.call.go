package service_provider

import (
	"errors"
	"reflect"
)

func (sp *ServiceProvider) Call(
	target interface{},
	otherArgs ...interface{},
) (results []interface{}, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultVals := sp.invokeFunction(targetValue, otherArgs...)
		results = make([]interface{}, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			results[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("Only functions can be invoked")
	}
	return
}
