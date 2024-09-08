package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
)

type ScopeFactoryMap struct {
	factoryInfo map[reflect.Type]IScopeFactoryInfo
}

func (*ScopeFactoryMap) New() IScopeFactoryMap {
	return &ScopeFactoryMap{
		factoryInfo: make(map[reflect.Type]IScopeFactoryInfo),
	}
}

func (sfm *ScopeFactoryMap) SetFactoryInfo(
	factoryInfo *collection.FactoryInfo,
) {
	if factoryInfo.Lifecycle() == factory.Transient {
		sfm.factoryInfo[factoryInfo.ObjectType()] = (*TransientFactoryInfo)(nil).New(
			factoryInfo.FactoryFunc(),
		)

		return
	}

	sfm.factoryInfo[factoryInfo.ObjectType()] = (*ScopeFactoryInfo)(nil).New(
		factoryInfo.FactoryFunc(),
		factoryInfo.Lifecycle(),
	)
}

func (sfm *ScopeFactoryMap) GetFactoryInfo(
	target interface{},
) (IScopeFactoryInfo, bool) {
	targetValue := reflect.ValueOf(target)
	serviceType := targetValue.Elem().Type()

	return sfm.GetFactoryInfoReflectType(serviceType)
}

func (sfm *ScopeFactoryMap) GetFactoryInfoReflectType(
	reflectType reflect.Type,
) (IScopeFactoryInfo, bool) {
	binding, found := sfm.factoryInfo[reflectType]

	return binding, found
}
