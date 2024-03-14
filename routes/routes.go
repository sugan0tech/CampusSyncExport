package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/export", indexGreet)
	server.POST("/export", export)
}
