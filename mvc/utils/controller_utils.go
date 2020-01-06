package utils

import (
	"github.com/gin-gonic/gin"
)

// Response make response json to client
func Response(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}

	c.JSON(status, body)

}
