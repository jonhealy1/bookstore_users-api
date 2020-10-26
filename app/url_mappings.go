package app

import (
	"github.com/jonhealy1/bookstore_/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
