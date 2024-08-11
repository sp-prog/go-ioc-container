package interfaces

// IServiceProvider Интерфейс поставщика зависимостей
// Используется для создания объектов из ранее зарегистрированных зависимостей
type IServiceProvider interface {
	// Call Вызвать зарегистрованную ранее матод-фабрику
	Call(
		target interface{},
		otherArgs ...interface{},
	) ([]interface{}, error)

	// Populate Заполнить поля структуры зависимостями
	Populate(target interface{}) error

	// Build Собрать поставщика зависимостей из коллекции
	Build(sc IServiceCollection)

	// GetService Получить зависимость по типу
	GetService(target interface{}) error

	// Создать дочерний поставщик зависимостей с новой областью
	CreateScopedServiceProvider() IServiceProvider
}
