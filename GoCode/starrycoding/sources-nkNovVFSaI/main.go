package main

import (
	"os"
)

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//r.Static("/public", "./public") // 静态文件服务
	//r.LoadHTMLGlob("views/**/*") // 载入html模板目录

	// web路由
	r.GET("/", Home)

	return r
}

func Home(c *gin.Context) {
	c.String(200, "Hello, world!")
}
func main() {
	r := InitRouter()
	port := os.Getenv("HTTP_PORT")
	go r.Run(":" + port) // 监听并在 0.0.0.0:8080 上启动服务
	select {}
}
