package cmd

import "github.com/spf13/cobra"

type ProxyCmdConfig struct {
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run the InfraKit RPC proxy",
	Run:   proxyMain,
}

func proxyMain(cmd *cobra.Command, args []string) {
	// Do Stuff Here
}
