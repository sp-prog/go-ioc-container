package provider

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	icollection "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/stretchr/testify/mock"
)

type iCollectionMocked struct {
	mock.Mock
	icollection.ICollection
}

func (cm *iCollectionMocked) Services() *icollection.FactoryMap {
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
