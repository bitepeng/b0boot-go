package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ECHO 直接返回
func ECHO(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

// OK 成功返回 code=0
func OK(msg string, data interface{}, c *gin.Context) {
	JSON(0, msg, data, c)
}

// ERROR 失败返回 code=400
func ERROR(msg string, c *gin.Context) {
	JSON(400, msg, nil, c)
}

// JSON 通用返回参数
// code=0 成功OK
// code=400 错误ERROR
// code=401 需要登录
func JSON(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// PAGE 返回分页数据
func PAGE(count int64, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
	})
}
