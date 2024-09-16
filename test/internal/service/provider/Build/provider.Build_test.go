package Build

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	icollection "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify"
	mock2 "github.com/sp-prog/go-ioc-container/test/extensions/testify/mock"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type collectionMocked struct {
	mock.Mock
	icollection.ICollection
}

func (cm *collectionMocked) Services() *icollection.FactoryMap {
	args := cm.Called()

	return args.Get(0).(*icollection.FactoryMap)
}

type scopeFactoryMapMocked struct {
	mock.Mock
	factory2.IScopeFactoryMap
}

func (sfmm *scopeFactoryMapMocked) New() factory2.IScopeFactoryMap {
	args := sfmm.Called()

	return args.Get(0).(factory2.IScopeFactoryMap)
}

func (sfmm *scopeFactoryMapMocked) SetFactoryInfo(factoryInfo factory2.IScopeFactoryInfo) {
	sfmm.Called()
}

func (sfmm *scopeFactoryMapMocked) GetFactoryInfo(target interface{}) (factory2.IScopeFactoryInfo, bool) {
	args := sfmm.Called()

	return args.Get(0).(factory2.IScopeFactoryInfo), args.Bool(1)
}

// Проверка работы метода сборки провайдера из контейнера
func TestProviderAndBuildThenSuccess(t *testing.T) {
	//Test data
	sfmm := scopeFactoryMapMocked{}
	cm := collectionMocked{}
	fm := (*icollection.FactoryMap)(nil).New()
	fm.SetFactoryInfo(reflect.TypeOf(""), nil)

	(&mock2.MockR1[factory2.IScopeFactoryMap]{Mock: &sfmm.Mock}).
		OnExt(sfmm.New).
		ReturnExt(&sfmm)

	(&mock2.MockR0{Mock: &sfmm.Mock}).
		OnExt(sfmm.SetFactoryInfo).
		Return()

	(&mock2.MockR1[*icollection.FactoryMap]{Mock: &cm.Mock}).
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
