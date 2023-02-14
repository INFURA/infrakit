package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/INFURA/infrakit/cmd/opts"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/spf13/cobra"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	k = koanf.New(".")
)

var rootCmd = &cobra.Command{
	Use:   "infrakit",
	Short: "InfraKit enables participation on the Decentralized Infrastructure Network",
	Long: `InfraKit is a mono-binary for participating in the Decentralized 
Infrastructure Network. It can contains the kubernetes operator as well
as other ancillary runtimes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	setupLog.Info("initializing config")

	// See if a config file was provided via env var
	if configFile := os.Getenv("INFRAKIT_CONFIG"); configFile != "" {
		if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
	}

	if err := k.Load(env.Provider("INFRAKIT_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "INFRAKIT_")), "__", ".", -1)
	}), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	var logConfig opts.LogConfig

	if err := k.Unmarshal("logger", &logConfig); err != nil {
		log.Fatalf("error parsing log config: %v", err)
	}

	logOpts := opts.FromLogConfig(logConfig)
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&logOpts)))
}
