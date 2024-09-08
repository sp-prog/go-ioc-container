package mock

import (
	"github.com/sp-prog/go-ioc-container/test/extensions/reflect"
	"github.com/sp-prog/go-ioc-container/test/extensions/testify/call"
	"github.com/stretchr/testify/mock"
)

type MockR0 struct {
	*mock.Mock
}

func (m MockR0) OnExt(
	funcInfo interface{},
) *call.CallR0 {
	funcName := reflect.GetSimpleFuncName(funcInfo)

	return &call.CallR0{Call: m.On(funcName)}
}
