package provider

import (
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка работы конструктора
func TestProviderAndNewThenNewObject(t *testing.T) {
	//Test data
	//Action
	res := (*provider.Provider)(nil).New(nil)

	//Validate
	assert.NotNil(t, res)
}
