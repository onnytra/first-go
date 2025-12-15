package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

func HTTPResponse(c *gin.Context, statusCode int, message string, data any, err any) {
	c.JSON(statusCode, Response{
		Message: message,
		Data:    data,
		Error:   err,
	})
}
