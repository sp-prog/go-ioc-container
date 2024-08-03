package interfaces

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
)

// IServiceProvider Интерфейс поставщика зависимостей
// Используется для создания объектов из ранее зарегистрированных зависимостей
type IServiceProvider interface {
	// Call Вызвать зарегистрованную ранее матод-фабрику
	Call(
		target interface{},
		otherArgs ...interface{},
	) ([]interface{}, error)

	// CallForContext Вызвать зарегистрованную ранее матод-фабрику из контекста
	CallForContext(
		c context.Context,
		target interface{},
		otherArgs ...interface{},
	) (results []interface{}, err error)

	// Populate Заполнить поля структуры зависимостями
	Populate(target interface{}) error

	// PopulateForContext Заполнить поля структуры зависимостями из контекста
	PopulateForContext(
		c context.Context,
		target interface{},
	) (err error)

	// Build Собрать поставщика зависимостей из коллекции
	Build(sc IServiceCollection)

	// GetService Получить зависимость по типу
	GetService(target interface{}) error

	// GetServiceForContext Получить зависимость по типу из контекста
	GetServiceForContext(
		c context.Context,
		target interface{},
	) (err error)

	// PopulateForContextWithExtras Заполнить поля структуры зависимостями из контекста и дополнительными параметрами
	PopulateForContextWithExtras(
		c context.Context,
		target interface{},
		extras type_factory.ServiceMap,
	) (err error)
}
