package interfaces

import (
	"reflect"
)

type FactoryMap struct {
	factoryInfo map[reflect.Type]*FactoryInfo
}

func (*FactoryMap) New() *FactoryMap {
	return &FactoryMap{
		factoryInfo: make(map[reflect.Type]*FactoryInfo),
	}
}

func (fm *FactoryMap) GetFactoryInfos() (res []*FactoryInfo) {
	for _, v := range fm.factoryInfo {
		res = append(res, v)
	}

	return res
}

func (fm *FactoryMap) SetFactoryInfo(
	objectType reflect.Type,
	factoryInfo *FactoryInfo,
) {
	fm.factoryInfo[objectType] = factoryInfo
}
