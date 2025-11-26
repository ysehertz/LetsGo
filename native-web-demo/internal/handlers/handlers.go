package handlers

import (
	"log"
	"net/http"

	"native-web-demo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm" // 导入 gorm 包
)

// Handler 结构体持有数据库连接，以便其方法可以访问数据库。
// 这是在 Go 中实现依赖注入的一种常见模式。
type Handler struct {
	DB *gorm.DB
}

// Home 是一个 Gin 处理器函数。
func Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gin!")
}

// GetUsers 处理 GET /users 请求，现在从数据库中查询用户。
func (h *Handler) GetUsers(c *gin.Context) {
	var users []models.User
	// 使用 h.DB.Find 从数据库查询所有用户。
	// GORM 会自动将查询结果映射到 users 切片中。
	if result := h.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser 处理 POST /users 请求，现在将用户保存到数据库。
func (h *Handler) CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用 h.DB.Create 将新用户保存到数据库。
	// newUser 结构体嵌入了 gorm.Model，GORM 会自动为其分配 ID 和时间戳。
	if result := h.DB.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	log.Printf("User created: %+v", newUser)
	c.JSON(http.StatusCreated, newUser)
}
