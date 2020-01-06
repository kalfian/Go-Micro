package app

import (
	"github.com/kalfian/Go-Micro/mvc/controllers"
)

func routerApp() {
	router.GET("/users/:user_id", controllers.GetUsers)
}
