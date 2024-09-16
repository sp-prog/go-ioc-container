package Build

import (
	"github.com/brianvoe/gofakeit"
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify"
	mock2 "github.com/sp-prog/go-ioc-container/test/extensions/testify/mock"
	"github.com/stretchr/testify/mock"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type scopeFactoryMapMocked struct {
	mock.Mock
	factory2.IScopeFactoryMap
}

func (sfmm *scopeFactoryMapMocked) GetFactoryInfo(target interface{}) (factory2.IScopeFactoryInfo, bool) {
	args := sfmm.Called()

	return args.Get(0).(factory2.IScopeFactoryInfo), args.Bool(1)
}

func (sfmm *scopeFactoryMapMocked) GetFactoryInfoReflectType(reflectType reflect.Type) (factory2.IScopeFactoryInfo, bool) {
	args := sfmm.Called([]interface{}{reflectType})

	return args.Get(0).(factory2.IScopeFactoryInfo), args.Bool(1)
}

func (sfmm *scopeFactoryMapMocked) Copy() factory2.IScopeFactoryMap {
	args := sfmm.Called()

	return args.Get(0).(factory2.IScopeFactoryMap)
}

// Проверка работы метода получения объекта при отсутствии зависимости
func TestProviderAndGetServiceAndNotExistsThenError(t *testing.T) {
	//Test data
	sfmm := scopeFactoryMapMocked{}
	p := (*provider.Provider)(nil).New(&sfmm)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmm.Mock}).
		OnExt(sfmm.GetFactoryInfo).
		ReturnExt((*factory2.ScopeFactoryInfo)(nil), false)

	//Action
	err := p.GetService("")

	//Validate
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "target factory not found")
}

// Проверка работы метода получения объекта из родительской области при передаче неуказателя
func TestProviderAndGetServiceAndSingletonThenError(t *testing.T) {
	//Test data
	str := ""
	f := func() string { return str }

	sfi := (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(f), factory.Singleton)

	sfmmGlobal := scopeFactoryMapMocked{}
	sfmmScoped := scopeFactoryMapMocked{}

	p := (*provider.Provider)(nil).New(&sfmmGlobal)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmmGlobal.Mock}).
		OnExt(sfmmGlobal.GetFactoryInfo).
		ReturnExt(sfi, true)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmmScoped.Mock}).
		OnExt(sfmmScoped.GetFactoryInfo).
		ReturnExt(sfi, true)

	(&mock2.MockR1[factory2.IScopeFactoryMap]{Mock: &sfmmGlobal.Mock}).
		OnExt(sfmmGlobal.Copy).
		ReturnExt(factory2.IScopeFactoryMap(&sfmmScoped))

	//Action
	ssp := p.CreateScopedServiceProvider()
	err := ssp.GetService(str)

	//Validate
	testify.Assert{Mock: &sfmmGlobal.Mock}.
		AssertNumberOfCallsEx(t, sfmmGlobal.GetFactoryInfo, 1)
	testify.Assert{Mock: &sfmmScoped.Mock}.
		AssertNumberOfCallsEx(t, sfmmScoped.GetFactoryInfo, 1)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "type cannot be used as target")
}

// Проверка работы метода разрешения зависимостей при отсуствии зависимости`
func TestProviderAndGetServiceAndNoFactoryThenError(t *testing.T) {
	//Test data
	str := ""
	strValue := reflect.ValueOf(&str)
	strType := strValue.Elem().Type()

	fStr := func(int) string { return str }

	i := 0
	fi := i
	intValue := reflect.ValueOf(&fi)
	intType := intValue.Elem().Type()

	sfi := (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(fStr), factory.Scoped)

	sfmm := scopeFactoryMapMocked{}

	p := (*provider.Provider)(nil).New(&sfmm)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmm.Mock}).
		OnExt(sfmm.GetFactoryInfo).
		ReturnExt(sfi, true)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmm.Mock}).
		OnExt(sfmm.GetFactoryInfoReflectType, strType).
		ReturnExt(sfi, true)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmm.Mock}).
		OnExt(sfmm.GetFactoryInfoReflectType, intType).
		ReturnExt((*factory2.ScopeFactoryInfo)(nil), false)

	//Action

	//Validate
	assert.PanicsWithError(t, "cannot find service int", func() {
		_ = p.GetService(&str)
	})
}

