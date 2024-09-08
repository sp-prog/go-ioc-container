package factory

import (
	"github.com/sp-prog/go-ioc-container/internal/service/factory"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestTransientFactoryInfoAndNewThenCreated(t *testing.T) {
	//Test data

	//Action
	res := (*factory.TransientFactoryInfo)(nil).New(
		reflect.ValueOf(func() string {
			return ""
		}),
	)

	//Validate
	assert.NotNil(t, res)
}

// Проверка сеттеров и геттеров
func TestTransientFactoryInfoAndObjectType(t *testing.T) {
	//Test data

	//Action
	res := (*factory.TransientFactoryInfo)(nil).New(
		reflect.ValueOf(func() string {
			return ""
		}),
	)

	//Validate
	assert.NotNil(t, res)
}
