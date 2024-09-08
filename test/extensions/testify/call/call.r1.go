package call

import "github.com/stretchr/testify/mock"

type CallR1[TReturn1 interface{}] struct {
	*mock.Call
}

func (c CallR1[TReturn1]) ReturnExt(
	returnArguments1 TReturn1,
) *mock.Call {
	return c.Return(returnArguments1)
}
