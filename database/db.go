package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sql/config"
	"sql/models"
)

var DB *gorm.DB

func InitDB(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	fmt.Println("数据库连接成功")
	err = db.AutoMigrate(&models.Role{}, &models.User{})
	if err != nil {
		log.Fatal("自动迁移失败：", err)
	}
	DB = db
	return db
}
