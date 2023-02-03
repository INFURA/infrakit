package proxy

import (
	"net/http"
	"os"

	"github.com/INFURA/infrakit/cmd/opts"
	"github.com/INFURA/infrakit/pkg/proxy/health"
	"github.com/go-logr/logr"
)

func RunProxy(log logr.Logger) error {
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

	return nil
}
