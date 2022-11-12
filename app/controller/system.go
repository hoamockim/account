package controller

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
	success(c, nil)
}
