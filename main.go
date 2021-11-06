package main

/*
这里是原先的cmd/version2/main.go 因为go run main.go时 文件路径的问题，所以把main移到最外层
这个是经过修改之后的重新封装的main，在这里只有启动初始化等工作
*/
import (
	_ "Bubble-simple-web/internal/DatabaseStore"
	"Bubble-simple-web/server"
	"Bubble-simple-web/store/factory"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db,err:=factory.New("mysql")
	if err!=nil {
		panic(err)
	}
	srv:=server.NewTodoServer(db)
	errChan,err:=srv.ListenAndServe(":8080")
	if err!=nil{
		log.Println("Web server start failer :",err)
	}
	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGINT,syscall.SIGTERM)
	select { //监听来自errChan和c的事件
	case err=<-errChan:
		log.Println("web server run failed:",err)
	case <-c:
		log.Println("web program is exiting....")
		srv.ShutDown()
	}
	log.Println("bookstore program exit ok")
}



