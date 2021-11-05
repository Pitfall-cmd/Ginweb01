package main

/*
这是一个简单版本，所有功能都在一个页面上，所有逻辑都在这里
 */

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

var (
	DB *gorm.DB
)

func initMysql() error {
	var err error
	DB,err=gorm.Open("mysql","root:109456@(127.0.0.1:3306)/bubblesql?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil {
		return err
	}
	return DB.DB().Ping()
}

func main() {
	//数据库连接
	err:=initMysql()
	if err!=nil{
		panic(err)
	}
	DB.AutoMigrate(&Todo{})

	r:=gin.Default()
	// // 配置静态文件夹路径 第一个参数是api(index.html中请求静态资源时/static打头的)，第二个是文件夹路径
	r.Static("/static","./static")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	v1Gruop:=r.Group("/v1")
	{	//查看所有
		v1Gruop.GET("/todo", func(c *gin.Context) {
			//查询所有数据
			var todolist []Todo
			DB.Find(&todolist)
			c.IndentedJSON(http.StatusOK,todolist)
		})
		//添加
		v1Gruop.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			err:=DB.Create(&todo).Error
			if err!=nil{
				c.IndentedJSON(http.StatusOK,gin.H{
					"err":err.Error(),
				})
			}else{
				c.IndentedJSON(http.StatusOK,todo)
			}
		})
		//查看某一个
		v1Gruop.GET("/todo/:id", func(c *gin.Context) {

		})

		//修改某一个待办事项
		v1Gruop.PUT("/todo/:id", func(c *gin.Context) {
			id:=c.Param("id")
			var todo Todo
			if err:=DB.Where("id=?",id).Find(&todo).Error;err!=nil{
				c.IndentedJSON(http.StatusOK,gin.H{"error":err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err:=DB.Save(&todo).Error;err!=nil{
				c.IndentedJSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.IndentedJSON(http.StatusOK,todo)
			}
		})
		//删除
		v1Gruop.DELETE("/todo/:id", func(c *gin.Context) {
			id:=c.Param("id")
			if err:=DB.Where("id=?",id).Delete(&Todo{}).Error;err!=nil{
				c.IndentedJSON(http.StatusOK,gin.H{"error":err})
			}else{
				c.IndentedJSON(http.StatusOK,nil)
			}

		})
	}
	r.Run()
}
