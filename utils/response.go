package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func HttpRespSuccess(c *gin.Context, status int, message string, data interface{}) {
	resp := Response{
		Success: true,
		Error:   nil,
		Message: message,
		Data:    data,
	}

	c.JSON(status, resp)
}

func HttpRespFailed(c *gin.Context, status int, message string) {
	resp := Response{
		Success: false,
		Data:    nil,
		Error:   message,
	}

	c.JSON(status, resp)
}
