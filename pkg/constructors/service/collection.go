package service

import (
	collection2 "github.com/sp-prog/go-ioc-container/internal/service/collection"
	"github.com/sp-prog/go-ioc-container/pkg/interfaces/service/collection"
)

// NewCollection Проксирование для скрытия реализации
// Создать новую коллекцию зависимостей, по которой в последствии необходимо
// собрать поставщика зависимостей
func NewCollection() collection.ICollection {
	return (*collection2.Collection)(nil).New()
}
