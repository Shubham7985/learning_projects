package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Success Response
Ye function har successful API response ke liye use hoga
*/
func Success(c *gin.Context, message string, data ...interface{}) {
	response := gin.H{
		"success": true,
		"message": message,
	}

	if len(data) > 0 {
		response["data"] = data[0]
	}

	c.JSON(http.StatusOK, response)
}

/*
Error Response
Ye function har error ke liye use hoga
*/
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   message,
	})
}
