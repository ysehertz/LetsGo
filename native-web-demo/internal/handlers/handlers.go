package handlers

import (
	"log"
	"net/http"

	"native-web-demo/internal/models"

	"github.com/gin-gonic/gin"
)

// Home 是一个 Gin 处理器函数。
func Home(c *gin.Context) {
	// c.String 简化了返回字符串响应的操作，自动处理 Content-Type。
	c.String(http.StatusOK, "Hello Gin!")
}

// GetUsers 处理 GET /users 请求。
func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}

	// c.JSON 简化了返回 JSON 响应的操作。
	// 它会自动设置 Content-Type: application/json 并进行 JSON 序列化。
	c.JSON(http.StatusOK, users)
}

// CreateUser 处理 POST /users 请求。
func CreateUser(c *gin.Context) {
	var newUser models.User

	// c.BindJSON 自动从请求体中解析 JSON 并填充到 newUser 结构体中。
	// 它还内置了验证功能（如果结构体有验证标签）。
	// 这比手动的 json.NewDecoder().Decode() 要方便得多。
	if err := c.BindJSON(&newUser); err != nil {
		// 如果绑定失败，Gin 会自动返回一个 400 Bad Request 响应。
		// 我们这里可以简单地返回，或者添加自定义的错误处理。
		return
	}

	log.Printf("Received new user: %+v", newUser)

	// 返回创建的用户对象和 201 Created 状态码。
	c.JSON(http.StatusCreated, newUser)
}