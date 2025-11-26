package main

import (
	"native-web-demo/internal/database" // 导入数据库连接包
	"native-web-demo/internal/handlers"
	"native-web-demo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 连接数据库
	db := database.ConnectDatabase()

	// 2. 创建一个 Gin 路由器
	router := gin.Default() // gin.Default() 已经自带了 Gin 自己的 Logger 和 Recovery 中间件。

	// 3. 注册我们的自定义中间件
	router.Use(middleware.Logging())

	// 4. 创建一个 Handler 实例，并将数据库连接注入进去。
	// 这样，handlers 包中的所有方法都可以访问到数据库连接。
	h := &handlers.Handler{DB: db}

	// 5. 注册路由和处理器
	// 对于需要数据库连接的处理器，我们使用 h.MethodName
	router.GET("/", handlers.Home) // Home 函数不需要数据库连接，保持独立
	router.GET("/users", h.GetUsers)
	router.POST("/users", h.CreateUser)

	// 6. 启动服务器
	router.Run(":8080")
}