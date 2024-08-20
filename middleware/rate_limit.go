package middleware

import (
	"cors/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func RateLimit(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		v, err := redisClient.Get(c, key).Int()
		if err != nil && err != redis.Nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if v >= 100 {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		redisClient.Incr(c, key)
		redisClient.Expire(c, key, time.Minute)

		c.Next()
	}
}