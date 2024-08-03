package service_collection

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
)

// ServiceProvider Поставщик зависимостей
// Для создания новой области необходимо создать новый контекст,
// например, из существующего
type ServiceProvider struct {
	contextReference     *context.Context
	contextReferenceType reflect.Type
	services             type_factory.FactoryInfo
}

// NewServiceProvider Конструктор
func NewServiceProvider(
	contextReference *context.Context,
) interfaces.IServiceProvider {
	//contextReference = (*context.Context)(nil)
	return &ServiceProvider{
		contextReference:     contextReference,
		contextReferenceType: reflect.TypeOf(contextReference).Elem(),
		services:             type_factory.NewFactoryInfo(),
	}
}
