package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging 返回一个 Gin 中间件函数，用于记录每个请求的信息。
// Gin 中间件的签名是 func(c *gin.Context)。
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 在请求到达核心业务逻辑之前执行
		startTime := time.Now()
		log.Printf("--> %s %s", c.Request.Method, c.Request.URL.Path)

		// 2. 调用链中的下一个中间件或最终的处理器。
		// c.Next() 类似于 net/http 中间件中的 next.ServeHTTP(w, r)，
		// 它将控制权传递给下一个处理器，直到所有处理器执行完毕。
		c.Next()

		// 3. 在所有后续处理器执行完成之后执行
		duration := time.Since(startTime)
		log.Printf("<-- %s %s (%v)", c.Request.Method, c.Request.URL.Path, duration)
	}
}