package api

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/blacklisted", AddBlacklisted)
	return r
}
