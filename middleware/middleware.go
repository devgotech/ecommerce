package middleware

import (
	"ecommerce/tokens"
	"net/http"

	// "net/url"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			ClientLogin := c.Request.Header.Get("/user/login")
			if ClientLogin == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
				c.Abort()
				return
			}
			ClientSignup := c.Request.Header.Get("/user/signup")
			if ClientSignup == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
				c.Abort()
				return
			}
		*/
		

		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}
		claims, err := tokens.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
