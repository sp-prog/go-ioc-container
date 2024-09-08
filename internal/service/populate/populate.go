package populate

import (
	"errors"
	"github.com/sp-prog/go-ioc-container/internal/service/provider"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/populate"
	iprovider "github.com/sp-prog/go-ioc-container/pkg/interfaces/service/provider"
	"reflect"
)

type Populate struct {
	provider iprovider.IProvider
}

func (*Populate) New(provider iprovider.IProvider) populate.IPopulate {
	return &Populate{
		provider: provider,
	}
}

func (p *Populate) Populate(target interface{}) error {
	return p.PopulateWithExtras(
		target,
		(*provider.ServiceMap)(nil).New(),
	)
}

// PopulateWithExtras Заполняить поля структуры указанными значениями
// поля, для которых значенияне указаны, заполнить зависимостями
func (p *Populate) PopulateWithExtras(
	target interface{},
	extras provider.ServiceMap,
) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().Kind() == reflect.Struct {
		targetValue = targetValue.Elem()
		for i := 0; i < targetValue.Type().NumField(); i++ {
			fieldVal := targetValue.Field(i)
			if fieldVal.CanSet() {
				if extra, ok := extras[fieldVal.Type()]; ok {
					fieldVal.Set(extra)
				} else {
					//err = p.provider.resolveServiceFromValue(fieldVal.Addr())
					err = p.provider.GetService(fieldVal.Addr())
				}
			}

		}
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
