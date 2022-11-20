package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/tipee/account/app/cmd/authentication"
	"github.com/tipee/account/app/cmd/profile"
	"github.com/tipee/account/app/middleware"
)

type GracefulShutdown interface {
	Switch(status bool)
}

func main() {
	cmd := &cobra.Command{
		Use:          "account",
		Short:        "account system",
		SilenceUsage: true,
	}

	cmd.AddCommand(authentication.Cmd())
	cmd.AddCommand(profile.Cmd())
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	shutdownService(cmd.Context())
}

//shutdownServer graceful shutdown server
func shutdownService(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	fmt.Println("service is shutting down gracefully")
	middleware.IsShuttingDown = true
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
