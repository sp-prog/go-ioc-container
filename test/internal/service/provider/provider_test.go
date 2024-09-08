package provider

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	icollection "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify"
	mock2 "github.com/sp-prog/go-ioc-container/test/extensions/testify/mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestProviderAndNewThenNewObject(t *testing.T) {
	//Test data
	//Action
	res := (*provider.Provider)(nil).New(nil)

	//Validate
	assert.NotNil(t, res)
}

// Проверка работы метода создания новой области жизни объектов
func TestProviderAndCreateScopedServiceProviderThenChildProvider(t *testing.T) {
	//Test data
	sfmm := iScopeFactoryMapMocked{}

	sp := (*provider.Provider)(nil).New(&sfmm)

	//Action
	res := sp.CreateScopedServiceProvider()

	//Validate
	assert.NotNil(t, res)
}

// Проверка работы метода сборки провайдера из контейнера
func TestProviderAndBuildThenSuccess(t *testing.T) {
	//Test data

	//f := func() string { return "" }
	sfmm := iScopeFactoryMapMocked{}
	cm := iCollectionMocked{}
	fm := (*icollection.FactoryMap)(nil).New()
	fm.SetFactoryInfo(reflect.TypeOf(""), nil)

	mock2.MockR1[factory2.IScopeFactoryMap]{Mock: &sfmm.Mock}.
		OnExt(sfmm.New).
		ReturnExt(&sfmm)

	mock2.MockR0{Mock: &sfmm.Mock}.
		OnExt(sfmm.SetFactoryInfo).
		Return()

	mock2.MockR1[*icollection.FactoryMap]{Mock: &cm.Mock}.
		OnExt(cm.Services).
		ReturnExt(fm)

	sp := (*provider.Provider)(nil).New(&sfmm)

	//Action
	sp.Build(&cm)

	//Validate
	testify.Assert{Mock: &cm.Mock}.
		AssertNumberOfCallsEx(t, cm.Services, 1)
	testify.Assert{Mock: &sfmm.Mock}.
		AssertNumberOfCallsEx(t, sfmm.SetFactoryInfo, 1)
}

//// Проверка работы метода вызова функции с незарегистрированной зависимостью
//func TestProviderCallAndArgsAndNoDependencyThenError(t *testing.T) {
//	//Test data
//	id := uuid.NewString()
//	f := func(i int) string { return id }
//
//	//Action
//	sp := (*provider.Provider)(nil).New()
//
//	//Validate
//	assert.PanicsWithError(t, "Cannot find service int", func() {
//		_, _ = sp.CallR2(f)
//	})
//}
//
//// Проверка работы метода вызова функции с зарегистрированной зависимостью
//func TestProviderCallAndArgsThenError(t *testing.T) {
//	t.Parallel()
//
//	//Test data
//	id := uuid.NewString()
//	arg := gofakeit.Number(1, 1000)
//
//	funcArg := func() int { return arg }
//	funcCall := func(i int) string { return id }
//
//	sc := (*collection.Collection)(nil).New()
//
//	testDatas := []struct {
//		name          string
//		lifecycleFunc func(interface{}) error
//	}{
//		{"test CallR2 with Transient lifecycle", sc.AddTransient},
//		{"test CallR2 with Scoped lifecycle", sc.AddScoped},
//		{"test CallR2 with Singleton lifecycle", sc.AddSingleton},
//	}
//
//	for _, testData := range testDatas {
//		t.Run(testData.name, func(t *testing.T) {
//			t.Parallel()
//
//			sp := (*provider.Provider)(nil).New()
//
//			_ = testData.lifecycleFunc(funcArg)
//
//			sp.Build(sc)
//
//			//Action
//			res, err := sp.CallR2(funcCall)
//
//			//Validate
//			assert.NoError(t, err)
//			assert.Equal(t, id, res[0])
//		})
//	}
//}
