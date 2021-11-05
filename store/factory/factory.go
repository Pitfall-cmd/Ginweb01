package factory
//工厂模式里的工厂，提供注册数据库的相关信息
import (
	"Bubble-simple-web/store"
	"fmt"
	"sync"
)

var (
	providerMu sync.RWMutex
	providers=make(map[string]store.Store)
)

func Register(name string ,p store.Store)  {
	providerMu.Lock()
	defer  providerMu.Unlock()
	if p==nil {
		panic("store:register provider is nil")
	}
	if _,exit:=providers[name];exit{
		panic("store:"+name+" has been registered before")
	}
	providers[name]=p
}

func New(name string) (store.Store,error) {
	providerMu.Lock()
	defer providerMu.Unlock()
	p,ok:=providers[name]
	if !ok{
		return nil,fmt.Errorf("store:Unknown provider %s",name)
	}
	return p,nil
}