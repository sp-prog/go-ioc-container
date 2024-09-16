package call

import "github.com/stretchr/testify/mock"

type CallR2[TReturn1 interface{}, TReturn2 interface{}] struct {
	*mock.Call
}

func (c *CallR2[TReturn1, TReturn2]) ReturnExt(
	returnArguments1 TReturn1,
	returnArguments2 TReturn2,
) *CallR2[TReturn1, TReturn2] {
	c.Return(returnArguments1, returnArguments2)

	return c
}
