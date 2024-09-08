package service

import (
	"github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	iprovider "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/provider"
)

// NewProvider Проксирование для скрытия реализации
// Создать новый поставщик зависимостей
func NewProvider() iprovider.IProvider {
	sfm := (*factory.ScopeFactoryMap)(nil).New()

	return (*provider.Provider)(nil).New(sfm)
}
