package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging 是一个中间件函数，用于记录每个请求的信息。
// 它接收一个 http.Handler 作为参数，并返回一个新的 http.Handler。
func Logging(next http.Handler) http.Handler {
	// http.HandlerFunc 是一个适配器，允许普通的函数作为 http.Handler 使用。
	// 我们返回的这个函数就是实际处理请求的中间件处理器。
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 在请求到达核心业务逻辑之前执行
		startTime := time.Now()
		log.Printf("--> %s %s", r.Method, r.URL.Path)
		log.Printf("第一次提交")
		log.Printf("第二次提交")
		log.Printf("第三次提交")
		log.Printf("第四次提交")
		log.Printf("第五次提交")

		// 2. 调用下一个处理器 (next.ServeHTTP)
		// 这可能是另一个中间件，也可能是我们最终的 mux。
		// 请求会在这里被"暂停"，直到核心业务逻辑处理完成。
		next.ServeHTTP(w, r)

		// 3. 在核心业务逻辑处理完成之后执行
		duration := time.Since(startTime)
		log.Printf("<-- %s %s (%v)", r.Method, r.URL.Path, duration)
	})
}
