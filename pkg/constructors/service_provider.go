package constructors

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/service_collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
)

// NewServiceProvider Проксирование для скрытия реализации
// Создать новый поставщик зависимостей
func NewServiceProvider(
	contextReference *context.Context,
) interfaces.IServiceProvider {
	//contextReference = (*context.Context)(nil)
	//reflect.TypeOf(contextReference).Elem()
	return service_collection.NewServiceProvider(
		contextReference,
	)
}
