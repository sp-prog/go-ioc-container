// Package interfaces Интерфейсы контейнера инверсии управления
package interfaces

import "github.com/sp-prog/go-ioc-container/internal/type_factory"

// IServiceCollection Интерфейс коллекции зависимостей
// используется для сбора зависимостей, по которым
// в последствии будет собран поставщик зависимостей
type IServiceCollection interface {
	// Services Получить список зависимостей коллекции
	Services() type_factory.FactoryInfo

	// AddTransient Добавить зависимость со временем жизни "каждый раз новый"
	AddTransient(factoryFunc interface{}) (err error)

	// AddScoped Добавить зависимость со временем жизни "заданная область"
	AddScoped(factoryFunc interface{}) (err error)

	// AddSingleton Добавить зависимость со временем жизни "только один"
	AddSingleton(factoryFunc interface{}) (err error)
}
