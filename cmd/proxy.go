package cmd

import (
	"github.com/spf13/cobra"
	ctrl "sigs.k8s.io/controller-runtime"
)

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
	log := ctrl.Log.WithName("proxy")

	log.Info("logger proxy called")
}
