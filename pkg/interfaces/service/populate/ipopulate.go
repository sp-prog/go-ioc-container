package populate

// IPopulate Интерфейс для заполнения полей структуры через DI
type IPopulate interface {
	// Populate Заполнить поля структуры зависимостями
	Populate(target interface{}) error
}
