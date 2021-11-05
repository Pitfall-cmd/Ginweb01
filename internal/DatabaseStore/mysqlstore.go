package DatabaseStore
//这一部分实现了mysql的数据存储和操作
import (
	"Bubble-simple-web/store"
	"Bubble-simple-web/store/factory"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	db *gorm.DB
	err error

)

func init() {
	db,err=gorm.Open("mysql","root:109456@(127.0.0.1:3306)/bubblesql?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	db.AutoMigrate(&store.Todo{})

	factory.Register("mysql",&Mysql{})
}
type Mysql struct {
	sync.RWMutex
}

func (m *Mysql) CreateATodo(todo *store.Todo) error {
	m.Lock()
	defer m.Unlock()
	err=db.Create(todo).Error
	if err!=nil{
		return err
	}
	return nil
}

func (m *Mysql) UpdateATodo(todo *store.Todo) error {
	m.Lock()
	defer m.Unlock()
	err=db.Save(todo).Error
	if err!=nil{
		return err
	}
	return nil
}

func (m *Mysql) GetATodo(i int) (store.Todo, error) {
	panic("implement me")
}

func (m *Mysql) GetAllTodos() ([]store.Todo, error) {
	m.RLock()
	defer m.RUnlock()
	var todos []store.Todo
	if err=db.Find(&todos).Error;err!=nil{
		return nil,err
	}
	return todos,nil
}

func (m *Mysql) DeleteATodo(todo *store.Todo) error {
	m.Lock()
	defer m.Unlock()

	if err=db.Where("id=?",todo.ID).Delete(todo).Error;err!=nil{
		return err
	}
	return nil
}

