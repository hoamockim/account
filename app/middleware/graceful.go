package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tipee/account/errors"
)

var (
	IsShuttingDown bool
)

func GracfefulShutdown() gin.HandlerFunc {
	return func(c *gin.Context) {
		if IsShuttingDown {
			c.Abort()
			handlerShuttingDown(c)
		}
		c.Next()
	}
}

func handlerShuttingDown(c *gin.Context) {
	res := errors.HandlerErr(c.GetHeader("X-Request-ID"), c.Request, errors.OffService,
		&errors.ClientInfo{
			IP:     c.ClientIP(),
			Uri:    c.Request.URL.Path,
			Method: c.Request.Method,
		})
	c.JSON(500, res)
}
