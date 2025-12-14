package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key_123") // 和你生成 token 用的保持一致

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization 格式错误"})
			c.Abort()
			return
		}

		tokenStr := parts[1]
		println("RAW TOKEN:", tokenStr)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token 无效"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token 解析失败"})
			c.Abort()
			return
		}

		// 👇 把用户信息存进 Context
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("username", claims["username"].(string))

		c.Next()
	}
}
