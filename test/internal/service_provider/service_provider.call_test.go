package service_provider

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/sp-prog/go-ioc-container/internal/service_collection"
	"github.com/sp-prog/go-ioc-container/internal/service_provider"
	"github.com/stretchr/testify/assert"
)

// Проверка работы метода вызова функции с ошибочным параметром
func TestServiceProviderCallAndNotFunctionThenError(t *testing.T) {
	//Test data
	sp := (*service_provider.ServiceProvider)(nil).New()

	//Action
	res, err := sp.Call(struct{}{})

	//Validate
	assert.Nil(t, res)
	assert.Error(t, err)
}

// Проверка работы метода вызова функции без аргументов
func TestServiceProviderCallAndNoArgsThenOk(t *testing.T) {
	//Test data
	id := uuid.NewString()

	sp := (*service_provider.ServiceProvider)(nil).New()

	//Action
	res, err := sp.Call(func() string { return id })

	//Validate
	assert.NoError(t, err)
	assert.Equal(t, id, res[0])
}

// Проверка работы метода вызова функции с незарегистрированной зависимостью
func TestServiceProviderCallAndArgsAndNoDependencyThenError(t *testing.T) {
	//Test data
	id := uuid.NewString()
	f := func(i int) string { return id }

	sp := (*service_provider.ServiceProvider)(nil).New()

	//Validate
	assert.PanicsWithError(t, "Cannot find service int", func() {
		_, _ = sp.Call(f)
	})
}

// Проверка работы метода вызова функции с зарегистрированной зависимостью
func TestServiceProviderCallAndArgsThenError(t *testing.T) {
	t.Parallel()

	//Test data
	id := uuid.NewString()
	arg := gofakeit.Number(1, 1000)

	funcArg := func() int { return arg }
	funcCall := func(i int) string { return id }

	sc := (*service_collection.ServiceCollection)(nil).New()

	testDatas := []struct {
		name          string
		lifecycleFunc func(interface{}) error
	}{
		{"test Call with Transient lifecycle", sc.AddTransient},
		{"test Call with Scoped lifecycle", sc.AddScoped},
		{"test Call with Singleton lifecycle", sc.AddSingleton},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			sp := (*service_provider.ServiceProvider)(nil).New()

			testData.lifecycleFunc(funcArg)

			sp.Build(sc)

			res, err := sp.Call(funcCall)

			//Validate
			assert.NoError(t, err)
			assert.Equal(t, id, res[0])
		})
	}
}
