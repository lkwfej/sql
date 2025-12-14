package main

import (
	"sql/config"

	"github.com/gin-gonic/gin"
	"sql/controllers"
	"sql/database"
	"sql/routers"
	"sql/services"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)

	db := database.InitDB(cfg)

	userService := services.NewUserService(db) // ★★★ 传入 db
	userController := controllers.NewUserController(userService, cfg.JWTSecret, cfg.TokenTTL)
	//router := gin.Default()
	r := routers.SetupRouter(userController, cfg.JWTSecret)
	r.Run(cfg.Addr)
}
