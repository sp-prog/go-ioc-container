package service

import (
	"github.com/sp-prog/go-ioc-container/pkg/constructors/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работы метода расширения
func TestCollectionNewCollection(t *testing.T) {
	res := service.NewCollection()

	assert.NotNil(t, res)
}
