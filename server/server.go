package server
//服务端，包括创建服务器，路由表的设置，加载静态资源等
import (
	"Bubble-simple-web/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type TodoServer struct {
	database store.Store
	srv *gin.Engine
}

func NewTodoServer(db store.Store) TodoServer {
	todosrv:=TodoServer{
		database: db,
		srv: gin.Default(),
	}

	todosrv.loadStaticResource()
	todosrv.initGruopRouters()
	return todosrv
}
func (t *TodoServer)initGruopRouters()  {
	t.srv.GET("/index",t.getIndexHandler)
	v1Group:=t.srv.Group("/v1")
	{
		v1Group.GET("/todo",t.getAllTodoHandler)
		v1Group.POST("/todo",t.createTodoHandler)
		v1Group.PUT("/todo/:id",t.updateTodoHandler)
		v1Group.DELETE("/todo/:id",t.deleteTodoHandler)
	}
	t.srv.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path="/index"
		t.srv.HandleContext(c)

	})
}
func (t*TodoServer)ListenAndServe(addr string) (<-chan error,error)  {
	var err error
	errChan:=make(chan error)
	go func() {
		err=t.srv.Run(addr)
		errChan<-err
	}()
	select {
	case err=<-errChan:
		return nil,err
	case <-time.After(time.Second):
		return errChan,nil
	}
}

func (t*TodoServer)ShutDown()  {

}

func (t *TodoServer) loadStaticResource() {
	t.srv.Static("/static", "./static")
	t.srv.LoadHTMLGlob("templates/*")
}

func (t* TodoServer)getIndexHandler(c*gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}
func (t* TodoServer)getAllTodoHandler(c*gin.Context)  {
	todos,err:=t.database.GetAllTodos()
	if err!=nil{
		c.IndentedJSON(http.StatusOK,gin.H{"error":"cannot get all todos"})
		return
	}
	c.IndentedJSON(http.StatusOK,todos)
}
func (t* TodoServer)createTodoHandler(c*gin.Context)  {
	var todo store.Todo
	c.BindJSON(&todo)
	err:=t.database.CreateATodo(&todo)
	if err!=nil {
		c.IndentedJSON(http.StatusOK,gin.H{"error":"cannot create todo"})
		return
	}
	c.IndentedJSON(http.StatusOK,todo)
}
func (t*TodoServer)updateTodoHandler(c*gin.Context)  {
	var todo store.Todo
	todo.ID,_=strconv.Atoi(c.Param("id"))
	c.BindJSON(&todo)
	fmt.Printf("%v\n",todo)
	if err:=t.database.UpdateATodo(&todo);err!=nil{
		c.IndentedJSON(http.StatusOK,gin.H{
			"error":"cannot update todo",
		})
		return
	}
	c.IndentedJSON(http.StatusOK,todo)
}
func (t*TodoServer)deleteTodoHandler(c*gin.Context)  {
	var todo store.Todo
	todo.ID,_ =strconv.Atoi(c.Param("id"))
	fmt.Printf("%v\n",todo)
	if err:=t.database.DeleteATodo(&todo);err!=nil{
		c.IndentedJSON(http.StatusOK,gin.H{
			"error":"cannot delete todo",
		})
		return
	}
	c.IndentedJSON(http.StatusOK,nil)
	//c.Request.URL.Path="/index"
	//t.srv.HandleContext(c)
}