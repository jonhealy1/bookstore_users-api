package app

import (
	pingcontroller "github.com/jonhealy1/bookstore_/bookstore_users-api/controllers/ping"
	"github.com/jonhealy1/bookstore_/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", pingcontroller.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
