package constructors

import (
	"github.com/sp-prog/go-ioc-container/internal/service_provider"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
)

// NewServiceProvider Проксирование для скрытия реализации
// Создать новый поставщик зависимостей
func NewServiceProvider() interfaces.IServiceProvider {
	return (*service_provider.ServiceProvider)(nil).New()
}
