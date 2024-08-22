package service_collection

import (
	"reflect"
	"testing"

	"github.com/sp-prog/go-ioc-container/internal/service_collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"github.com/sp-prog/go-ioc-container/test/assert_extensions"
	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestServiceCollectionAndNewThenCreated(t *testing.T) {
	res := (*service_collection.ServiceCollection)(nil).New()

	assert.NotNil(t, res)
}

// Проверка метода получения списка завимимостей
func TestServiceCollectionAndServicesThenNotNull(t *testing.T) {
	res := (*service_collection.ServiceCollection)(nil).New()

	assert.NotNil(t, res.Services())
}

// Проверка метода добавления зависимости с областью жизни Transient
func TestServiceCollectionAndAddLifecycleThenExists(t *testing.T) {
	t.Parallel()

	//Test data
	f := func() string { return "" }
	sc := (*service_collection.ServiceCollection)(nil).New()

	testDatas := []struct {
		name          string
		lifecycleFunc func(interface{}) error
		lifecycle     interfaces.Lifecycle
	}{
		{"test AddTransient lifecycle", sc.AddTransient, interfaces.Transient},
		{"test AddScoped lifecycle", sc.AddScoped, interfaces.Scoped},
		{"test AddSingleton lifecycle", sc.AddSingleton, interfaces.Singleton},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			//Action
			res := testData.lifecycleFunc(f)
			fis := sc.Services().GetFactoryInfos()

			//Validate
			assert.NoError(t, res, testData.name)
			assert_extensions.ContainsByFunctionf(
				t,
				fis,
				func(element *interfaces.FactoryInfo) bool {
					return element.FactoryFunc() == reflect.ValueOf(f) &&
						element.Lifecycle() == testData.lifecycle
				},
				testData.name,
			)
		})
	}
}

// Проверка возникновения ошибки при попытке зарегистировать не-метод
func TestServiceCollectionAndAndWithhNotFunctionThenError(t *testing.T) {
	//Test data
	var f struct{}

	sc := (*service_collection.ServiceCollection)(nil).New()

	//Action
	res := sc.AddSingleton(f)

	//Validate
	assert.Error(t, res)
}
