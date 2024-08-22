package factory

import (
	"reflect"
	"testing"

	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
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
		reflect.TypeOf(""),
	)

	//Validate
	assert.NotNil(t, res)
}
