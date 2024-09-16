package populate

import (
	"github.com/sp-prog/go-ioc-container/internal/service/factory"
	"github.com/sp-prog/go-ioc-container/internal/service/populate"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работы конструктора
func TestPopulateAndNewThenCreated(t *testing.T) {
	//Test data
	sfm := (*factory.ScopeFactoryMap)(nil).New()

	p := (*provider.Provider)(nil).New(sfm)

	//Action
	res := (*populate.Populate)(nil).New(
		p,
	)

	//Validate
	assert.NotNil(t, res)
}
