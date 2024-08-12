package service_collection

import (
	"github.com/sp-prog/go-ioc-container/internal/service_collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"github.com/sp-prog/go-ioc-container/test/assert_extensions"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Проверка работы конструктора
func TestServiceCollectionNew(t *testing.T) {
	res := (*service_collection.ServiceCollection)(nil).New()

	assert.NotNil(t, res)
}

// Проверка метода получения списка завимимостей
func TestServiceCollectionServices(t *testing.T) {
	res := (*service_collection.ServiceCollection)(nil).New()

	assert.NotNil(t, res.Services())
}

// Проверка метода добавления зависимости с областью жизни Transient
func TestServiceCollectionAddTransient(t *testing.T) {
	//Test data
	f := func() string { return "" }

	sc := (*service_collection.ServiceCollection)(nil).New()

	//Action
	res := sc.AddTransient(f)
	fis := sc.Services().GetFactoryInfos()

	//Validate
	assert.NoError(t, res)
	assert_extensions.ContainsByFunction(
		t,
		fis,
		func(element *interfaces.FactoryInfo) bool {
			return element.FactoryFunc() == reflect.ValueOf(f) &&
				element.Lifecycle() == interfaces.Transient
		},
	)
}

// Проверка метода добавления зависимости с областью жизни Scoped
func TestServiceCollectionAddScoped(t *testing.T) {
	//Test data
	f := func() string { return "" }

	sc := (*service_collection.ServiceCollection)(nil).New()

	//Action
	res := sc.AddScoped(f)
	fis := sc.Services().GetFactoryInfos()

	//Validate
	assert.NoError(t, res)
	assert_extensions.ContainsByFunction(
		t,
		fis,
		func(element *interfaces.FactoryInfo) bool {
			return element.FactoryFunc() == reflect.ValueOf(f) &&
				element.Lifecycle() == interfaces.Scoped
		},
	)
}

// Проверка метода добавления зависимости с областью жизни Singleton
func TestServiceCollectionAddSingleton(t *testing.T) {
	//Test data
	f := func() string { return "" }

	sc := (*service_collection.ServiceCollection)(nil).New()

	//Action
	res := sc.AddSingleton(f)
	fis := sc.Services().GetFactoryInfos()

	//Validate
	assert.NoError(t, res)
	assert_extensions.ContainsByFunction(
		t,
		fis,
		func(element *interfaces.FactoryInfo) bool {
			return element.FactoryFunc() == reflect.ValueOf(f) &&
				element.Lifecycle() == interfaces.Singleton
		},
	)
}

// Проверка возникновения ошибки при попытке зарегистировать не-метод
func TestServiceCollectionNotFunction(t *testing.T) {
	//Test data
	var f struct{}

	sc := (*service_collection.ServiceCollection)(nil).New()

	//Action
	res := sc.AddSingleton(f)

	//Validate
	assert.Error(t, res)
}
