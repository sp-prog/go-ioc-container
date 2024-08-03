package service_collection

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
)

func (sp *ServiceProvider) Populate(target interface{}) error {
	return sp.PopulateForContext(context.Background(), target)
}

func (sp *ServiceProvider) PopulateForContext(
	c context.Context,
	target interface{},
) (err error) {
	return sp.PopulateForContextWithExtras(
		c,
		target,
		type_factory.NewServiceMap(),
	)
}
