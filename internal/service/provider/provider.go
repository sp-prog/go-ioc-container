package provider

import (
	"errors"
	"fmt"
	"github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	pkgFactory "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/provider"
	"reflect"
)

// Provider Поставщик зависимостей
// Для создания новой области необходимо создать новый контекст,
// например, из существующего
type Provider struct {
	serviceMap      ServiceMap
	scopeFactoryMap factory.IScopeFactoryMap
	rootProvider    provider.IProvider
}

// New Конструктор
func (*Provider) New(
	scopeFactoryMap factory.IScopeFactoryMap,
) provider.IProvider {

	return &Provider{
		serviceMap:      (*ServiceMap)(nil).New(),
		scopeFactoryMap: scopeFactoryMap,
		rootProvider:    nil,
	}
}

func (p *Provider) CreateScopedServiceProvider() provider.IProvider {

	return &Provider{
		serviceMap:      p.serviceMap.New(),
		scopeFactoryMap: p.scopeFactoryMap.Copy(),
		rootProvider:    p,
	}
}

func (p *Provider) Build(c collection.ICollection) {
	services := p.scopeFactoryMap.New()

	fs := c.Services().GetFactoryInfos()

	for _, v := range fs {
		services.SetFactoryInfo(
			v,
		)
	}

	p.scopeFactoryMap = services
}

func (p *Provider) GetService(target interface{}) (err error) {
	fi, exists := p.scopeFactoryMap.GetFactoryInfo(target)
	if !exists {
		err = errors.New("target factory not found")

		return
	}

	if fi.Lifecycle() == pkgFactory.Singleton &&
		p.rootProvider != nil {
		return p.rootProvider.GetService(target)
	}

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = p.resolveServiceFromValue(targetValue)
	} else {
		err = errors.New("type cannot be used as target")
	}

	return
}

func (p *Provider) resolveServiceFromValue(
	val reflect.Value,
) (err error) {
	serviceType := val.Elem().Type()
	if factoryInfo, found := p.scopeFactoryMap.GetFactoryInfoReflectType(serviceType); found {
		if factoryInfo.Lifecycle() != pkgFactory.Transient {
			err = p.resolveScopedService(val, factoryInfo)
		} else {
			val.Elem().Set(p.invokeFunction(factoryInfo.FactoryFunc())[0])
		}
	} else {
		err = fmt.Errorf("cannot find service %v", serviceType)
	}
	return
}

func (p *Provider) resolveScopedService(
	val reflect.Value,
	factoryInfo factory.IScopeFactoryInfo,
) (err error) {
	sMap := p.serviceMap
	serviceVal, ok := sMap[val.Type()]
	if !ok {
		serviceVal = p.invokeFunction(factoryInfo.FactoryFunc())[0]
		sMap[val.Type()] = serviceVal
	}
	val.Elem().Set(serviceVal)

	return nil
}

func (p *Provider) resolveFunctionArguments(
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
		err := p.resolveServiceFromValue(pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}

func (p *Provider) invokeFunction(
	f reflect.Value,
	otherArgs ...interface{},
) []reflect.Value {
	return f.Call(p.resolveFunctionArguments(f, otherArgs...))
}
