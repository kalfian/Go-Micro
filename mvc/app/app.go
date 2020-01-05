package app

import (
	"net/http"

	"github.com/kalfian/Go-Micro/mvc/controllers"
)

// StartApp ...
func StartApp() {
	http.HandleFunc("/users", controllers.GetUsers)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
