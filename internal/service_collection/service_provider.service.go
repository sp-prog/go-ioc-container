package service_collection

import (
	"context"
	"errors"
	"fmt"
	"github.com/sp-prog/go-ioc-container/internal/extensions"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
	"sync"
)

func (sp *ServiceProvider) Build(sc interfaces.IServiceCollection) {
	services := type_factory.NewFactoryInfo()

	for k, v := range sc.Services() {
		if v.Lifecycle() == type_factory.Singleton {
			var results []reflect.Value
			once := sync.Once{}
			wrapper := reflect.MakeFunc(v.FactoryFunc().Type(),
				func([]reflect.Value) []reflect.Value {
					once.Do(func() {
						results = sp.invokeFunction(nil, v.FactoryFunc())
					})
					return results
				})

			factoryFuncType := reflect.TypeOf(wrapper)
			services[factoryFuncType.Out(0)] = type_factory.NewBindingMap(
				reflect.ValueOf(wrapper),
				v.Lifecycle(),
			)
		} else {
			services[k] = v
		}
	}

	sp.services = services
}

func (sp *ServiceProvider) GetService(target interface{}) error {
	return sp.GetServiceForContext(context.Background(), target)
}

func (sp *ServiceProvider) GetServiceForContext(
	c context.Context,
	target interface{},
) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = sp.resolveServiceFromValue(c, targetValue)
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}

func (sp *ServiceProvider) PopulateForContextWithExtras(
	c context.Context,
	target interface{},
	extras type_factory.ServiceMap,
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
					err = sp.resolveServiceFromValue(c, fieldVal.Addr())
				}
			}

		}
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}

func (sp *ServiceProvider) resolveServiceFromValue(
	c context.Context,
	val reflect.Value,
) (err error) {
	serviceType := val.Elem().Type()
	if serviceType == sp.contextReferenceType {
		val.Elem().Set(reflect.ValueOf(c))
	} else if binding, found := sp.services[serviceType]; found {
		if binding.Lifecycle() == type_factory.Scoped {
			err = sp.resolveScopedService(c, val, binding)
		} else {
			val.Elem().Set(sp.invokeFunction(c, binding.FactoryFunc())[0])
		}
	} else {
		err = fmt.Errorf("Cannot find service %v", serviceType)
	}
	return
}

func (sp *ServiceProvider) resolveScopedService(
	c context.Context,
	val reflect.Value,
	binding type_factory.BindingMap,
) (err error) {
	sMap, ok := c.Value(extensions.ServiceKey).(type_factory.ServiceMap)
	if ok {
		serviceVal, ok := sMap[val.Type()]
		if !ok {
			serviceVal = sp.invokeFunction(c, binding.FactoryFunc())[0]
			sMap[val.Type()] = serviceVal
		}
		val.Elem().Set(serviceVal)
	} else {
		val.Elem().Set(sp.invokeFunction(c, binding.FactoryFunc())[0])
	}
	return
}

func (sp *ServiceProvider) resolveFunctionArguments(
	c context.Context,
	f reflect.Value,
	otherArgs ...interface{},
) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	if otherArgs != nil {
		for ; i < len(otherArgs); i++ {
			params[i] = reflect.ValueOf(otherArgs[i])
		}
	}
	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		err := sp.resolveServiceFromValue(c, pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}

func (sp *ServiceProvider) invokeFunction(
	c context.Context,
	f reflect.Value,
	otherArgs ...interface{},
) []reflect.Value {
	return f.Call(sp.resolveFunctionArguments(c, f, otherArgs...))
}
