package CreateScopedServiceProvider

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

type scopeFactoryMapMocked struct {
	mock.Mock
	factory2.IScopeFactoryMap
}

func (sfmm *scopeFactoryMapMocked) New() factory2.IScopeFactoryMap {
	args := sfmm.Called()

	return args.Get(0).(factory2.IScopeFactoryMap)
}

func (sfmm *scopeFactoryMapMocked) Copy() factory2.IScopeFactoryMap {
	return nil
}

// Проверка работы метода создания новой области жизни объектов
func TestProviderAndCreateScopedServiceProviderThenChildProvider(t *testing.T) {
	//Test data
	sfmm := scopeFactoryMapMocked{}

	sp := (*provider.Provider)(nil).New(&sfmm)

	//Action
	res := sp.CreateScopedServiceProvider()

	//Validate
	assert.NotNil(t, res)
}
