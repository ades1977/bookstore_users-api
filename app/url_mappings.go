package app

import (
	"github.com/ades1977/bookstore_users-api/controllers/ping"
	"github.com/ades1977/bookstore_users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/GetUser/:user_id", users.GetUser)
	router.POST("/users/CreateUser", users.CreateUser)
}
