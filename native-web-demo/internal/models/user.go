package models

import "gorm.io/gorm"

// User 定义了我们的用户模型，并嵌入了 gorm.Model。
type User struct {
	gorm.Model        // 嵌入 gorm.Model，自动获得 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"` // gorm:"unique" 确保 Email 字段在数据库中是唯一的
}
