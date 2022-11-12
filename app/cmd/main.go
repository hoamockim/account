package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tipee/account/app/cmd/authentication"
	"github.com/tipee/account/app/cmd/profile"
)

func main() {
	cmd := &cobra.Command{
		Use:          "account",
		Short:        "account system",
		SilenceUsage: true,
	}
	//register command
	cmd.AddCommand(authentication.Cmd())
	cmd.AddCommand(profile.Cmd())
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
