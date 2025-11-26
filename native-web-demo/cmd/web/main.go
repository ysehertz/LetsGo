package main

import (
	"native-web-demo/internal/handlers"
	"native-web-demo/internal/middleware" // 导入我们的自定义中间件包

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建一个 Gin 路由器
	router := gin.Default() // gin.Default() 已经自带了 Gin 自己的 Logger 和 Recovery 中间件。

	// 2. 注册我们的自定义中间件
	// router.Use() 方法用于注册全局中间件，它会应用到所有请求。
	// 这里我们使用自己实现的 Logging 中间件。
	router.Use(middleware.Logging())

	// 3. 注册路由和处理器
	router.GET("/", handlers.Home)
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)

	// 4. 启动服务器
	router.Run(":8080")
}
