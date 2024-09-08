package mock

import (
	"github.com/sp-prog/go-ioc-container/test/extensions/reflect"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify/call"
	"github.com/stretchr/testify/mock"
)

type MockR2[TReturn1 interface{}, TReturn2 interface{}] struct {
	*mock.Mock
}

func (m MockR2[TReturn1, TReturn2]) OnExt(
	funcInfo interface{},
) *call.CallR2[TReturn1, TReturn2] {
	funcName := reflect.GetSimpleFuncName(funcInfo)

	return &call.CallR2[TReturn1, TReturn2]{Call: m.On(funcName)}
}
