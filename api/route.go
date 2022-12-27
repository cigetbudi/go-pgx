package api

import (
	"go-pgx/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/blacklisted", AddBlacklisted)
	r.POST("/auth/register", Register)
	r.POST("/auth/login", Login)

	post := r.Group("/post")
	post.Use(middleware.JwtAuthMiddleware())
	post.POST("/", CreatePost)
	post.POST("/like/:id", LikePost)
	post.POST("/unlike/:id", UnlikePost)
	post.GET("/likes/:id", GetLikesCount)

	return r
}
