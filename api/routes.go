package api

import (
	"cors/auth"
	"cors/config"
	"cors/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors(cfg))
	r.Use(middleware.RateLimit(cfg))

	r.POST("/login", Login)

	authorized := r.Group("/api")
	authorized.Use(auth.JWTMiddleware())
	{
		authorized.GET("/users", GetUsers)
		authorized.POST("/users", CreateUser)
	}

	return r
}
