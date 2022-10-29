package profile

import "github.com/spf13/cobra"

func Run(args []string) *cobra.Command {
	return &cobra.Command{
		Use:   "userProfile",
		Short: "start user profile service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
