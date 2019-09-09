package middleware

import "github.com/gin-gonic/gin"

func Session(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
