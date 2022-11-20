package profile

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tipee/account/app/controller"
	"github.com/tipee/account/app/middleware"
)

func setupRouter(r *gin.Engine) {
	//set-up middleware
	corsConfig := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "TimezoneOffset"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	r.Use(corsConfig, middleware.Recovery(), middleware.GracfefulShutdown())

	// add route
	v1 := r.Group(USERPROFILE + "/v1")
	{
		app := v1.Group("/sys")
		{
			app.GET("/health-check", controller.HealthCheck)
		}
	}
}
