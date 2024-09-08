package factory

import (
	"github.com/stretchr/testify/mock"
)

type contextMocked struct {
	mock.Mock
}

func (cm *contextMocked) fakeFunc() string {
	args := cm.Called()

	return args.Get(0).(string)
}
