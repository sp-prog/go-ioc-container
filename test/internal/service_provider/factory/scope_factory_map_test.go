package factory

import (
	"github.com/sp-prog/go-ioc-container/internal/service_provider/factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работы конструктора
func TestScopeFactoryMapNew(t *testing.T) {
	res := (*factory.ScopeFactoryMap)(nil).New()

	assert.NotNil(t, res)
}
