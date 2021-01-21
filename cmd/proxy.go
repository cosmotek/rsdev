package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/therecspot/rsdev/auth"
	"gitlab.com/therecspot/rsdev/net"
)

var Proxy = &cobra.Command{
	Use:   "proxy",
	Short: "HTTP Authentication Proxy",
	Long:  `Automatically handles authentication under the hood allowing you to make requests without ever setting up authorization`,
	Run: func(cmd *cobra.Command, args []string) {
		grapqlEndpoint := args[0]

		proxy, err := net.NewHeaderProxy(grapqlEndpoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// create a context with cancel in case the user interrupts the process
		ctx, cancel := context.WithCancel(context.Background())

		// get a starting token
		headerInfo, err := auth.StartPromptedSMSVerification(ctx, grapqlEndpoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		proxy.Set(headerInfo.HeaderKey, headerInfo.Value)

		go func() {
		loop:
			for {
				timeUntil := time.Now().Sub(headerInfo.ExpiresAt)
				timer := time.NewTimer(timeUntil - (time.Minute * 1))

				// block until timer completes or exit started
				select {
				case <-timer.C:
				case <-ctx.Done():
					break loop
				}

				headerInfo, err = auth.RefreshToken(ctx, grapqlEndpoint, headerInfo.Value)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				proxy.Set(headerInfo.HeaderKey, headerInfo.Value)
			}
		}()

		go func() {
			err = proxy.StartProxy(ctx, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}()

		exitSignal := make(chan os.Signal)
		signal.Notify(exitSignal, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGINT)

		// block until exit
		<-exitSignal
		cancel()
		fmt.Println("\nProxy stopped...\nGoodbye.")
	},
}
