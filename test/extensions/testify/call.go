package testify

import "github.com/stretchr/testify/mock"

type Call[TReturn1 interface{}, TReturn2 interface{}] struct {
	*mock.Call
}

func (c Call[TReturn1, _]) ReturnExt1(
	returnArguments1 TReturn1,
) *mock.Call {
	return c.Return(returnArguments1)
}

func (c Call[TReturn1, TReturn2]) ReturnExt2(
	returnArguments1 TReturn1,
	returnArguments2 TReturn2,
) *mock.Call {
	return c.Return(returnArguments1, returnArguments2)
}
