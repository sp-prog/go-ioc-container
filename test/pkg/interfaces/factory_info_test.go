package interfaces

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Проверка работы конструктора
func TestFactoryInfoNew(t *testing.T) {
	res := (*interfaces.FactoryInfo)(nil).New(
		reflect.Value{},
		interfaces.Singleton,
		nil,
	)

	assert.NotNil(t, res)
}

// Проверка получения типа конструктора
func TestFactoryFunc(t *testing.T) {

	//Test data
	factoryFunc := reflect.Value{}

	//Action
	res := (*interfaces.FactoryInfo)(nil).New(
		factoryFunc,
		interfaces.Singleton,
		nil,
	)

	//Validate
	assert.Equal(t, factoryFunc, res.FactoryFunc())
}

// Проверка получения области жизни
func TestLifecycle(t *testing.T) {

	//Test data
	lifetime := interfaces.Singleton

	//Action
	res := (*interfaces.FactoryInfo)(nil).New(
		reflect.Value{},
		lifetime,
		nil,
	)

	//Validate
	assert.Equal(t, lifetime, res.Lifecycle())
}

// Проверка получения области жизни
func TestObjectType(t *testing.T) {

	//Test data
	var objectType reflect.Type = nil

	//Action
	res := (*interfaces.FactoryInfo)(nil).New(
		reflect.Value{},
		interfaces.Singleton,
		objectType,
	)

	//Validate
	assert.Equal(t, objectType, res.ObjectType())
}
