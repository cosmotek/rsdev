package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"gitlab.com/therecspot/rsdev/auth"
	"gitlab.com/therecspot/rsdev/net"
)

var (
	red   = ansi.ColorCode("red+h")
	cyan  = ansi.ColorCode("cyan+h")
	reset = ansi.ColorCode("reset")
)

var port string
var Proxy = &cobra.Command{
	Use:   "proxy [endpoint]",
	Short: "HTTP Authentication Proxy",
	Long:  `Automatically handles authentication under the hood allowing you to make requests without ever setting up authorization`,
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostEndpoint := args[0]
		grapqlEndpoint := hostEndpoint + "/query"

		proxy, err := net.NewHeaderProxy(hostEndpoint)
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
				timeUntil := headerInfo.ExpiresAt.Sub(time.Now())
				timer := time.NewTimer(timeUntil - (time.Minute * 1))
				defer timer.Stop()

				// block until timer completes or exit started
				select {
				case <-timer.C:
					headerInfo, err = auth.RefreshToken(ctx, grapqlEndpoint, headerInfo.Value)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					proxy.Set(headerInfo.HeaderKey, headerInfo.Value)
					fmt.Println(cyan, "Token refreshed...", reset)
				case <-ctx.Done():
					break loop
				}
			}
		}()

		go func() {
			var err error
			if port != "" {
				err = proxy.StartProxy(ctx, &port)
			} else {
				err = proxy.StartProxy(ctx, nil)
			}
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
		fmt.Println(red, "\nProxy stopped... Goodbye.", reset)
	},
}

func init() {
	Proxy.Flags().StringVarP(&port, "port", "p", "", "sets the proxy listener port")
}
