package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"sql/models"
	"sql/services"
	"sql/utils"
)

type UserController struct {
	userService *services.UserService
	jwtSecret   []byte
	tokenTTL    time.Duration
}

func (uc *UserController) Profile(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		models.Fail(c, http.StatusUnauthorized, models.CodeUnauthorized, "未登录")
		return
	}

	user, err := uc.userService.GetByID(userID)
	if err != nil {
		models.Fail(c, http.StatusNotFound, models.CodeNotFound, "用户不存在")
		return
	}

	models.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
} //能返回当前登录的用户信息

func NewUserController(us *services.UserService, jwtSecret []byte, tokenTTL time.Duration) *UserController {
	return &UserController{userService: us, jwtSecret: jwtSecret, tokenTTL: tokenTTL} // ★★★ 注入 service
}

func (uc *UserController) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.Fail(c, http.StatusBadRequest, models.CodeInvalidParam, "请求参数错误")
		return
	}

	user, err := uc.userService.Register(req.Username, req.Password)
	if err != nil {
		models.Fail(c, http.StatusBadRequest, models.CodeInvalidParam, err.Error())
		return
	}

	models.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

// 新的登录方式，通过JWT能返回Token
func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Fail(c, http.StatusBadRequest, models.CodeInvalidParam, "请求参数错误")
		return
	}

	user, err := uc.userService.Login(req.Username, req.Password)
	if err != nil {
		models.Fail(c, http.StatusUnauthorized, models.CodeUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateTokenWithSecret(user.ID, user.Username, uc.jwtSecret, uc.tokenTTL)
	if err != nil {
		models.Fail(c, http.StatusInternalServerError, models.CodeServerError, "生成 token 失败")
		return
	}

	models.Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

/*
func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	user, err := uc.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}
*/
