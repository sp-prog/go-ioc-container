package factory

import (
	"reflect"
	"testing"

	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestScopeFactoryMapNew(t *testing.T) {
	res := (*factory.ScopeFactoryMap)(nil).New()

	assert.NotNil(t, res)
}

// Проверка добавления и получения зависимости с указанным временем жизни
// "Табличный тест" https://go.dev/wiki/TableDrivenTests
func TestScopeFactoryMapAndSetFactoryInfoAndGetFactoryInfoThenTableResult(t *testing.T) {
	t.Parallel()

	//Test data
	testDatas := []struct {
		name         string
		lifecycle    interfaces.Lifecycle
		res          interface{}
		byReflection bool
	}{
		{"test Transient lifecycle with get by type", interfaces.Transient, &factory.TransientFactoryInfo{}, false},
		{"test Singleton lifecycle with get by type", interfaces.Singleton, &factory.ScopeFactoryInfo{}, false},
		{"test Scoped lifecycle with get by type", interfaces.Scoped, &factory.ScopeFactoryInfo{}, false},
		{"test Transient lifecycle with get by reflection", interfaces.Transient, &factory.TransientFactoryInfo{}, true},
		{"test Singleton lifecycle with get by reflection", interfaces.Singleton, &factory.ScopeFactoryInfo{}, true},
		{"test Scoped lifecycle with get by reflection", interfaces.Scoped, &factory.ScopeFactoryInfo{}, true},
	}

	typeValue := ""
	factoryFunc := func() string { return typeValue }
	objectType := reflect.TypeOf(typeValue)

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			scopeFactoryMap := (*factory.ScopeFactoryMap)(nil).New()
			factoryInfo := (*interfaces.FactoryInfo)(nil).New(
				reflect.ValueOf(factoryFunc),
				testData.lifecycle,
				objectType,
			)

			//Action
			scopeFactoryMap.SetFactoryInfo(
				factoryInfo,
			)

			var val factory.IScopeFactoryInfo
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
