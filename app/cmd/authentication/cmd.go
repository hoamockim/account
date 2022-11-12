package authentication

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "auth",
		Short: "authentication service",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Account service is not available")
			return nil
		},
	}
}
