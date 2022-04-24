package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartAplication() {
	MapUrls() // Add in router url's endpoints
	router.Run(":8080")
}
