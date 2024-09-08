package provider

import "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"

// IProvider Интерфейс поставщика зависимостей
// Используется для создания объектов из ранее зарегистрированных зависимостей
type IProvider interface {
	// Build Собрать поставщика зависимостей из коллекции
	Build(sc collection.ICollection)

	// GetService Получить зависимость по типу
	GetService(target interface{}) error

	// Создать дочерний поставщик зависимостей с новой областью
	CreateScopedServiceProvider() IProvider
}