// Проверка работы метода разрешения зависимостей с временем жизни Transient
func TestProviderAndGetServiceAndTransientThenSuccess(t *testing.T) {
	//Test data
	str := ""
	strValue := reflect.ValueOf(&str)
	strType := strValue.Elem().Type()

	fStr := func() string { return str }

	sfiStr := (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(fStr), factory.Transient)

	sfmm := scopeFactoryMapMocked{}

	p := (*provider.Provider)(nil).New(&sfmm)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmm.Mock}).
		OnExt(sfmm.GetFactoryInfo).
		ReturnExt(sfiStr, true)

	(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmm.Mock}).
		OnExt(sfmm.GetFactoryInfoReflectType, strType).
		ReturnExt(sfiStr, true)

	//Action
	var res string
	err := p.GetService(&res)

	//Validate
	assert.NoError(t, err)
	assert.Equal(t, str, res)
}

// Проверка работы метода разрешения зависимостей с временем жизни Singleton
func TestProviderAndGetServiceAndLifeScopeThenTableResult(t *testing.T) {
	t.Parallel()

	//Test data
	grCount := gofakeit.Number(1, 1000)

	strVal := ""
	var str = &strVal
	strValue := reflect.ValueOf(&str)
	strType := strValue.Elem().Type()

	fStr := func() *string {
		newStr := ""

		return &newStr
	}

	testDatas := []struct {
		name        string
		factoryInfo factory2.IScopeFactoryInfo
		assertFunc  func(
			interface{},
			interface{},
		) bool
		isDifferentScope bool
	}{
		{"test Transient lifecycle with same scope", (*factory2.TransientFactoryInfo)(nil).New(reflect.ValueOf(fStr)), func(
			ptr1 interface{},
			ptr2 interface{},
		) bool {
			return ptr1 != ptr2
		}, false},
		{"test Transient lifecycle with different scope", (*factory2.TransientFactoryInfo)(nil).New(reflect.ValueOf(fStr)), func(
			ptr1 interface{},
			ptr2 interface{},
		) bool {
			return ptr1 != ptr2
		}, true},
		{"test Scoped lifecycle with same scope", (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(fStr), factory.Scoped), func(
			ptr1 interface{},
			ptr2 interface{},
		) bool {
			return ptr1 == ptr2
		}, false},
		{"test Scoped lifecycle with different scope", (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(fStr), factory.Scoped), func(
			ptr1 interface{},
			ptr2 interface{},
		) bool {
			return ptr1 != ptr2
		}, true},
		{"test Singleton lifecycle with same scope", (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(fStr), factory.Singleton), func(
			ptr1 interface{},
			ptr2 interface{},
		) bool {
			return ptr1 == ptr2
		}, false},
		{"test Singleton lifecycle with different scope", (*factory2.ScopeFactoryInfo)(nil).New(reflect.ValueOf(fStr), factory.Singleton), func(
			ptr1 interface{},
			ptr2 interface{},
		) bool {
			return ptr1 == ptr2
		}, true},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			sfmmGlobal := scopeFactoryMapMocked{}
			sfmmScoped := scopeFactoryMapMocked{}

			pGlobal := (*provider.Provider)(nil).New(&sfmmGlobal)

			(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmmGlobal.Mock}).
				OnExt(sfmmGlobal.GetFactoryInfo).
				ReturnExt(testData.factoryInfo, true)

			(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmmGlobal.Mock}).
				OnExt(sfmmGlobal.GetFactoryInfoReflectType, strType).
				ReturnExt(testData.factoryInfo, true)

			(&mock2.MockR1[factory2.IScopeFactoryMap]{Mock: &sfmmGlobal.Mock}).
				OnExt(sfmmGlobal.Copy).
				ReturnExt(factory2.IScopeFactoryMap(&sfmmScoped))

			(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmmScoped.Mock}).
				OnExt(sfmmScoped.GetFactoryInfo).
				ReturnExt(testData.factoryInfo.Copy(), true)

			(&mock2.MockR2[factory2.IScopeFactoryInfo, bool]{Mock: &sfmmScoped.Mock}).
				OnExt(sfmmScoped.GetFactoryInfoReflectType, strType).
				ReturnExt(testData.factoryInfo.Copy(), true)

			//Action
			var res0 *string
			_ = pGlobal.GetService(&res0)

			wg := sync.WaitGroup{}
			wg.Add(grCount)
			for i := 0; i < grCount; i++ {
				go func() {
					defer wg.Done()

					var res1 *string

					if testData.isDifferentScope {
						_ = pGlobal.CreateScopedServiceProvider().GetService(&res1)
					} else {
						_ = pGlobal.GetService(&res1)
					}

					//Validate
					assert.True(t, testData.assertFunc(res0, res1))
				}()
			}

			wg.Wait()
		})
	}
}
