package middleware

import (
	"github.com/gin-gonic/gin"
	"plover.com/spider/handler"
	"plover.com/spider/pkg/errno"
	"plover.com/spider/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
