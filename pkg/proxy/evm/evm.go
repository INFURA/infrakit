package evm

import (
	"fmt"
	"net/http"

	"github.com/INFURA/infrakit/pkg/proxy/server"
	"github.com/go-logr/logr"
)

type Proxy struct {
	server *server.Server
	log    logr.Logger
}

type Opts struct {
	Addr string
	Log  logr.Logger
}

func NewProxy(opts Opts) (*Proxy, error) {
	serv := server.NewServer(server.Opts{
		Addr: opts.Addr,
	})

	serv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello infrakit."))
	})

	return &Proxy{
		server: serv,
		log:    opts.Log,
	}, nil
}

func (p Proxy) Start() error {
	err := p.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("evm proxy Start returned error: %w", err)
	}

	p.log.Info("evm proxy server closed successfully")

	return nil
}
