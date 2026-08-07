package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(
	c *gin.Context,
	code int,
	message string,
	errors interface{},
) {

	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"errors":  errors,
	})
}

func SuccessResponse(
	c *gin.Context,
	code int,
	message string,
	data interface{},
) {

	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}
