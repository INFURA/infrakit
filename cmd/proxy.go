package cmd

import (
	"os"

	"github.com/INFURA/infrakit/cmd/opts"
	"github.com/INFURA/infrakit/pkg/proxy"
	"github.com/spf13/cobra"
	ctrl "sigs.k8s.io/controller-runtime"
)

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

	if err := k.Unmarshal("proxy", opts.ProxyConfig()); err != nil {
		log.Error(err, "unable to parse proxy config")
		os.Exit(1)
	}

	log.Info("logger proxy called")

	if err := proxy.RunProxy(log); err != nil {
		log.Error(err, "error running proxy")
		os.Exit(1)
	}
}
