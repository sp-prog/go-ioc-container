package factory

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"reflect"
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
			factoryInfo.Lifecycle(),
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

func (fm *ScopeFactoryMap) GetFactoryInfo(
	target interface{},
) (IScopeFactoryInfo, bool) {
	targetValue := reflect.ValueOf(target)
	serviceType := targetValue.Elem().Type()

	return fm.GetFactoryInfoReflectType(serviceType)
}

func (fm *ScopeFactoryMap) GetFactoryInfoReflectType(
	reflectType reflect.Type,
) (IScopeFactoryInfo, bool) {
	binding, found := fm.factoryInfo[reflectType]

	return binding, found
}
