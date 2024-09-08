package testify

import (
	"github.com/sp-prog/go-ioc-container/test/extensions/reflect"
	"github.com/stretchr/testify/mock"
)

type Mock[TReturn1 interface{}, TReturn2 interface{}] struct {
	*mock.Mock
}

func (m Mock[TReturn1, TReturn2]) OnExt(
	funcInfo interface{},
) *Call[TReturn1, TReturn2] {
	funcName := reflect.GetSimpleFuncName(funcInfo)

	return &Call[TReturn1, TReturn2]{Call: m.On(funcName)}
}
