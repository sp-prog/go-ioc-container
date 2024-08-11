package service_provider

import (
	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
)

// ServiceProvider Поставщик зависимостей
// Для создания новой области необходимо создать новый контекст,
// например, из существующего
type ServiceProvider struct {
	serviceMap      ServiceMap
	scopeFactoryMap *factory.ScopeFactoryMap
	rootProvider    interfaces.IServiceProvider
}

// NewServiceProvider Конструктор
func (*ServiceProvider) New() interfaces.IServiceProvider {

	return &ServiceProvider{
		serviceMap:      (*ServiceMap)(nil).New(),
		scopeFactoryMap: (*factory.ScopeFactoryMap)(nil).New(),
		rootProvider:    nil,
	}
}

func (sp *ServiceProvider) CreateScopedServiceProvider() interfaces.IServiceProvider {
	return &ServiceProvider{
		serviceMap:      (*ServiceMap)(nil).New(),
		scopeFactoryMap: sp.scopeFactoryMap,
		rootProvider:    sp,
	}
}
