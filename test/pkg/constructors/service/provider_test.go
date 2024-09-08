package service

import (
	"github.com/sp-prog/go-ioc-container/pkg/constructors/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работы метода расширения
func TestProviderNewProvider(t *testing.T) {
	res := service.NewProvider()

	assert.NotNil(t, res)
}
