package routers //路由保护
import (
	"github.com/gin-gonic/gin"
	"sql/controllers"
	"sql/middleware"
)

func SetupRouter(userController *controllers.UserController, jwtSecret []byte) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// 🔐 受保护路由
	api := r.Group("/api")
	api.Use(middleware.JWTAuth(jwtSecret))
	{
		api.GET("/profile", userController.Profile)
	}

	return r
}

/*
func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// 受保护接口
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user_id":  c.GetUint("user_id"),
				"username": c.GetString("username"),
			})
		})
	}

	return r
}
*/
/*
func SetupRouter(userController *controllers.UserController) *gin.Engine {

	r := gin.Default()

	// ★★★ 使用 main.go 注入进来的 controller
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	return r
}
*/
