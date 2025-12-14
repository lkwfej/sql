package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 用于统一返回结构
type Response struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 错误信息
	Data    interface{} `json:"data"`    // 返回数据
}

// 常见的错误码
const (
	CodeSuccess      = 0     // 成功
	CodeInvalidParam = 40001 // 参数无效
	CodeUnauthorized = 40101 // 未授权
	CodeForbidden    = 40301 // 无权限
	CodeNotFound     = 40401 // 资源未找到
	CodeServerError  = 50000 // 服务器错误
)

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "ok",
		Data:    data,
	})
}

// Fail 返回失败响应（建议：HTTP 状态码与业务 code 同时返回）
func Fail(c *gin.Context, httpStatus int, code int, message string) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
