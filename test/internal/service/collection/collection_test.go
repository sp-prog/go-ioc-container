package collection

import (
	"github.com/sp-prog/go-ioc-container/internal/service/collection"
	collection2 "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	assert2 "github.com/sp-prog/go-ioc-container/test/extensions/assert"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Проверка работы конструктора
func TestCollectionAndNewThenCreated(t *testing.T) {
	res := (*collection.Collection)(nil).New()

	assert.NotNil(t, res)
}

// Проверка метода получения списка завимимостей
func TestCollectionAndServicesThenNotNull(t *testing.T) {
	res := (*collection.Collection)(nil).New()

	assert.NotNil(t, res.Services())
}

// Проверка метода добавления зависимости с областью жизни Transient
func TestCollectionAndAddLifecycleThenExists(t *testing.T) {
	t.Parallel()

	//Test data
	f := func() string { return "" }
	sc := (*collection.Collection)(nil).New()

	testDatas := []struct {
		name          string
		lifecycleFunc func(interface{}) error
		lifecycle     factory.Lifecycle
	}{
		{"test AddTransient lifecycle", sc.AddTransient, factory.Transient},
		{"test AddScoped lifecycle", sc.AddScoped, factory.Scoped},
		{"test AddSingleton lifecycle", sc.AddSingleton, factory.Singleton},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			//Action
			res := testData.lifecycleFunc(f)
			fis := sc.Services().GetFactoryInfos()

			//Validate
			assert.NoError(t, res, testData.name)
			assert2.ContainsByFunctionf(
				t,
				fis,
				func(element *collection2.FactoryInfo) bool {
					return element.FactoryFunc() == reflect.ValueOf(f) &&
						element.Lifecycle() == testData.lifecycle
				},
				testData.name,
			)
		})
	}
}

// Проверка возникновения ошибки при попытке зарегистировать не-метод
func TestCollectionAndAndWithNotFunctionThenError(t *testing.T) {
	//Test data
	var f struct{}

	sc := (*collection.Collection)(nil).New()

	//Action
	res := sc.AddSingleton(f)

	//Validate
	assert.Error(t, res)
}
