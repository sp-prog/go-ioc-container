package provider

import (
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestServiceMapAndNewThenCreated(t *testing.T) {
	//Test data

	//Action
	res := (*provider.ServiceMap)(nil).New()

	//Validate
	assert.NotNil(t, res)
}
