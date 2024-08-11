package assert_extensions

import "github.com/stretchr/testify/assert"

// ContainsByFunction Проверка наличия элемента в списке
// вызовом функции, проверяющей условия равенства
// equalsFunction - пользовательская функция, проверяющая элемент списка
func ContainsByFunction[T any](
	t assert.TestingT,
	array []T,
	equalsFunction func(element T) bool,
) {
	for _, v := range array {
		if equalsFunction(v) {
			assert.True(t, true)

			return
		}
	}

	assert.True(t, false)
}
