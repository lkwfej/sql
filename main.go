package main

import (
	"sql/controllers"
	"sql/database"
	"sql/routers"
	"sql/services"
)

func main() {
	db := database.InitDB()

	userService := services.NewUserService(db) // ★★★ 传入 db
	userController := controllers.NewUserController(userService)
	//router := gin.Default()
	r := routers.SetupRouter(userController)
	r.Run(":9090")
}
