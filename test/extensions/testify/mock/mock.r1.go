package mock

import (
	"github.com/sp-prog/go-ioc-container/test/extensions/reflect"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify/call"
	"github.com/stretchr/testify/mock"
)

type MockR1[TReturn1 interface{}] struct {
	*mock.Mock
}

func (m *MockR1[TReturn1]) OnExt(
	funcInfo interface{},
	arguments ...interface{},
) *call.CallR1[TReturn1] {
	funcName := reflect.GetSimpleFuncName(funcInfo)

	if len(arguments) > 0 {
		return &call.CallR1[TReturn1]{Call: m.On(funcName, arguments)}
	}

	return &call.CallR1[TReturn1]{Call: m.On(funcName)}
}
