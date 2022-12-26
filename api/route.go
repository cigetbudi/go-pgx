package api

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/blacklisted", AddBlacklisted)
	r.POST("/auth/register", Register)
	r.POST("/auth/login", Login)
	return r
}
