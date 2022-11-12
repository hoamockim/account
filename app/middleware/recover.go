package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tipee/account/errors"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				err := fmt.Errorf("%v", rec)
				zap.S().Errorw("Panicking during process request", zap.Error(err))
				c.Abort()
				handlerErrPanic(c, errors.Panicked)
			}
		}()
		c.Next()
	}
}

func handlerErrPanic(c *gin.Context, err errors.ErrorCode) {
	code, res := errors.HandlerErr(c.GetHeader("X-Request-ID"), c.Request, errors.Panicked,
		&errors.ClientInfo{
			IP:     c.ClientIP(),
			Uri:    c.Request.URL.Path,
			Method: c.Request.Method,
		})
	c.JSON(code, res)
}
