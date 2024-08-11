package constructors

import (
	"github.com/sp-prog/go-ioc-container/pkg/constructors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работы метода расширения
func TestServiceCollectionNewServiceCollection(t *testing.T) {
	res := constructors.NewServiceCollection()

	assert.NotNil(t, res)
}
