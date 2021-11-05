package store
//相当于存储模型的层次
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

type Store interface {
	CreateATodo(todo *Todo) error
	UpdateATodo(todo *Todo) error
	GetATodo(int)(Todo,error)
	GetAllTodos()([]Todo,error)
	DeleteATodo(todo *Todo)	error
}