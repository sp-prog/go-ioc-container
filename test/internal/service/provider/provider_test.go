package provider

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	icollection "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type icollectionMocked struct {
	mock.Mock
	icollection.ICollection
}

func (cm *icollectionMocked) Services() *icollection.FactoryMap {
	args := cm.Called()

	return args.Get(0).(*icollection.FactoryMap)
}

type iScopeFactoryMapMocked struct {
	mock.Mock
	factory2.IScopeFactoryMap
}

func (sfmm *iScopeFactoryMapMocked) New() factory2.IScopeFactoryMap {
	args := sfmm.Called()

	return args.Get(0).(factory2.IScopeFactoryMap)
}

func (sfmm *iScopeFactoryMapMocked) SetFactoryInfo(factoryInfo *icollection.FactoryInfo) {
	sfmm.Called()
}

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
	cm := icollectionMocked{}
	fm := (*icollection.FactoryMap)(nil).New()
	fm.SetFactoryInfo(reflect.TypeOf(""), nil)

	testify.Mock[factory2.IScopeFactoryMap, testify.FakeType]{Mock: &sfmm.Mock}.
		OnExt(sfmm.New).
		ReturnExt1(&sfmm)

	testify.Mock[factory2.IScopeFactoryMap, testify.FakeType]{Mock: &sfmm.Mock}.
		OnExt(sfmm.SetFactoryInfo).
		Return()

	testify.Mock[*icollection.FactoryMap, testify.FakeType]{Mock: &cm.Mock}.
		OnExt(cm.Services).
		ReturnExt1(fm)

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
//		_, _ = sp.Call(f)
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
//		{"test Call with Transient lifecycle", sc.AddTransient},
//		{"test Call with Scoped lifecycle", sc.AddScoped},
//		{"test Call with Singleton lifecycle", sc.AddSingleton},
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
//			res, err := sp.Call(funcCall)
//
//			//Validate
//			assert.NoError(t, err)
//			assert.Equal(t, id, res[0])
//		})
//	}
//}
