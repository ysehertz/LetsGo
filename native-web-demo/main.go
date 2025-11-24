package main

import (
	"fmt"
	"log"
	"net/http"
)

// home 是一个处理器函数 (handler function)
// 它的作用类似于 Java Servlet 中的 doGet 或 doPost 方法。
func home(w http.ResponseWriter, r *http.Request) {
	// w: http.ResponseWriter -> 对应 Java 的 HttpServletResponse
	// r: *http.Request      -> 对应 Java 的 HttpServletRequest

	// 向响应写入内容
	fmt.Fprintln(w, "Hello Native Go")
}

func main() {
	// http.HandleFunc 类似于在 web.xml 中配置一个 <servlet-mapping>。
	// 它将 URL 路径 "/" 与我们的 home 函数绑定。
	http.HandleFunc("/", home)

	// http.ListenAndServe 启动一个 HTTP 服务器并监听 8080 端口。
	// 这类似于启动一个内嵌的 Tomcat 或 Jetty 服务器。
	// log.Fatal 会在出现错误时打印日志并退出程序。
	log.Println("Starting web server on port 8080")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
