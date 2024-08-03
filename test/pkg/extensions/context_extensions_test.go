package extensions

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/extensions"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ContextMocked struct {
	mock.Mock
	context.Context
}

func (*ContextMocked) Value(key any) any {
	return type_factory.ServiceMap{}
}

// Проверка работы метода расширения
func TestNewServiceContext(t *testing.T) {
	contextMocked := new(ContextMocked)

	contextMocked.On("Value")

	res := extensions.NewServiceContext(contextMocked)

	assert.NotNil(t, res)
}
