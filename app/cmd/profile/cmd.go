package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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
			container := initDependecies()
			if err := container.Invoke(func() {
				router := gin.New()
				setupRouter(router)
				go func() {
					if err := router.Run(configs.AppURL()); err != nil {
						zap.S().Errorf("application error: %v", err)
					}
				}()

			}); err != nil {
				return err
			}

			return nil
		},
	}
}
