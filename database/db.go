package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sql/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "admin_sql:admin_sql@tcp(127.0.0.1:3306)/goland_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	fmt.Println("数据库连接成功")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("自动迁移失败：", err)
	}
	DB = db
	return db
}
func InitTables() {
	DB.AutoMigrate(
		&models.Role{},
		&models.User{},
	)
}
