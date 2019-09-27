package router

import (
	. "web_haoge/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/haoge", "./static")

	router.GET("/users", Users)

	router.POST("/user", Store)

	router.GET("/user", FindUser)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	return router
}
