package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"sql/models"
	"sql/utils"
	"strings"
)

//var jwtKey = []byte("secret_key_123") // 和你生成 token 用的保持一致

func JWTAuth(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			models.Fail(c, http.StatusUnauthorized, models.CodeUnauthorized, "未登录")
			c.Abort()
			return
		}

		// Safer parsing than Split: handle extra spaces/newlines copied from Postman.
		if !strings.HasPrefix(authHeader, "Bearer ") {
			models.Fail(c, http.StatusUnauthorized, models.CodeUnauthorized, "Authorization 格式错误，应为 Bearer <token>")
			c.Abort()
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		// Guard against invisible characters
		tokenStr = string(bytes.TrimSpace([]byte(tokenStr)))
		if tokenStr == "" {
			models.Fail(c, http.StatusUnauthorized, models.CodeUnauthorized, "Token 为空")
			c.Abort()
			return
		}

		claims, err := utils.ParseTokenWithSecret(tokenStr, jwtSecret)
		if err != nil {
			models.Fail(c, http.StatusUnauthorized, models.CodeUnauthorized, "Token 无效或已过期")
			c.Abort()
			return
		}
		// 👇 把用户信息存进 Context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
