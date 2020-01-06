package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp ...
func StartApp() {
	routerApp()

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
