package health

import (
	"net/http"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	"github.com/INFURA/infrakit/cmd/opts"
)

const (
	healthcheckInitialDelay = 1 * time.Second
	healthcheckTimeout      = 1 * time.Second
	healthcheckInterval     = 5 * time.Second
)

func HTTPHandler() http.HandlerFunc {
	h := gosundheit.New()

	pinger := checks.NewDialPinger("tcp", opts.ProxyConfig().ListenAddr)
	pingCheck := checks.Must(checks.NewPingCheck("api.reachable", pinger))

	err := h.RegisterCheck(
		pingCheck,
		gosundheit.InitialDelay(healthcheckInitialDelay),
		gosundheit.ExecutionPeriod(healthcheckInterval),
		gosundheit.ExecutionTimeout(healthcheckTimeout),
	)
	if err != nil {
		panic(err)
	}

	return healthhttp.HandleHealthJSON(h)
}
