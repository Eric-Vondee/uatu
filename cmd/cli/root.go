package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uatu-cli",
	Short: "Administrative and data-seeding CLI for the uatu service",
	// Runtime failures are returned from RunE and reported by main; don't let
	// Cobra also dump usage/errors on top of them.
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command with a context that is cancelled on SIGINT or
// SIGTERM, so long-running commands can stop cleanly on Ctrl+C.
func Execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return rootCmd.ExecuteContext(ctx)
}
