package mock

import (
	"github.com/sp-prog/go-ioc-container/test/extensions/reflect"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify/call"
	"github.com/stretchr/testify/mock"
)

type MockR1[TReturn1 interface{}] struct {
	*mock.Mock
}

func (m MockR1[TReturn1]) OnExt(
	funcInfo interface{},
) *call.CallR1[TReturn1] {
	funcName := reflect.GetSimpleFuncName(funcInfo)

	return &call.CallR1[TReturn1]{Call: m.On(funcName)}
}
