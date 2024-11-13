package middlewares

import (
	"courses-api/src/config/envs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	KEY := envs.LoadEnvs(".env").Get("COURSES_API_KEY")
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		fmt.Println(apiKey)

		if apiKey != KEY {
			ErrorResponse(c, 401, "Invalid API Key")
			return
		}

		c.Next()
	}
}

// ErrorResponse sets CORS headers and aborts the request with a JSON error response
func ErrorResponse(c *gin.Context, status int, message string) {
	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Header("Access-Control-Allow-Credentials", "true")
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}
