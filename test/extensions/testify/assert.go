package testify

import (
	"github.com/sp-prog/go-ioc-container/test/extensions/reflect"
	"github.com/stretchr/testify/mock"
)

type Assert struct {
	*mock.Mock
}

func (a Assert) AssertNumberOfCallsEx(
	t mock.TestingT,
	funcInfo interface{},
	expectedCalls int,
) bool {
	funcName := reflect.GetSimpleFuncName(funcInfo)

	return a.AssertNumberOfCalls(t, funcName, expectedCalls)
}
