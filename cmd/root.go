package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/spf13/cobra"
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
	// See if a config file was provided via env var
	if configFile := os.Getenv("INFRAKIT_CONFIG"); configFile != "" {
		if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
	}
}
