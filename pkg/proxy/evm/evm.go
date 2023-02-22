package evm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/INFURA/infrakit/pkg/proxy/server"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-logr/logr"
)

type Proxy struct {
	server    *server.Server
	log       logr.Logger
	ethclient *ethclient.Client
}

type Opts struct {
	Addr                 string
	EthBootstrapEndpoint string
	Log                  logr.Logger
}

func NewProxy(opts Opts) (*Proxy, error) {
	serv := server.NewServer(server.Opts{
		Addr: opts.Addr,
	})

	ethclient, err := ethclient.Dial(opts.EthBootstrapEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum: %w", err)
	}

	serv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		blockNumber, err := ethclient.BlockNumber(context.TODO())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := fmt.Sprintf("Hello infrakit. Current block number is %d", blockNumber)
		_, _ = w.Write([]byte(response))
	})

	return &Proxy{
		server:    serv,
		log:       opts.Log,
		ethclient: ethclient,
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
