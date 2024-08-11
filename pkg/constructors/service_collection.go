package constructors

import (
	"github.com/sp-prog/go-ioc-container/internal/service_collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
)

// NewServiceCollection Проксирование для скрытия реализации
// Создать новую коллекцию зависимостей, по которой в последствии необходимо
// собрать поставщика зависимостей
func NewServiceCollection() interfaces.IServiceCollection {
	return (*service_collection.ServiceCollection)(nil).New()
}
