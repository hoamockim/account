package main

import "github.com/spf13/cobra"

type IApp interface {
	Run(args []string) *cobra.Command
}

func main() {

}
