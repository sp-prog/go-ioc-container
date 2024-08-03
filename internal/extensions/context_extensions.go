package extensions

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/type_factory"
)

const ServiceKey = "services"

func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, type_factory.NewServiceMap())
	} else {
		return c
	}
}
