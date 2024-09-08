package interfaces

import (
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/factory"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestFactoryInfoNew(t *testing.T) {
	res := (*collection.FactoryInfo)(nil).New(
		reflect.Value{},
		factory.Singleton,
		nil,
	)

	assert.NotNil(t, res)
}

// Проверка получения типа конструктора
func TestFactoryFunc(t *testing.T) {

	//Test data
	factoryFunc := reflect.Value{}

	//Action
	res := (*collection.FactoryInfo)(nil).New(
		factoryFunc,
		factory.Singleton,
		nil,
	)

	//Validate
	assert.Equal(t, factoryFunc, res.FactoryFunc())
}

// Проверка получения области жизни
func TestLifecycle(t *testing.T) {

	//Test data
	lifecycle := factory.Singleton

	//Action
	res := (*collection.FactoryInfo)(nil).New(
		reflect.Value{},
		lifecycle,
		nil,
	)

	//Validate
	assert.Equal(t, lifecycle, res.Lifecycle())
}

// Проверка получения области жизни
func TestObjectType(t *testing.T) {

	//Test data
	var objectType reflect.Type = nil

	//Action
	res := (*collection.FactoryInfo)(nil).New(
		reflect.Value{},
		factory.Singleton,
		objectType,
	)

	//Validate
	assert.Equal(t, objectType, res.ObjectType())
}
