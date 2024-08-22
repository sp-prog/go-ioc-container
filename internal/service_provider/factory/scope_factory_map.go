package factory

import (
	"reflect"

	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
)

type ScopeFactoryMap struct {
	factoryInfo map[reflect.Type]IScopeFactoryInfo
}

func (*ScopeFactoryMap) New() *ScopeFactoryMap {
	return &ScopeFactoryMap{
		factoryInfo: make(map[reflect.Type]IScopeFactoryInfo),
	}
}

func (sfm *ScopeFactoryMap) SetFactoryInfo(
	factoryInfo *interfaces.FactoryInfo,
) {
	if factoryInfo.Lifecycle() == interfaces.Transient {
		sfm.factoryInfo[factoryInfo.ObjectType()] = (*TransientFactoryInfo)(nil).New(
			factoryInfo.FactoryFunc(),
			factoryInfo.ObjectType(),
		)

		return
	}

	sfm.factoryInfo[factoryInfo.ObjectType()] = (*ScopeFactoryInfo)(nil).New(
		factoryInfo.FactoryFunc(),
		factoryInfo.Lifecycle(),
		factoryInfo.ObjectType(),
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
