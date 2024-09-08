package factory

import (
	"github.com/sp-prog/go-ioc-container/internal/service/factory"
	icollection "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	factory2 "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify"
	mock2 "github.com/sp-prog/go-ioc-container/test/extensions/testify/mock"
	"reflect"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestScopeFactoryInfoAndNewThenCreated(t *testing.T) {
	//Test data

	//Action
	res := (*factory.ScopeFactoryInfo)(nil).New(
		reflect.ValueOf(func() string {
			return ""
		}),
		factory2.Singleton,
	)

	//Validate
	assert.NotNil(t, res)
}

// Проверка конструктора и многопоточного вызова метода-конструктора зависимости
func TestScopeFactoryInfoAndGoCallThenCallOnce(t *testing.T) {
	//Test data
	grCount := gofakeit.Number(1, 1000)
	wg := sync.WaitGroup{}

	cm := contextMocked{}

	sc := (*factory.ScopeFactoryInfo)(nil).New(
		reflect.ValueOf(cm.fakeFunc),
		factory2.Singleton,
	)

	mock2.MockR1[*icollection.FactoryMap]{Mock: &cm.Mock}.
		OnExt(cm.fakeFunc).
		Return("")

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
	testify.Assert{Mock: &cm.Mock}.
		AssertNumberOfCallsEx(t, cm.fakeFunc, 1)
}
