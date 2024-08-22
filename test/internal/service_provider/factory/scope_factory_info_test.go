package factory

import (
	"reflect"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type contextMocked struct {
	mock.Mock
}

func (cm *contextMocked) fakeFunc() string {
	cm.Called()

	return ""
}

func TestScopeFactoryInfoAndNewThenCreated(t *testing.T) {
	//Test data

	//Action
	res := (*factory.ScopeFactoryInfo)(nil).New(
		reflect.ValueOf(func() string {
			return ""
		}),
		interfaces.Singleton,
		reflect.TypeOf(""),
	)

	//Validate
	assert.NotNil(t, res)
}

// Проверка конструктора и многопоточного вызова метода-конструктора зависимости
func TestScopeFactoryInfoAndGoCallThenCallOnce(t *testing.T) {
	//Test data
	funcName := "fakeFunc"
	grCount := gofakeit.Number(1, 1000)
	wg := sync.WaitGroup{}

	m := contextMocked{}

	sc := (*factory.ScopeFactoryInfo)(nil).New(
		reflect.ValueOf(m.fakeFunc),
		interfaces.Singleton,
		reflect.TypeOf(""),
	)

	m.On(funcName).Return("")

	//Action
	wg.Add(grCount)
	for i := 0; i < grCount; i++ {
		go func() {
			sc.FactoryFunc().Call(nil)

			wg.Done()
		}()
	}

	wg.Wait()

	//Validate
	m.AssertNumberOfCalls(t, funcName, 1)
}
