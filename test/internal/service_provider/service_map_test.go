package service_provider

import (
	"testing"

	"github.com/sp-prog/go-ioc-container/internal/service_provider"
	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestServiceMapAndNewThenCreated(t *testing.T) {
	//Test data

	//Action
	res := (*service_provider.ServiceMap)(nil).New()

	//Validate
	assert.NotNil(t, res)
}
