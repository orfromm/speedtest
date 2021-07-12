package cmd

import (
	"fmt"
	l "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"speedtest/client"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Speed Test",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			l.Fatal("can't init", err)
		}

		client, err := client.NewDefaultClientOf(viper.GetString("interface"))
		if err != nil {
			fmt.Printf("error creating client: %v", err)
		}

		// Pass an empty string to select the fastest server
		server, err := client.GetServer("")
		if err != nil {
			fmt.Printf("error getting server: %v", err)
		}

		dmbps, err := client.Download(server)
		if err != nil {
			fmt.Printf("error getting download: %v", err)
		}

		umbps, err := client.Upload(server)
		if err != nil {
			fmt.Printf("error getting upload: %v", err)
		}

		fmt.Printf("Ping: %3.2f ms | Download: %3.2f Mbps | Upload: %3.2f Mbps\n", server.Latency, dmbps, umbps)
	},
}

func defineStringFlag(cmd *cobra.Command, long, short, def, desc string, persistent bool) {
	if persistent {
		cmd.PersistentFlags().StringP(long, short, def, desc)
	} else {
		cmd.Flags().StringP(long, short, def, desc)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	defineStringFlag(RootCmd, "interface", "", "", "source IP or interface name", true)
}
