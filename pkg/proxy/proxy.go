package proxy

import (
	"fmt"
	"net/http"
	"os"

	"github.com/INFURA/infrakit/cmd/opts"
	"github.com/INFURA/infrakit/pkg/proxy/evm"
	"github.com/INFURA/infrakit/pkg/proxy/health"
	"github.com/go-logr/logr"
)

func RunProxy(log logr.Logger) error {
	proxy, err := evm.NewProxy(evm.Opts{
		Log:                      log,
		Addr:                     opts.ProxyConfig().ListenAddr,
		EthBootstrapEndpoint:     opts.ProxyConfig().EthBootstrapEndpoint,
		NodeRegistryContract:     opts.ProxyConfig().NodeRegistryContract,
		CustomerRegistryContract: opts.ProxyConfig().CustomerRegistryContract,
	})
	if err != nil {
		return fmt.Errorf("failed to create proxy: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/ready", health.HTTPHandler())

	healthServer := &http.Server{
		Addr:    opts.ProxyConfig().HealthListenAddr,
		Handler: mux,
	}

	go func() {
		log.Error(healthServer.ListenAndServe(), "health server finished unexpectedly")
		os.Exit(1)
	}()

	log.Info("Starting up server", "addr", opts.ProxyConfig().ListenAddr)

	if err := proxy.Start(); err != nil {
		log.Error(err, "evm proxy server finished unexpectedly")
		os.Exit(1)
	}

	return nil
}
