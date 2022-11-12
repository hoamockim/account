package profile

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tipee/account/app/controller"
	"github.com/tipee/account/app/middleware"
	"github.com/tipee/account/pkg/configs"
	"go.uber.org/zap"
)

const (
	USERPROFILE = "user-profile"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   USERPROFILE,
		Short: "user profile service",
		RunE: func(cmd *cobra.Command, args []string) error {
			router := gin.New()
			corsConfig := cors.New(cors.Config{
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
				AllowOrigins:     []string{"*"},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "TimezoneOffset"},
				AllowCredentials: false,
				MaxAge:           12 * time.Hour,
			})
			router.Use(corsConfig, middleware.Recovery())
			v1 := router.Group(USERPROFILE + "/v1")
			{
				app := v1.Group("/sys")
				{
					app.GET("/health-check", controller.HealthCheck)
				}
			}
			if err := router.Run(configs.AppURL()); err != nil {
				zap.S().Errorf("application error: %v", err)
				return err
			}
			return nil
		},
	}
}
