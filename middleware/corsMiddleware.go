package middleware

import (
	"Go_server/auth"
	"Go_server/model"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func VerifyAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	err := auth.TokenValid(token)
	if err != nil {
		c.JSON(401, "ku")
		return
	} else {
		_, ok := model.SaveToken[token]
		if !ok {
			c.JSON(401, "sa")
			return
		}
	}
	c.Next()
}
