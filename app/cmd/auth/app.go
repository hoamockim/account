package auth

import "github.com/spf13/cobra"

func Run(args []string) *cobra.Command {
	return &cobra.Command{
		Use:   "auth",
		Short: "start authentication service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
