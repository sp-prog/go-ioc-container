package factory

import (
	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Проверка работы конструктора
func TestScopeFactoryMapNew(t *testing.T) {
	res := (*factory.ScopeFactoryMap)(nil).New()

	assert.NotNil(t, res)
}

// Проверка добавления зависимости со временм жизни Transient
func TestSetFactoryInfoTransient(t *testing.T) {
	//Test data
	objectType := reflect.TypeOf("")
	scopeFactoryMap := (*factory.ScopeFactoryMap)(nil).New()
	factoryInfo := (*interfaces.FactoryInfo)(nil).New(
		reflect.Value{},
		interfaces.Transient,
		objectType,
	)

	//Action
	scopeFactoryMap.SetFactoryInfo(
		factoryInfo,
	)

	//Validate
	val, exists := scopeFactoryMap.GetFactoryInfo(objectType)

	assert.True(t, exists)
	assert.IsType(t, &factory.TransientFactoryInfo{}, val)
}
