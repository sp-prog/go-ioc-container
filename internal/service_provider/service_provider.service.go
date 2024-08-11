package service_provider

import (
	"errors"
	"fmt"
	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
)

func (sp *ServiceProvider) Build(sc interfaces.IServiceCollection) {
	services := (*factory.ScopeFactoryMap)(nil).New()

	serviceCollection := sc.Services().GetFactoryInfos()

	for _, v := range serviceCollection {
		services.SetFactoryInfo(
			v,
		)
	}

	sp.scopeFactoryMap = services
}

func (sp *ServiceProvider) GetService(target interface{}) (err error) {
	if fi, exists := sp.scopeFactoryMap.GetFactoryInfo(target); exists && fi.Lifecycle() == interfaces.Singleton {
		return sp.rootProvider.GetService(target)
	}

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = sp.resolveServiceFromValue(targetValue)
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}

func (sp *ServiceProvider) resolveServiceFromValue(
	val reflect.Value,
) (err error) {
	serviceType := val.Elem().Type()
	if factoryInfo, found := sp.scopeFactoryMap.GetFactoryInfoReflectType(serviceType); found {
		if factoryInfo.Lifecycle() == interfaces.Scoped {
			err = sp.resolveScopedService(val, factoryInfo)
		} else {
			val.Elem().Set(sp.invokeFunction(factoryInfo.FactoryFunc())[0])
		}
	} else {
		err = fmt.Errorf("Cannot find service %v", serviceType)
	}
	return
}

func (sp *ServiceProvider) resolveScopedService(
	val reflect.Value,
	factoryInfo factory.IScopeFactoryInfo,
) (err error) {
	sMap := sp.serviceMap
	serviceVal, ok := sMap[val.Type()]
	if !ok {
		serviceVal = sp.invokeFunction(factoryInfo.FactoryFunc())[0]
		sMap[val.Type()] = serviceVal
	}
	val.Elem().Set(serviceVal)

	return nil
}

func (sp *ServiceProvider) resolveFunctionArguments(
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
		err := sp.resolveServiceFromValue(pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}

func (sp *ServiceProvider) invokeFunction(
	f reflect.Value,
	otherArgs ...interface{},
) []reflect.Value {
	return f.Call(sp.resolveFunctionArguments(f, otherArgs...))
}
