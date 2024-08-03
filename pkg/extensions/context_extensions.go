// Package extensions Пакет назван "расширения" по аналогии с Net'овскими файлами расширений
// Т.е. в этом файле "расширяется" функционал структуры context.Context
package extensions

import (
	"context"
	"github.com/sp-prog/go-ioc-container/internal/extensions"
)

// NewServiceContext Проксирование для скрытия реализации
// Создать контекст со списком зависимостей
func NewServiceContext(c context.Context) context.Context {
	return extensions.NewServiceContext(c)
}
