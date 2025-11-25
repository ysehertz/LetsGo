package main

import (
	"log"
	"net/http"

	"native-web-demo/internal/handlers"
	"native-web-demo/internal/middleware"
)

func main() {
	// 创建一个新的多路复用器 (Mux)
	mux := http.NewServeMux()

	// 从 'handlers' 包中导入处理器函数并注册它们
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/users", handlers.UsersHandler)

	log.Println("Starting web server on port 8080")
	log.Println("Available routes: / (GET) and /users (GET, POST)")

	// 使用 middleware.Logging 来包裹我们的 mux。
	// 请求的流转路径现在是：Logging Middleware -> mux -> 具体的 handler
	err := http.ListenAndServe(":8080", middleware.Logging(mux))
	if err != nil {
		log.Fatal(err)
	}
}