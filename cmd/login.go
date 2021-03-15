package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/therecspot/rsdev/auth"
)

var Login = &cobra.Command{
	Use:   "login [endpoint]",
	Short: "Service Login",
	Long:  `Service Login Helper`,
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostEndpoint := args[0]
		grapqlEndpoint := hostEndpoint + "/query"

		ctx := context.Background()

		_, err := auth.StartPromptedSMSVerification(ctx, grapqlEndpoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
