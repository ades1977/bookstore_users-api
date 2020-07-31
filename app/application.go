package app

import (
	"github.com/gin-gonic/gin"
)

var (
	route = gin.Default()
)

func StartApplication() {
	MapUrls()
	route.Run(":8080")
}
