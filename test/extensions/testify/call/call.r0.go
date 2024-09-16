package call

import "github.com/stretchr/testify/mock"

type CallR0 struct {
	*mock.Call
}

func (c *CallR0) ReturnExt() *mock.Call {
	return c.Return()
}
