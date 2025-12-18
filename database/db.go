package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"sql/config"
	"sql/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	cfg := config.LoadDBConfig()
	dsn := cfg.DSN()

	db, err := openWithRetry(dsn, 5, time.Second)
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

func openWithRetry(dsn string, attempts int, delay time.Duration) (*gorm.DB, error) {
	var lastErr error
	for i := 0; i < attempts; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		lastErr = err
		time.Sleep(delay)
	}
	return nil, lastErr
}
