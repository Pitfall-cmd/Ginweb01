package DatabaseStore

import (
	"Bubble-simple-web/store"
	"sync"
)
//因为使用了MYSQL作为存储 所以这一部分就不实现了
type Memstore struct {
	sync.RWMutex
	todos map[int]*store.Todo
}

func (m *Memstore) CreateATodo(todo *store.Todo) error {
	panic("implement me")
}

func (m *Memstore) UpdateATodo(todo *store.Todo) error {
	panic("implement me")
}

func (m *Memstore) GetATodo(i int) (store.Todo, error) {
	panic("implement me")
}

func (m *Memstore) GetAllTodos() ([]store.Todo, error) {
	panic("implement me")
}

func (m *Memstore) DeleteATodo(todo *store.Todo) error {
	panic("implement me")
}
