package database

import (
	"log"

	"native-web-demo/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDatabase 初始化数据库连接并返回 *gorm.DB 实例。
func ConnectDatabase() *gorm.DB {
	// DSN (Data Source Name)
	// IMPORTANT: You MUST replace this with your actual MySQL DSN before running.
	// Example: "root:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:01010111@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移 (AutoMigrate)
	// 这将根据 models.User 结构体创建或更新数据库表。
	// 如果表不存在，GORM 会自动创建。如果表已存在，它会检查并添加新的字段。
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	log.Println("Database connection successful and auto-migrated.")
	return db
}
