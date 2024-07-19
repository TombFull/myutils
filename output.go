package myutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// OutputSuccess 请求成功返回
func OutputSuccess(c *gin.Context, code int, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"result":  data,
		"code":    code,
		"status":  1,
	})
}

// OutputFail 请求失败返回
func OutputFail(c *gin.Context, code int, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"result":  data,
		"code":    code,
		"status":  -1,
	})
}

type SwaggerResponse struct {
	Message string                 `json:"message"`
	Result  map[string]interface{} `json:"result"`
	Code    int                    `json:"code"`
	Status  int                    `json:"status"`
}
