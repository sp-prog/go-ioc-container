package call

// ICall Интерфейс вызова метода с заполнением части параметров через DI
type ICall interface {
	// Call Вызвать зарегистрованную ранее матод-фабрику
	Call(
		target interface{},
		otherArgs ...interface{},
	) ([]interface{}, error)
}
