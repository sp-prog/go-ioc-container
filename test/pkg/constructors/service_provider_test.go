package constructors

import (
	"github.com/sp-prog/go-ioc-container/pkg/constructors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работы метода расширения
func TestServiceProviderNewServiceProvider(t *testing.T) {
	res := constructors.NewServiceProvider()

	assert.NotNil(t, res)
}
