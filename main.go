package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/therecspot/rsdev/cmd"
)

var rootCmd = &cobra.Command{
	Use:     "rsdev",
	Version: "v1.0.0",
	Short:   "RecSpot Dev Tool",
	Long:    `A collection of useful helpers, automated processes and other utilities for developing RecSpot software`,
}

func main() {
	rootCmd.AddCommand(cmd.Proxy)
	rootCmd.AddCommand(cmd.Login)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
