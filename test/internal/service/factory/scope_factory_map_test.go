package factory

import (
	factory2 "github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestScopeFactoryMapNew(t *testing.T) {
	res := (*factory2.ScopeFactoryMap)(nil).New()

	assert.NotNil(t, res)
}

// Проверка добавления и получения зависимости с указанным временем жизни
// "Табличный тест" https://go.dev/wiki/TableDrivenTests
func TestScopeFactoryMapAndSetFactoryInfoAndGetFactoryInfoThenTableResult(t *testing.T) {
	t.Parallel()

	//Test data
	testDatas := []struct {
		name         string
		lifecycle    factory.Lifecycle
		res          interface{}
		byReflection bool
	}{
		{"test Transient lifecycle with get by type", factory.Transient, &factory2.TransientFactoryInfo{}, false},
		{"test Singleton lifecycle with get by type", factory.Singleton, &factory2.ScopeFactoryInfo{}, false},
		{"test Scoped lifecycle with get by type", factory.Scoped, &factory2.ScopeFactoryInfo{}, false},
		{"test Transient lifecycle with get by reflection", factory.Transient, &factory2.TransientFactoryInfo{}, true},
		{"test Singleton lifecycle with get by reflection", factory.Singleton, &factory2.ScopeFactoryInfo{}, true},
		{"test Scoped lifecycle with get by reflection", factory.Scoped, &factory2.ScopeFactoryInfo{}, true},
	}

	typeValue := ""
	factoryFunc := func() string { return typeValue }
	objectType := reflect.TypeOf(typeValue)

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			scopeFactoryMap := (*factory2.ScopeFactoryMap)(nil).New()
			factoryInfo := (*collection.FactoryInfo)(nil).New(
				reflect.ValueOf(factoryFunc),
				testData.lifecycle,
				objectType,
			)

			//Action
			scopeFactoryMap.SetFactoryInfo(
				factoryInfo,
			)

			var val factory2.IScopeFactoryInfo
			var exists bool
			if testData.byReflection {
				val, exists = scopeFactoryMap.GetFactoryInfoReflectType(objectType)
			} else {
				val, exists = scopeFactoryMap.GetFactoryInfo(&typeValue)
			}

			//Validate
			assert.Truef(t, exists, testData.name)
			assert.IsTypef(t, testData.res, val, testData.name)
		})
	}
}
